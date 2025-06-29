package array_study

func SearchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return findInterval(nums, mid)
		}
	}
	return res
}

func findInterval(nums []int, mid int) []int {
	res := []int{mid, mid}
	left, right := mid-1, mid+1
	for left >= 0 && nums[left] == nums[mid] {
		res[0] = left
		left--
	}
	for right < len(nums) && nums[right] == nums[mid] {
		res[1] = right
		right++
	}
	return res
}

func SearchInsertBinary(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		middle := left + (right-left)/2
		if nums[middle] > target {
			right = middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}
	return left
}

func Search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}

func SearchV2(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		middle := left + (right-left)/2
		if nums[middle] > target {
			right = middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}
