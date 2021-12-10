package main

import "sort"

// given an array A of intergers, find the smallest positive integer that is not in A
// A will have at least 1 positive integer
func smallestMissingPositive(A []int) int {
	n := len(A)
	if n < 1 {
		return 0
	}
	if n == 1 {
		if A[0] == 1 {
			return 2
		} else {
			return 1
		}
	}
	// check if 1 does not exist in A then return 1
	if !contains(A, 1) {
		return 1
	}
	for i := 0; i < n; i++ {
		for A[i] > 0 && A[i] <= n && A[i] != A[A[i]-1] {
			A[i], A[A[i]-1] = A[A[i]-1], A[i]
		}
	}
	for i := 0; i < n; i++ {
		if A[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

func contains(A []int, n int) bool {
	for _, v := range A {
		if v == n {
			return true
		}
	}
	return false
}

// approach 2 - use sorting
func smallestMissingPositive2(A []int) int {
	n := len(A)
	if n < 1 {
		return 0
	}
	if n == 1 {
		if A[0] == 1 {
			return 2
		} else {
			return 1
		}
	}
	// check if 1 does not exist in A then return 1
	if !contains(A, 1) {
		return 1
	}
	sort.Ints(A)
	var result int = 1
	for i := 0; i < n; i++ {
		if A[i] == result {
			result++
		}
	}
	return result
}
