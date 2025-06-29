package array_study

func GenerateMatrix(n int) [][]int {
	nums := make([][]int, n)
	for i := range nums {
		nums[i] = make([]int, n)
	}
	startX, startY := 0, 0
	offset, count := 1, 1
	for r := 0; r < n/2; r++ {
		for j := startX; j < n-offset; j++ {
			nums[startX][j] = count
			count++
		}
		for i := startY; i < n-offset; i++ {
			nums[i][n-offset] = count
			count++
		}
		for j := n - offset; j > startY; j-- {
			nums[n-offset][j] = count
			count++
		}
		for i := n - offset; i > startX; i-- {
			nums[i][startY] = count
			count++
		}
		startX++
		startY++
		offset++
	}
	if n%2 != 0 {
		nums[n/2][n/2] = count
	}
	return nums
}
