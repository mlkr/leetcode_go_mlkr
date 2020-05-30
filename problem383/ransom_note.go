package problem383

import "strings"

func canConstruct(ransomNote string, magazine string) bool {
	charCount := make(map[rune]int,0)
	for _, char := range ransomNote{
		if charCount[char] == 0 {
			charCount[char] = strings.Count(ransomNote, string(char))
		}
	}

	for char := range charCount {
		if charCount[char] > strings.Count(magazine, string(char)) {
			return false
		}
	}

	return true
}
