package array_study

import "math"

func MinSubArrayLen(target int, nums []int) int {
	sum, result := 0, math.MaxInt
	hasResult := false
	for fast, slow := 0, 0; fast < len(nums); fast++ {
		sum += nums[fast]
		for sum >= target && slow <= fast {
			l := fast - slow + 1
			hasResult = true
			result = min(result, l)
			sum -= nums[slow]
			slow++
		}
	}
	if !hasResult {
		return 0
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
