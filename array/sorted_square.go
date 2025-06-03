package array_study

func SortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	left, right := 0, len(nums)-1
	for i := len(nums) - 1; i >= 0; i-- {
		leftPow := pow(nums[left])
		rightPow := pow(nums[right])
		if leftPow >= rightPow {
			res[i] = leftPow
			left++
		} else {
			res[i] = rightPow
			right--
		}
	}
	return res
}

func pow(num int) int {
	return num * num
}
