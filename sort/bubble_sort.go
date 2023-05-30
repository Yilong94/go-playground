package sort

// Bubble sort is named as such because the largest number will bubble up to the top

// Time complexity
// - Best case: 	O(N)
// - Worst case:	O(N^2)
// - Average case: 	O(N^2)
func BubbleSort(s []int) []int {
	sCopy := make([]int, len(s))
	copy(sCopy, s)

	// Loop until a passthrough has no swap
	for {
		noSwap := false
		for i := 0; i < len(sCopy)-1; i++ {
			// Swap when number on left is bigger than number on right
			if sCopy[i] > sCopy[i+1] {
				sCopy[i], sCopy[i+1] = sCopy[i+1], sCopy[i]
				noSwap = true
			}
		}
		if !noSwap {
			break
		}
	}
	return sCopy
}
