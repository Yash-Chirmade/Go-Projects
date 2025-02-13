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

	// wordsNext := []string{"a","b"}
	// count := solution.FindWords(wordsNext)
	// log.Printf("array result is %v", count)

	// nums := []int{0,0,1,1,1,2,2,3,3,4}
	// target:= 1
	// // result := solution.TwoSum(nums,target)
	// k := solution.RemoveElement(nums,target)

	// nums := []int{0,1,4,5,3,6,6,23}
	// k := solution.MajorityElement(nums)
	// log.Printf("Result is %v ", k)

	// s := solution.SecondLargestOfArray(nums)
	// log.Printf("Second Max is %v ", s)

	// numbers := []int{1, 0, 3, 4, 0, 6, 0}
	// // k := 3
	// solution.MoveZeroes(numbers)
	// log.Printf("Rotated Array is %v ", arr)

	checkMissing := []int{0,1}
	n:=solution.MissingNumber(checkMissing)
	log.Printf("Missing Number is %v ", n)

}
