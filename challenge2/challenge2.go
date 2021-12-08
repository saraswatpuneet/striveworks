package main

// write a function give two integers A and B
// return a string that contains letters A and B with no three consecutive letters
// example:
// A = 1, B = 2
// return "abb" or "bab"
// A = 1, B = 3
// return babb
// A = 2, B = 3
// return babab
func Solution(A int, B int) string {
	var result string
	var writeA bool
	// while A or B is not zero
	for A > 0 || B > 0 {
		// if result length is greater than 2
		if len(result) > 1 && result[len(result)-2] == result[len(result)-1] {
			// if last letter is B and writeA is true else writeA is false
			if result[len(result)-1] == 'b' {
				writeA = true
			} else {
				writeA = false
			}
		} else {
			// if A is greater than B
			if A >= B {
				writeA = true
			} else {
				writeA = false
			}
		}
		if writeA {
			result += "a"
			// decrement A
			A--
		} else {
			result += "b"
			// decrement B
			B--
		}
	}

	return result
}

func main() {
	println(Solution(5, 3))
}
