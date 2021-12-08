package main

import (
	"fmt"
	"strings"
)

// given a string with letter A, B and C apply transformtion to replace one occurence of AA,BB and CC with empty string
func StringTransformation(s string) string {
	var aaApplied, bbApplied, ccApplied bool = true, true, true
	// while aaApplied, bbApplied and ccApplied are true
	for aaApplied || bbApplied || ccApplied {
		length := len(s)
		if length < 2 {
			break
		}
		aaApplied = false
		s = strings.Replace(s, "AA", "", 1)
		if length-len(s) == 2 {
			length = length - 2
			aaApplied = true
		}
		bbApplied = false
		s = strings.Replace(s, "BB", "", 1)
		if length-len(s) == 2 {
			length = length - 2
			bbApplied = true
		}
		ccApplied = false
		s = strings.Replace(s, "CC", "", 1)
		if length-len(s) == 2 {
			ccApplied = true
		}
	}
	return s
}

func main() {
	s := "ACCAABBC"
	s = StringTransformation(s)
	fmt.Println(s)
}
