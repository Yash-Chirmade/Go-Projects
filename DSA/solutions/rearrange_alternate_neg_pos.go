package solution

// 2149. Rearrange Array Elements by Sign
// You are given a 0-indexed integer array nums of even length consisting of an equal number of positive and negative integers.

// You should return the array of nums such that the the array follows the given conditions:

// Every consecutive pair of integers have opposite signs.
// For all integers with the same sign, the order in which they were present in nums is preserved.
// The rearranged array begins with a positive integer.
// Return the modified array after rearranging the elements to satisfy the aforementioned conditions.
//{3,1,-2,-5,2,-4}

func RearrangeArray(nums []int) []int {
	negIdx := 1
	posIdx := 0
	rearrangedArr := make([]int, len(nums))
	for _, num := range nums {
		if num < 0 {
			rearrangedArr[negIdx] = num
			negIdx += 2
		} else {
			rearrangedArr[posIdx] = num
			posIdx += 2
		}

	}

	return rearrangedArr
}
