package main

// count conforming bitmasks
// given 3 30-bit unsigned integers, count the number of intergers that conform to atleast one of the given integers
func count(a, b, c int) int {
	var count int
	// use bitwise AND to find the number of bitmasks that are conforming to atleast one of the given integers
	for i := 0; i < 1<<30; i++ {
		if (a&i) == a || (b&i) == b || (c&i) == c {
			count++
		}
	}
	return count
}
