package main

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start := 0
	end := 0

	for i := 0; i < len(s); i++ {
		len1 := expandFromMiddle(s, i, i)   // palíndromo ímpar
		len2 := expandFromMiddle(s, i, i+1) // palíndromo par

		length := len1
		if len2 > len1 {
			length = len2
		}

		if length > end-start {
			start = i - (length-1)/2
			end = i + length/2
		}
	}

	return s[start : end+1]
}

func expandFromMiddle(s string, left int, right int) int {
	if len(s) == 0 || left > right {
		return 0
	}

	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return right - left - 1
}
