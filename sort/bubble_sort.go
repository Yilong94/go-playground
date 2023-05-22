package sort

// Time complexity
// - Best case: 	O(N)
// - Worst case:	O(N^2)
// - Average case: 	O(N^2)
func BubbleSort(s []int) []int {
	sCopy := make([]int, len(s))
	copy(sCopy, s)
	for {
		noSwap := false
		for i := 0; i < len(sCopy)-1; i++ {
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
