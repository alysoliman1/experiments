package main

import (
	"encoding/json"
	"os"

	"github.com/asoliman1/experiments/gaps/internal/pkg/nodes"
)

func main() {
	root := nodes.Node{
		Sequence:        "",
		GapDepth:        10,
		Lengths:         []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Tails:           []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		LastRecurrences: []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		GapBuckets:      [][]int{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}},
		Level:           0,
		Value:           -1,
	}

	N := 24

	queue := []nodes.Node{root}
	sets := []string{}
	for len(queue) > 0 {
		next := queue[0]
		if len(queue) > 0 {
			queue = queue[1:]
		}
		if next.Level == N {
			sets = append(sets, next.Sequence)
			continue
		}
		if left, ok := next.Left(); ok {
			queue = append(queue, left)
		}
		if right, ok := next.Right(); ok {
			queue = append(queue, right)
		}
	}

	raw, _ := json.MarshalIndent(sets, "", " ")
	os.WriteFile("sets.json", raw, os.ModePerm)
}
