package solution

func MissingNumberHashing(nums []int) int {
	freqarr := make([]int, len(nums)+1)
	for _, num := range nums {
		freqarr[num]++
	}
	for i, v := range freqarr {
		if v == 0 {
			return i
		}
	}
	return nums[0]
}
func MissingNumberBruteForce(nums []int) int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}
	for i := 0; i <= len(nums); i++ {
		if _, ok := freqMap[i]; !ok {
			return i
		}
	}
	return nums[0]
}
func MissingNumberMathematics(nums []int) int {
	var sum int
	n := len(nums)
	for _, num := range nums {
		sum += num
	}
	
	return (n * (n+1))/2 - sum
}
func MissingNumber(nums []int) int {
	var xor1,xor2 int
	for i, num := range nums {
		xor2=xor2^num
		xor1=xor1^i
	}
	xor1=xor1^(len(nums))
	return xor2^xor1
}
