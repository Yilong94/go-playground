package sort

// Time complexity
// - Best case: 	O(N^2)
// - Average case: 	O(N^2)
// - Worst case: 	O(N^2)
func SelectionSort(s []int) []int {
	sCopy := make([]int, len(s))
	copy(sCopy, s)

	numPassthroughs := len(sCopy)
	for startIndex := 0; startIndex < numPassthroughs; startIndex++ {
		lowestIndex := startIndex

		// Find lowest value
		for i := startIndex + 1; i < len(sCopy); i++ {
			if sCopy[i] < sCopy[lowestIndex] {
				lowestIndex = i
			}
		}

		// Swap lowest value with first value in passthrough
		sCopy[startIndex], sCopy[lowestIndex] = sCopy[lowestIndex], sCopy[startIndex]
	}
	return sCopy
}
