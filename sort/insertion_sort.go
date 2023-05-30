package sort

// Time complexity
// - Best case: 	O(N)
// - Worst case:	O(N^2)
// - Average case: 	O(N^2)
func InsertionSort(s []int) []int {
	sCopy := make([]int, len(s))
	copy(sCopy, s)

	for startIndex := 1; startIndex < len(sCopy); startIndex++ {
		temp := sCopy[startIndex] // Store temp value
		gapIndex := startIndex    // and create a gap

		for i := gapIndex - 1; i >= 0; i-- {
			if sCopy[i] > temp {
				sCopy[gapIndex] = sCopy[i] // Move number to fill the gap
				gapIndex--                 // Move gap down by one index
			} else {
				break
			}
		}
		sCopy[gapIndex] = temp // Move temp value into gap
	}
	return sCopy
}
