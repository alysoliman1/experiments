package main

import (
	"fmt"
	"math"
	"sync"

	"github.com/asoliman1/experiments/gaps/internal/pkg/nodes"
)

/*
0 2
1 4
2 8
3 16
4 32
5 64
6 128
7 256
8 512
9 1024
10 2048
11 4048
12 7928
13 15328
14 29276
15 55156
16 102956
17 190304
18 349316
19 637152
20 1157292
21 2094925
22 3783595
23 6822418
24 12290399

*/

func main() {
	numWorkers := 100
	maxTreeLevel := 25
	rootLayer := make(chan nodes.Node, 1)
	rootLayer <- nodes.NewRoot(10)
	close(rootLayer)
	layer := rootLayer
	counts := []int{}
	for range maxTreeLevel {
		layer = PopulateTreeLayer(layer, numWorkers)
		counts = append(counts, len(layer))
	}

	for i, count := range counts {
		fmt.Println(i+1, count/2, math.Pow(float64(i+1), 2)/2)
	}
}

func PopulateTreeLayer(
	previousLayerNodes <-chan nodes.Node,
	numWorkers int,
) chan nodes.Node {
	layerNodes := make(chan nodes.Node, 2*len(previousLayerNodes))
	defer close(layerNodes)
	wg := new(sync.WaitGroup)
	wg.Add(numWorkers)
	for range numWorkers {
		go func() {
			defer wg.Done()
			for node := range previousLayerNodes {
				if left, ok := node.Left(); ok {
					layerNodes <- left
				}
				if right, ok := node.Right(); ok {
					layerNodes <- right
				}
			}
		}()
	}
	wg.Wait()
	return layerNodes
}
