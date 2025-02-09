package main

import (
	solution "dsa/prep/solutions"
	"log"
)

func main() {

	// wordsNext := []string{"a", "aba", "ababa", "aa"}
	// wordsNext := []string{"bb","bb"}
	// count := solution.CountPrefixSuffix(wordsNext)
	// log.Printf("Count is %d", count)

	wordsNext := []string{"a","b"}
	count := solution.FindWords(wordsNext)
	log.Printf("array result is %v", count)

}
