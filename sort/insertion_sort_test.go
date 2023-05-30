package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := InsertionSort(testCase.slice)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
