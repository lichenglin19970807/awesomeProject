package test

import (
	"awesomeProject/array"
	"github.com/stretchr/testify/assert"
	"math/rand"
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
