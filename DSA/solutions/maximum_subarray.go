package solution

// 53. Maximum Subarray
// Given an integer array nums, find the
// subarray
//  with the largest sum, and return its sum.

// Example 1:

// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: The subarray [4,-1,2,1] has the largest sum 6.

func MaxSubArray(nums []int) int {
	maxCurrent, maxGlobal := nums[0], nums[0]
	for i := 1; i <= len(nums)-1; i++ {
		maxCurrent = max(nums[i], maxCurrent+nums[i])
		if maxGlobal < maxCurrent {
			maxGlobal = maxCurrent
		}
	}
	return maxGlobal
}
