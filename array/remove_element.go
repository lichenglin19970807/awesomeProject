package array_study

func RemoveElement(nums []int, val int) int {
	fast, slow := 0, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
