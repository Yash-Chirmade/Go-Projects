package solution

//  Longest Subarray With Sum K
// You are given an array 'a' of size 'n' and an integer 'k'.

// Find the length of the longest subarray of 'a' whose sum is equal to 'k'.

// Example :
// Input: ‘n’ = 7 ‘k’ = 3
// ‘a’ = [1, 2, 3, 1, 1, 1, 1]

// Output: 3
// Explanation: Subarrays whose sum = ‘3’ are:
// [1, 2], [3], [1, 1, 1] and [1, 1, 1]
// Here, the length of the longest subarray is 3, which is our final answer.

func longestSubarrayWithSumK(nums []int, requiredSum int) int {
	n := len(nums)
	var sum, longestSubarraylength int
	for i, _ := range nums {
		j := i
		sum = 0
		for j <= n-1 {
			sum += nums[j]
			if sum == requiredSum {
				if longestSubarraylength < j-i+1 {
					longestSubarraylength = j - i + 1
				}
				break
			} else if sum > requiredSum {
				break
			}
			j++
		}
	}
	return longestSubarraylength

}

// Prefix Sum Hashing approach 
func LongestSubarrayWithSum(a []int, k int) int {
	prefixSum := make(map[int]int)
	sum := 0
	maxLen := 0

	for i, num := range a {
		sum += num

		if sum == k {
			maxLen = i + 1
		}

		if _, found := prefixSum[sum-k]; found {
			maxLen = max(maxLen, i-prefixSum[sum-k])
		}

		if _, exists := prefixSum[sum]; !exists {
			prefixSum[sum] = i
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}