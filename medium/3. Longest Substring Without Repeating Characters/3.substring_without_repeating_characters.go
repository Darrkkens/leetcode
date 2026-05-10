package main

func lengthOfLongestSubstring(s string) int {
	lastSeen := make(map[byte]int)

	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		char := s[right]

		if lastIndex, exists := lastSeen[char]; exists && lastIndex >= left {
			left = lastIndex + 1
		}

		lastSeen[char] = right

		currentLen := right - left + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}
