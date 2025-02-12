package solution

import "fmt"

// 283. Move Zeroes
// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Note that you must do this in-place without making a copy of the array.

func MoveZeroes(nums []int) {
	fmt.Println(nums)
	j := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			j = i
			break
		}
	}
	if j == -1 {
		return
	}
	for i := j + 1; i < len(nums); i++ {
		var temp int
		if nums[i] != 0 {
			temp = nums[j]
			nums[j] = nums[i]
			nums[i] = temp
			j++
		}
	}
	fmt.Print(nums)
}
