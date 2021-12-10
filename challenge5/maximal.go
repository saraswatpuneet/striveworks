package main

// maximal sum of non-adjacent elements
func MaxSliceSum(A []int) int {
	if len(A) == 0 {
		return 0
	}
	if len(A) == 1 {
		return A[0]
	}
	if len(A) == 2 {
		return max(A[0], A[1])
	}

	// kadane's algorithm 
	// local max array going forward
	max_forward := make([]int, len(A))
	// initialize to 0
	for i := 1; i < len(A)-1; i++ {
		max_forward[i] = max(max_forward[i-1]+A[i], 0)

	}
	// local max array going backward
	max_backward := make([]int, len(A))
	for i := len(A) - 2; i >= 0; i-- {
		max_backward[i] = max(max_backward[i+1]+A[i], 0)
	}
	// find max sum
	max_sum := 0
	for i := 1; i < len(A)-1; i++ {
		max_sum = max(max_sum, max_forward[i-1]+max_backward[i+1])
	}	
	return max_sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
