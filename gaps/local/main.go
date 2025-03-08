package main

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/asoliman1/experiments/gaps/internal/pkg/nodes"
)

func main() {
	numWorkers := 5
	maxTreeLevel := 20

	jobs := make(chan nodes.Node, 100000000)
	jobs <- nodes.NewRoot(10)

	var counter atomic.Int32

	wg := new(sync.WaitGroup)
	wg.Add(numWorkers)
	for range numWorkers {
		go func() {
			defer wg.Done()
			for node := range jobs {
				if node.Level == maxTreeLevel {
					counter.Add(1)
					continue
				}
				if left, ok := node.Left(); ok {
					jobs <- left
				}
				if right, ok := node.Right(); ok {
					jobs <- right
				}
			}
		}()
	}

	go func() {
		fmt.Println(counter.Load())
	}()
	wg.Wait()

	//raw, _ := json.MarshalIndent(sets, "", " ")
	//os.WriteFile("sets.json", raw, os.ModePerm)
}
