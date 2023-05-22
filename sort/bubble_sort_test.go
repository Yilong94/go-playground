package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	testCases := []struct {
		name     string
		expected []int
		slice    []int
	}{
		{
			name:     "sort unsorted slice",
			expected: []int{1, 2, 3, 4, 5},
			slice:    []int{3, 1, 5, 2, 4},
		},
		{
			name:     "sort sorted slice",
			expected: []int{1, 2, 3, 4, 5},
			slice:    []int{1, 2, 3, 4, 5},
		},
		{
			name:     "sort slice of length 1",
			expected: []int{1},
			slice:    []int{1},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := BubbleSort(testCase.slice)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
