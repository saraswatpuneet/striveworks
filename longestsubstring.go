package main
///
// find longest substring with repeating characters
// special case: given string is a palindrome
func longestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	if isPalindrome(s) {
		return len(s)
	}
	max := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if isPalindrome(s[i:j]) && j-i+1 > max {
				max = j - i + 1
			}
		}
	}
	return max
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
