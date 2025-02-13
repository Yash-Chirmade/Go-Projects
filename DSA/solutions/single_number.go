package solution

// 136. Single Number

func SingleNumber(nums []int) int {
	var xor int
    for _,num :=range nums{
		xor=xor^num
	}
	return xor
}