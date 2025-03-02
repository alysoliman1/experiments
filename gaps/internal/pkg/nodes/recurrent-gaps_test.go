package nodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeftRight(t *testing.T) {
	tests := []struct {
		name    string
		node    Node
		leftOk  bool
		left    Node
		rightOk bool
		right   Node
	}{
		{
			name: "empty word",
			node: Node{
				Sequence:        "",
				GapDepth:        4,
				Lengths:         []int{0, 0, 0, 0},
				Tails:           []int{-1, -1, -1, -1},
				LastRecurrences: []int{-1, -1, -1, -1},
				GapBuckets:      [][]int{{}, {}, {}, {}},
				Level:           0,
				Value:           -1,
			},
			leftOk: true,
			left: Node{
				Sequence:        "0",
				GapDepth:        4,
				Lengths:         []int{1, 0, 0, 0},
				Tails:           []int{0, -1, -1, -1},
				LastRecurrences: []int{-1, -1, -1, -1},
				GapBuckets:      [][]int{{}, {}, {}, {}},
				Level:           1,
				Value:           0,
			},
			rightOk: true,
			right: Node{
				Sequence:        "1",
				GapDepth:        4,
				Lengths:         []int{1, 0, 0, 0},
				Tails:           []int{1, -1, -1, -1},
				LastRecurrences: []int{-1, -1, -1, -1},
				GapBuckets:      [][]int{{}, {}, {}, {}},
				Level:           1,
				Value:           1,
			},
		},
		{
			name: "sequence: 1",
			node: Node{
				Sequence:        "1",
				GapDepth:        4,
				Lengths:         []int{1, 0, 0, 0},
				Tails:           []int{1, -1, -1, -1},
				LastRecurrences: []int{-1, -1, -1, -1},
				GapBuckets:      [][]int{{}, {}, {}, {}},
				Level:           1,
				Value:           1,
			},
			leftOk: true,
			left: Node{
				Sequence:        "10",
				GapDepth:        4,
				Lengths:         []int{2, 0, 0, 0},
				Tails:           []int{0, -1, -1, -1},
				LastRecurrences: []int{-1, -1, -1, -1},
				GapBuckets:      [][]int{{}, {}, {}, {}},
				Level:           2,
				Value:           0,
			},
			rightOk: true,
			right: Node{
				Sequence:        "11",
				GapDepth:        4,
				Lengths:         []int{2, 0, 0, 0},
				Tails:           []int{1, -1, -1, -1},
				LastRecurrences: []int{1, -1, -1, -1},
				GapBuckets:      [][]int{{}, {}, {}, {}},
				Level:           2,
				Value:           1,
			},
		},
		{
			name: "sequence: 11",
			node: Node{
				Sequence:        "11",
				GapDepth:        4,
				Lengths:         []int{2, 0, 0, 0},
				Tails:           []int{1, -1, -1, -1},
				LastRecurrences: []int{1, -1, -1, -1},
				GapBuckets:      [][]int{{}, {}, {}, {}},
				Level:           2,
				Value:           1,
			},
			leftOk: true,
			left: Node{
				Sequence:        "110",
				GapDepth:        4,
				Lengths:         []int{3, 0, 0, 0},
				Tails:           []int{0, -1, -1, -1},
				LastRecurrences: []int{1, -1, -1, -1},
				GapBuckets:      [][]int{{}, {}, {}, {}},
				Level:           3,
				Value:           0,
			},
			rightOk: true,
			right: Node{
				Sequence:        "111",
				GapDepth:        4,
				Lengths:         []int{3, 1, 0, 0},
				Tails:           []int{1, 1, -1, -1},
				LastRecurrences: []int{2, -1, -1, -1},
				GapBuckets:      [][]int{{1}, {}, {}, {}},
				Level:           3,
				Value:           1,
			},
		},
		{
			name: "sequence: 11",
			node: Node{
				Sequence:        "11",
				GapDepth:        4,
				Lengths:         []int{2, 0, 0, 0},
				Tails:           []int{1, -1, -1, -1},
				LastRecurrences: []int{1, -1, -1, -1},
				GapBuckets:      [][]int{{}, {}, {}, {}},
				Level:           2,
				Value:           1,
			},
			leftOk: true,
			left: Node{
				Sequence:        "110",
				GapDepth:        4,
				Lengths:         []int{3, 0, 0, 0},
				Tails:           []int{0, -1, -1, -1},
				LastRecurrences: []int{1, -1, -1, -1},
				GapBuckets:      [][]int{{}, {}, {}, {}},
				Level:           3,
				Value:           0,
			},
			rightOk: true,
			right: Node{
				Sequence:        "111",
				GapDepth:        4,
				Lengths:         []int{3, 1, 0, 0},
				Tails:           []int{1, 1, -1, -1},
				LastRecurrences: []int{2, -1, -1, -1},
				GapBuckets:      [][]int{{1}, {}, {}, {}},
				Level:           3,
				Value:           1,
			},
		},
		{
			name: "sequence: 111",
			node: Node{
				Sequence:        "111",
				GapDepth:        4,
				Lengths:         []int{3, 1, 0, 0},
				Tails:           []int{1, 1, -1, -1},
				LastRecurrences: []int{2, -1, -1, -1},
				GapBuckets:      [][]int{{1}, {}, {}, {}},
				Level:           3,
				Value:           1,
			},
			leftOk: true,
			left: Node{
				Sequence:        "1110",
				GapDepth:        4,
				Lengths:         []int{4, 1, 0, 0},
				Tails:           []int{0, 1, -1, -1},
				LastRecurrences: []int{2, -1, -1, -1},
				GapBuckets:      [][]int{{1}, {}, {}, {}},
				Level:           4,
				Value:           0,
			},
			rightOk: true,
			right: Node{
				Sequence:        "1111",
				GapDepth:        4,
				Lengths:         []int{4, 2, 0, 0},
				Tails:           []int{1, 1, -1, -1},
				LastRecurrences: []int{3, 1, -1, -1},
				GapBuckets:      [][]int{{1}, {}, {}, {}},
				Level:           4,
				Value:           1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			left, leftOk := tt.node.Left()
			assert.Equal(t, tt.leftOk, leftOk, tt.name)
			assert.Equal(t, tt.left, left, tt.name)
			right, rightOk := tt.node.Right()
			assert.Equal(t, tt.rightOk, rightOk, tt.name)
			assert.Equal(t, tt.right, right, tt.name)
		})
	}
}
