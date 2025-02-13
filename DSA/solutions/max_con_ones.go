package solution

// 485. Max Consecutive Ones
func FindMaxConsecutiveOnes(nums []int) int {
	var counter, lastCount int

	for _, num := range nums {
		if num != 1 {
			if counter > lastCount {
				lastCount = counter
			}
			counter = 0
			continue
		}
		counter++
	}
	return max(counter,lastCount)
}
