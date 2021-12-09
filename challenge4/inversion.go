package main

// write a function to find total number of inversions in an array A
// an inversion is defined as the pair of elements in an array
// where the first element is greater than the second element
//
// example:
// A = [-1,6 3, 4, 7, 4]
// inversions = 4
// (1,2) (1,3) (1, 5) (4, 5)
func FindTotalInversionsInArray(A []int) int {
	if len(A) == 0 {
		return 0
	}
	count:= findInversions(A, 0, len(A)-1)
	if count > 1000000000 {
		return -1
	}
}

func findInversions(A []int, start, end int) int {
	if start == end {
		return 0
	}
	mid := (start + end) / 2
	return findInversions(A, start, mid) + findInversions(A, mid+1, end) + mergeAndCount(A, start, mid, end)
}

func mergeAndCount(A []int, start, mid, end int) int {
	var left []int
	var right []int
	var count int
	for i := start; i <= mid; i++ {
		left = append(left, A[i])
	}
	for i := mid + 1; i <= end; i++ {
		right = append(right, A[i])
	}
	i, j := 0, 0
	for k := start; k <= end; k++ {
		if i < len(left) && j < len(right) {
			if left[i] <= right[j] {
				A[k] = left[i]
				i++
			} else {
				A[k] = right[j]
				j++
				count += len(left) - i
			}
		} else if i < len(left) {
			A[k] = left[i]
			i++
		} else {
			A[k] = right[j]
			j++
		}
	}
	return count
}

func main() {
	A := []int{-1, 6, 3, 4, 7, 4}
	println(FindTotalInversionsInArray(A))
}