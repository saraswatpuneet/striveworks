package main

import "strconv"

// given an integer A, return shuffled representation of A by altenating one digit from start and end of the number
// for example:
func CodilityShuffleInteger(A int) int {
	if A < 0 {
		return -1
	}
	if A == 0 {
		return 0
	}
	// convert to string
	strA := strconv.Itoa(A)
	shuffledA := shuffle(strA)
	// convert back to int
	intA, err := strconv.Atoi(shuffledA)
	if err != nil {
		panic(err)
	}
	return intA
}

func shuffle(str string) string {
	if len(str) == 1 {
		return str
	}
	if str ==""	{
		return ""
	}
	// get the first and last character
	firstChar := str[0]
	lastChar := str[len(str)-1]
	// join the first and last character and call the function again
	return string(firstChar) +  string(lastChar) + shuffle(str[1:len(str)-1])
}
