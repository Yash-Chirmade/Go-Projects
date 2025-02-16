package solution

// 75. Sort Colors||
// Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.
// We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.
// You must solve this problem without using the library's sort function.

// Example :
// Input: nums = [2,0,2,1,1,0]
// Output: [0,0,1,1,2,2]

func SortColors(nums []int) {

	for i, _ := range nums {
		j := len(nums) - 1
		minIdx := len(nums) - 1
		minElement := nums[len(nums) - 1]
		for j >= i {
			if nums[j] < minElement {
				minElement = nums[j]
				minIdx = j
			}
			j--
		}

		temp := nums[minIdx]
		nums[minIdx] = nums[i]
		nums[i] = temp
	}

}
