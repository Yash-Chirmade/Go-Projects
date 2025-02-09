package solution

import (
	"log"
	"slices"
)

// 1. Two Sum
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

func TwoSum(nums []int, target int) []int {
	log.Printf("Nums before sorting : %v", nums)
	sorted := append([]int{}, nums...)
	slices.Sort(sorted)
	log.Printf("Nums after sorting: %v", nums)
	i := 0
	j := len(sorted) - 1
	var firstIdx, secondIdx int

	for i < j {
		if sorted[i]+sorted[j] == target {
			firstIdx = slices.Index(nums, sorted[i])
            secondIdx = slices.Index(nums, sorted[j])
			if sorted[i] == sorted[j] {
				for i = firstIdx; i < len(nums); i++ {
					if sorted[j] == nums[i] {
						secondIdx = i
					}
				}
			}
			return []int{firstIdx, secondIdx}
		} else if sorted[i]+sorted[j] > target {
			j--
			log.Printf("Decrementing j to %d", j)
		} else {
			i++
			log.Printf("Incrementing i to %d", i)
		}

	}
	return []int{}
}
