package solution

// 169. Majority Element

func MajorityElement(nums []int) int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}
	for k, v := range freqMap {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}

func SecondLargestOfArray(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	secondMax := nums[0]
	for _, num := range nums {
		if num > secondMax && num < max {
			secondMax = num
		}
	}
	return secondMax
}
