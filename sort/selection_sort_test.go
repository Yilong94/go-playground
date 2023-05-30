package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := SelectionSort(testCase.slice)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
