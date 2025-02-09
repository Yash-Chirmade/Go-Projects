package solution

import (
	"strings"
)

const (
	ROW1 = "qwertyuiop"
	ROW2 = "asdfghjkl"
	ROW3 = "zxcvbnm"
)

// 500. Keyboard Row

// Given an array of strings words, return the words that can be typed using letters of the alphabet on only one row of American keyboard like the image below.

// Note that the strings are case-insensitive, both lowercased and uppercased of the same letter are treated as if they are at the same row.

// In the American keyboard:

// the first row consists of the characters "qwertyuiop",
// the second row consists of the characters "asdfghjkl", and
// the third row consists of the characters "zxcvbnm".

// Example 1:

// Input: words = ["Hello","Alaska","Dad","Peace"]

// Output: ["Alaska","Dad"]

// Explanation:

// Both "a" and "A" are in the 2nd row of the American keyboard due to case insensitivity.

func FindWords(words []string) []string {
	var result []string
	for _, word := range words {
		if wordFromSameRow := iswordFromSameRow(word); wordFromSameRow {
			result = append(result, word)
		}
	}
	return result
}

func iswordFromSameRow(word string) bool {
	var lenCounter int

	for _, letter := range strings.Split(word, "") {
		if !strings.Contains(ROW1, strings.ToLower(letter)) {
			break
		}
		if strings.Contains(ROW1, strings.ToLower(letter)) {
			lenCounter++
		}
	}
	if lenCounter != len(word) {
		lenCounter = 0
		for _, letter := range strings.Split(word, "") {
			if !strings.Contains(ROW2, strings.ToLower(letter)) {
				break
			}
			if strings.Contains(ROW2, strings.ToLower(letter)) {
				lenCounter++
			}
		}
	}
	if lenCounter != len(word) {
		lenCounter = 0
		for _, letter := range strings.Split(word, "") {
			if !strings.Contains(ROW3, strings.ToLower(letter)) {
				break
			}
			if strings.Contains(ROW3, strings.ToLower(letter)) {
				lenCounter++
			}
		}
	}
	if lenCounter == len(word) {
		return true
	}
	return false
}
