package mySort

import "fmt"

func SelectSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		m := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[m] {
				m = j
			}
		}
		arr[i], arr[m] = arr[m], arr[i]
		fmt.Println(arr)
	}
}

func BubbleSort1(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Println(arr)
	}
}

func BubbleSort2(arr []int) {
	for i := len(arr) - 1; i > 0; {
		nextI := 0
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				nextI = j
			}
		}
		i = nextI
		fmt.Println(arr)
	}
}

func DoubleBubbleSort(arr []int) {
	left, right := 0, len(arr)-1
	for left < right {
		nextRight := left
		for j := left; j < right; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				nextRight = j
			}
		}
		right = nextRight
		if left >= right {
			return
		}
		nextLeft := right
		for j := right; j > left; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				nextLeft = j
			}
		}
		left = nextLeft
		fmt.Println(arr)
	}
}

func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
		fmt.Println(arr)
	}
}
