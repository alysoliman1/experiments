package nodes

import (
	"fmt"

	"github.com/asoliman1/experiments/gaps/internal/pkg/utils"
)

// NewRoot creates a new root at a certain gap length.
func NewRoot(GapDepth int) Node {
	root := Node{
		Sequence: "",
		GapDepth: GapDepth,
		Level:    0,
		Value:    -1,
	}
	for range GapDepth {
		root.Lengths = append(root.Lengths, 0)
		root.Tails = append(root.Tails, -1)
		root.LastRecurrences = append(root.LastRecurrences, -1)
		root.GapBuckets = append(root.GapBuckets, []int{})
	}
	return root
}

// The Node object represents a finite collection of data associated
// with every node in a binary tree. Given the data for a node, we are able
// to decide if the node has a left or right child, and, if so, produce
// data for those children. Thus, given the data at a root node, we are
// able to generate an arbitrarly large binary tree.
type Node struct {
	// The entire 0 level sequence.
	Sequence string `dynamodbav:"sequence"`

	GapDepth int `dynamodbav:"gapDepth"`

	// The length of the sequence.
	Lengths []int `dynamodbav:"lengths"`

	// The tail is the value of the last element in the sequence
	// The index of the tail is simply the length of the sequence (assuming we start counting from 1).
	Tails []int `dynamodbav:"tails"`

	// The index of the last recurring element. An element is recurring if
	// it equals the element right after it. If there's no recurring elements
	// then the index is is set to -1.
	LastRecurrences []int `dynamodbav:"lastRecurrences"`

	GapBuckets [][]int `dynamodbav:"gapBuckets"`
	Level      int     `dynamodbav:"level"`
	Value      int     `dynamodbav:"value"`
}

func (n *Node) extend(newElement int) (child Node, ok bool) {
	child = Node{
		Sequence:        fmt.Sprintf("%s%d", n.Sequence, newElement),
		Value:           newElement,
		Level:           n.Level + 1,
		GapDepth:        n.GapDepth,
		Lengths:         make([]int, n.GapDepth),
		Tails:           make([]int, n.GapDepth),
		LastRecurrences: make([]int, n.GapDepth),
	}
	copy(child.Lengths, n.Lengths)
	copy(child.Tails, n.Tails)
	copy(child.LastRecurrences, n.LastRecurrences)
	var buckets []*utils.Bucket
	for i := range n.GapDepth {
		buckets = append(buckets, utils.NewBucket(n.GapBuckets[i], f(i)))
	}
	for i := range n.GapDepth {
		// If the new element matches the sequence tail then the tail
		// is a recurring element. If there's an existing recurring element
		// then we need to check for gaps between the existing recurring
		// element and the new recurring element.
		if newElement == child.Tails[i] && child.LastRecurrences[i] >= 0 {
			tailIndex := child.Lengths[i]
			gap := tailIndex - child.LastRecurrences[i]
			if !buckets[i].Add(gap) {
				return Node{}, false
			}
			newElement = gap
			child.LastRecurrences[i] = tailIndex
			child.Lengths[i] += 1
			continue
		}

		// If the new element matches the sequence tail then the tail
		// is a recurring element. If there are no recurring elements then
		// we simply update the LastRecurrences index. We then break
		// because there are no gaps and so there are no new elements
		// for the above layers.
		if newElement == child.Tails[i] && child.LastRecurrences[i] < 0 {
			tailIndex := child.Lengths[i]
			child.LastRecurrences[i] = tailIndex
			child.Lengths[i] += 1
			break
		}

		// If the new element does not match the sequence tail then we update
		// the tail value to be the new element's value. We then break
		// because there are no gaps and so there are no new elements
		// for the above layers.
		if newElement != child.Tails[i] {
			child.Tails[i] = newElement
			child.Lengths[i] += 1
			break
		}
	}
	ok = true
	for i := range n.GapDepth {
		child.GapBuckets = append(child.GapBuckets, buckets[i].list)
	}
	return
}

// If a node with the current context has a left child then the context
// for that child is returned. Otherwise if the node doesn't have a left
// child then a nil is returned
func (n Node) Left() (leftChild Node, ok bool) {
	return n.extend(0)
}

// If a node with the current context has a left child then the context
// for that child is returned. Otherwise if the node doesn't have a left
// child then a nil is returned.
func (n Node) Right() (rightChild Node, ok bool) {
	return n.extend(1)
}

func f(i int) int {
	if i == 0 {
		return 3
	}
	return 4
}
