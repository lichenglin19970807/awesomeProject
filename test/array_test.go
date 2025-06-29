package test

import (
	"awesomeProject/array"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

func TestBackStringCompare(t *testing.T) {
	input := "ab#c"
	fmt.Println(array_study.SubCompare(input))
}

func TestString(T *testing.T) {
	a := "abcd#"
	for _, v := range a {
		if v == '#' {
			fmt.Println("dafa")
		}
	}
	fmt.Println(a[:0])
}

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

func TestSearchRange(t *testing.T) {
	numsA := []int{5, 7, 7, 8, 8, 10}
	targetA := 8
	outputA := []int{3, 4}
	assert.Equal(t, outputA, array_study.SearchRange(numsA, targetA))

	numsB := []int{5, 7, 7, 8, 8, 10}
	targetB := 6
	outputB := []int{-1, -1}
	assert.Equal(t, outputB, array_study.SearchRange(numsB, targetB))

	numsC := []int{}
	targetC := 0
	outputC := []int{-1, -1}
	assert.Equal(t, outputC, array_study.SearchRange(numsC, targetC))
}

func TestSearchInsertBinary(t *testing.T) {
	numsA := []int{1, 3, 5, 6}
	targetA := 5
	outputA := 2
	assert.Equal(t, outputA, array_study.SearchInsertBinary(numsA, targetA))

	numsB := []int{1, 3, 5, 6}
	targetB := 2
	outputB := 1
	assert.Equal(t, outputB, array_study.SearchInsertBinary(numsB, targetB))

	numsC := []int{1, 3, 5, 6}
	targetC := 7
	outputC := 4
	assert.Equal(t, outputC, array_study.SearchInsertBinary(numsC, targetC))
}

func TestGenerateMatrix(t *testing.T) {
	n := 4
	nums := array_study.GenerateMatrix(n)
	for i := range nums {
		fmt.Println(nums[i])
	}
}

func TestMinSubArrayLen(t *testing.T) {
	inputTargetA := 7
	inputNumsA := []int{2, 3, 1, 2, 4, 3}
	outputA := 2
	assert.Equal(t, outputA, array_study.MinSubArrayLen(inputTargetA, inputNumsA))

	inputTargetB := 4
	inputNumsB := []int{1, 4, 4}
	outputB := 1
	assert.Equal(t, outputB, array_study.MinSubArrayLen(inputTargetB, inputNumsB))

	inputTargetC := 11
	inputNumsC := []int{1, 1, 1, 1, 1, 1, 1, 1}
	outputC := 0
	assert.Equal(t, outputC, array_study.MinSubArrayLen(inputTargetC, inputNumsC))
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
