package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRecurringGaps(t *testing.T) {
	tests := []struct {
		name       string
		word       []int
		gaps       []int
		uniqueGaps []int
	}{
		{
			name: "no recurrences",
			word: []int{0, 1, 0, 1, 0, 1},
		},
		{
			name: "one recurrence",
			word: []int{0, 0, 1, 0, 1, 0, 1, 0},
		},
		{
			name:       "two recurrences",
			word:       []int{0, 0, 1, 0, 1, 1, 0, 1},
			gaps:       []int{4},
			uniqueGaps: []int{4},
		},
		{
			name:       "three recurrences",
			word:       []int{0, 0, 1, 0, 1, 1, 0, 0},
			gaps:       []int{4, 2},
			uniqueGaps: []int{2, 4},
		},
		{
			name:       "four recurrences",
			word:       []int{0, 0, 1, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
			gaps:       []int{4, 2, 5},
			uniqueGaps: []int{2, 4, 5},
		},
		{
			name:       "five recurrences",
			word:       []int{0, 0, 1, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0, 1, 1},
			gaps:       []int{4, 2, 5, 2},
			uniqueGaps: []int{2, 4, 5},
		},
	}
	for _, tt := range tests {
		gaps, uniqueGaps := GetRecurringGaps(tt.word)
		assert.Equal(t, tt.gaps, gaps)
		assert.Equal(t, tt.uniqueGaps, uniqueGaps, tt.name)
	}
}
