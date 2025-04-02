package utils

import (
	"fmt"
	"slices"
)

func MakeWord(set []int, p int) (word []int) {
	for i := range p {
		if slices.Contains(set, i) {
			word = append(word, 1)
		} else {
			word = append(word, 0)
		}
	}
	return
}

func LegalWord(word []int) (int, bool) {
	depth := 10
	currentWord, uniqueGaps := GetRecurringGaps(word)
	if len(uniqueGaps) > 3 {
		fmt.Println("failed 3 gap condition with gaps", uniqueGaps)
		return -1, false
	}

	/*
		if len(uniqueGaps) == 3 && !haveThreeSumCondition(
			uniqueGaps[0],
			uniqueGaps[1],
			uniqueGaps[2],
		) {
			fmt.Println("failed 3 sum condition", uniqueGaps)
			return false
		}
	*/
	maxDepth := 0
	for i := range depth {
		currentWord, uniqueGaps = GetRecurringGaps(currentWord)
		if len(uniqueGaps) > 4 {
			fmt.Println("failed 4 gap condition at level", i+1, "with gaps", uniqueGaps)
			return -1, false
		}
		if len(currentWord) == 0 {
			maxDepth = i
			break
		}
	}
	return maxDepth, true
}

func GetRecurringGaps(word []int) (gaps []int, uniqueGaps []int) {
	return getGaps(word, func(index int, word []int) bool {
		if index >= 0 && index+1 <= len(word)-1 {
			return word[index] == word[index+1]
		}
		return false
	})
}

func Overlap(A, B []int) (overlap []int) {
	for _, a := range A {
		if slices.Contains(B, a) {
			overlap = append(overlap, a)
		}
	}
	return
}

// getGaps
func getGaps(
	word []int,
	flag func(index int, word []int) bool,
) (gaps []int, uniqueGaps []int) {
	lastFlaggedIndex := -1
	setOfUniqueGaps := map[int]struct{}{}
	for i := range len(word) {
		if flag(i, word) {
			if lastFlaggedIndex > -1 {
				gap := i - lastFlaggedIndex
				gaps = append(gaps, gap)
				setOfUniqueGaps[gap] = struct{}{}
			}
			lastFlaggedIndex = i
		}
	}
	for gap := range setOfUniqueGaps {
		uniqueGaps = append(uniqueGaps, gap)
	}
	slices.Sort(uniqueGaps)
	return
}

func MakeInterval(a, b int) (interval []int) {
	for i := a; i <= b; i++ {
		interval = append(interval, i)
	}
	return
}

func Sum(set []int, p int) (sum int) {
	for _, x := range set {
		sum = (sum + x) % p
	}
	return
}

func Multiply(a int, set []int, p int) []int {
	output := []int{}
	for _, x := range set {
		output = append(output, (a*x)%p)
	}
	slices.Sort(output)
	return output
}
