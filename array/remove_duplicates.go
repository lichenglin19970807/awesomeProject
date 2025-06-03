package array_study

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	fast, slow := 1, 1
	for ; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
