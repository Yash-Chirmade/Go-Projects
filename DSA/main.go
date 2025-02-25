package main

import (
	solution "dsa/prep/solutions"
	"fmt"
	"log"
	"reflect"
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

	// checkMissing := []int{0,1}
	// n:=solution.MissingNumber(checkMissing)
	// log.Printf("Missing Number is %v ", n)

	// conOnes := []int{1,0,1,1,0,1}
	// n:=solution.FindMaxConsecutiveOnes(conOnes)
	// log.Printf("Number of ones is %v ", n)

	// numbers := []int{0,1,1,1}
	// n:=solution.SingleNumber(numbers)
	// log.Printf("Single Number is %v ", n)

	// numbers := []int{1, 3, 1, 1, 2, 4, 7}
	// requiredSum := 7
	// n := solution.LongestSubarrayWithSum(numbers, requiredSum)
	// log.Printf("Longest subarray length is %v ", n)

	type testCase struct {
		n      int
		k      int
		a      []int
		expect int
	}
	tests := []testCase{
		{n: 7, k: 3, a: []int{1, 2, 3, 1, 1, 1, 1}, expect: 3},
		{n: 5, k: 15, a: []int{1, 2, 3, 4, 5}, expect: 5},
		{n: 4, k: 10, a: []int{1, 2, 3, 4}, expect: 4},
		{n: 6, k: 4, a: []int{1, 2, 4, 1, 3, 5}, expect: 2},
		{n: 9, k: 6, a: []int{1, 2, 3, 1, 1, 1, 2, 4, 2}, expect: 4},
		{n: 5, k: 0, a: []int{0, 0, 0, 0, 0}, expect: 1},
		{n: 6, k: 5, a: []int{2, -1, 3, 1, -2, 4}, expect: 5},
		{n: 10, k: 10, a: []int{1, 2, 3, 4, 5, 5, 0, 0, 10, 1}, expect: 4},
	}

	// Run each test case
	for i, test := range tests {
		result := solution.LongestSubarrayWithSum(test.a, test.k)
		if !reflect.DeepEqual(result, test.expect) {
			fmt.Printf("❌ Test %d failed: Input: %+v, Expected: %d, Got: %d\n", i+1, test.a, test.expect, result)
		} else {
			fmt.Printf("✅ Test %d passed!\n", i+1)
		}
	}

	// colors := []int{0,0,1,2,1}
	// solution.SortColors(colors)
	// log.Print(colors)

	Arr := []int{5,-4,0,7,8}
	maxSum:=solution.MaxSubArrayKadanes(Arr)
	log.Printf("Maximum sum is : %d" , maxSum)
}
