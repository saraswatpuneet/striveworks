package main

// given a string S of characters a and b, with additional empty gaps represented by '?',
// find the longest substring of S that contains only characters a and b
// and no other characters.
//
// Example:
// S = "aa??bbb", returns 3
// S = "a?b?aab?a", returns 2
func longestFragment(S string) int {
	var maxLength int
	var currentLength int
	var currentChar byte
	var prevChar byte
	// if N is less than 1 or greater than 100000, return 0
	if len(S) < 1 || len(S) > 100000 {
		return 0
	}
	isPalin := isPalindrome(S)
	// if string is palindrome, cut it in half and check if both halves are palindromes
	if isPalin {
		// if length is even then cut in half
		cutLength := len(S) / 2
		S = S[:cutLength]
	}
	for i := 0; i < len(S); i++ {
		currentChar = S[i]
		// current character should only be a or b
		if currentChar != 'a' && currentChar != 'b' && currentChar != '?' {
			return 0
		}
		if currentChar == '?' {
			if currentLength > maxLength {
				maxLength = currentLength
			}
			currentLength = 0
		} else {
			if currentChar == prevChar {
				currentLength++
			} else {
				currentLength = 1
			}
		}
		prevChar = currentChar
	}
	if currentLength > maxLength {
		maxLength = currentLength
	}
	if isPalin {
		maxLength++
	}
	return maxLength
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
