package solution

func Rotate(nums []int, k int) []int {
	n:=len(nums)
	rotatedArray := make([]int, n)
	for i, num := range nums {
		if i <= k {
			rotatedArray[(n-1)-(k-i)] = num
		} else {
			rotatedArray[i-k-1] = num
		}
	}
	return rotatedArray
}
