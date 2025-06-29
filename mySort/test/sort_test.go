package test

import (
	"awesomeProject/mySort"
	"awesomeProject/validator"
	"fmt"
	"sort"
	"testing"
)

func TestSanShaSort(t *testing.T) {
	a := validator.NewRandomArray(10, 1000)
	fmt.Println(a)
	// 选择排序
	fmt.Println("选择排序")
	a1 := DeepCopyArray(a)
	mySort.SelectSort(a1)
	fmt.Println(CheckArraySort(a1))

	// 冒泡排序
	fmt.Println("冒泡排序")
	a2 := DeepCopyArray(a)
	mySort.BubbleSort1(a2)
	fmt.Println(CheckArraySort(a2))

	fmt.Println("冒泡排序优化")
	a3 := DeepCopyArray(a)
	mySort.BubbleSort2(a3)
	fmt.Println(CheckArraySort(a3))

	fmt.Println("鸡尾酒排序")
	a4 := DeepCopyArray(a)
	mySort.DoubleBubbleSort(a4)
	fmt.Println(CheckArraySort(a4))

	fmt.Println("插入排序")
	a5 := DeepCopyArray(a)
	mySort.InsertSort(a5)
	fmt.Println(CheckArraySort(a5))
}

func DeepCopyArray(s []int) []int {
	res := make([]int, len(s))
	copy(res, s)
	return res
}

func CheckArraySort(arr []int) bool {
	compareSlice := DeepCopyArray(arr)
	sort.Ints(compareSlice)
	for i, v := range arr {
		if compareSlice[i] != v {
			return false
		}
	}
	return true
}
