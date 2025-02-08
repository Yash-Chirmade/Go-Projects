package solution

// 3042. Count Prefix and Suffix Pairs I

// You are given a 0-indexed string array words.
// Let's define a boolean function isPrefixAndSuffix that takes two strings, str1 and str2:
// isPrefixAndSuffix(str1, str2) returns true if str1 is both a  prefix and a  suffix of str2, and false otherwise.
// For example, isPrefixAndSuffix("aba", "ababa") is true because "aba" is a prefix of "ababa" and also a suffix, but isPrefixAndSuffix("abc", "abcd") is false.
// Return an integer denoting the number of index pairs (i, j) such that i < j, and isPrefixAndSuffix(words[i], words[j]) is true.

// Example 1:

// Input: words = ["a","aba","ababa","aa"]
// Output: 4
// Explanation: In this example, the counted index pairs are:
// i = 0 and j = 1 because isPrefixAndSuffix("a", "aba") is true.
// i = 0 and j = 2 because isPrefixAndSuffix("a", "ababa") is true.
// i = 0 and j = 3 because isPrefixAndSuffix("a", "aa") is true.
// i = 1 and j = 2 because isPrefixAndSuffix("aba", "ababa") is true.
// Therefore, the answer is 4.

func CountPrefixSuffix(words []string) int {
	var Counter int

	for i, word := range words {
		for j := (len(words) - 1); j > i; j-- {
			if IsPrefixAndSuffix(word, words[j]) {
				Counter++
			}
		}
	}
	return Counter
}

func IsPrefixAndSuffix(p, word string) bool {
	if len(p) > len(word) {
		return false
	}
	//check equality
	if word[(len(word)-len(p)):] == p && word[:len(p)] == p {
		return true
	}

	return false
}
