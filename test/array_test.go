package test

import (
	"awesomeProject/array"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

func InitialIncrementalArray(maxLen, interval int) []int {
	nums := make([]int, 0)
	offset := 0
	for i := 0; i < maxLen; i++ {
		nums = append(nums, rand.Intn(interval)+offset)
		offset += interval
	}
	return nums
}

func InitialNonStrictlyIncreasingArray(maxLen, interval int) []int {
	nums := make([]int, maxLen)
	for i := 0; i < maxLen; i++ {
		nums[i] = rand.Intn(interval)
	}
	sort.Ints(nums)
	return nums
}

func TestSortedSquares(t *testing.T) {
	nums := []int{-4, -1, 0, 3, 10}
	output := []int{0, 1, 9, 16, 100}
	assert.Equal(t, output, array_study.SortedSquares(nums))

	numsB := []int{-7, -3, 2, 3, 11}
	outputB := []int{4, 9, 9, 49, 121}
	assert.Equal(t, outputB, array_study.SortedSquares(numsB))
}

func TestRemoveElement(t *testing.T) {
	nums := InitialNonStrictlyIncreasingArray(10, 5)
	fmt.Println(nums)
	fmt.Println(array_study.RemoveElement(nums, 2))
}

func TestRemoveDuplicates(t *testing.T) {
	nums := InitialNonStrictlyIncreasingArray(100, 3)
	fmt.Println(nums)
	assert.Equal(t, 3, array_study.RemoveDuplicates(nums))
}

func Test_BinarySearch(t *testing.T) {
	assertHandler := assert.New(t)
	nums := InitialIncrementalArray(100, 10)
	indexA := rand.Intn(100)
	assertHandler.Equal(indexA, array_study.Search(nums, nums[indexA]))
	assertHandler.Equal(indexA, array_study.SearchV2(nums, nums[indexA]))
	targetB := -999
	assertHandler.Equal(-1, array_study.Search(nums, targetB))
	assertHandler.Equal(-1, array_study.SearchV2(nums, targetB))
}

func Benchmark_BinarySearch(b *testing.B) {
	nums := InitialIncrementalArray(100000000, 10)
	b.ResetTimer()
	indexA := rand.Intn(100000000)
	array_study.Search(nums, nums[indexA])
	b.StopTimer()
}

func Benchmark_BinarySearchV2(b *testing.B) {
	nums := InitialIncrementalArray(100000000, 10)
	b.ResetTimer()
	indexA := rand.Intn(100000000)
	array_study.Search(nums, nums[indexA])
	b.StopTimer()
}
