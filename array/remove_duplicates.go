package array_study

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	fast, slow := 1, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
