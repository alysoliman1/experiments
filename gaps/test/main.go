package main

import (
	"fmt"
	"math"

	"github.com/asoliman1/experiments/gaps/internal/pkg/utils"
)

// 28533
const p = 74353

func main() {
	s := int(math.Floor(p * 0.4999))
	E := utils.MakeInterval(s, p-s)
	k := 0
	for a := range p {
		H_a := utils.Multiply(a, E, p)
		for b := range a {
			H_b := utils.Multiply(b, E, p)
			overlap := utils.Overlap(H_a, H_b)
			if len(overlap) > k {
				k = len(overlap)
			}
			fmt.Println(a, b, s, p-s, k)
		}
	}
}
