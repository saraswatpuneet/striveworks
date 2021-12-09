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

// given a string with letter A, B and C apply transformtion to replace one occurence of AA,BB and CC with empty string
func StringTransformation2(s string) string {
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

// use a buffer to store the result
// use a loop to replace the string
func StringTransformation3(s string) string {
	var aaApplied, bbApplied, ccApplied bool = true, true, true

	for aaApplied || bbApplied || ccApplied {
		length := len(s)
		if length < 2 {
			break
		}
		aaApplied = false
		bbApplied = false
		ccApplied = false
		for i := 0; i < len(s); i++ {
			if i+1 < len(s) && s[i] == 'A' && s[i+1] == 'A' {
				if i+2 < len(s) {
					s = s[:i] + s[i+2:]
				} else {
					s = s[:i]
				}
				aaApplied = true
				continue
			}
			if i+1 < len(s) && s[i] == 'B' && s[i+1] == 'B' {
				if i+2 < len(s) {
					s = s[:i] + s[i+2:]
				} else {
					s = s[:i]
				}
				bbApplied = true
				continue
			}
			if i+1 < len(s) && s[i] == 'C' && s[i+1] == 'C' {
				if i+2 < len(s) {
					s = s[:i] + s[i+2:]
				} else {
					s = s[:i]
				}
				ccApplied = true
				continue
			}
		}
	}

	return s
}
func main() {
	s := "ABCBBCBA"
	s = StringTransformation3(s)
	fmt.Println(s)
}
