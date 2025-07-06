package main

import "fmt"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	i := make([]int, 1, 5)
	fmt.Println(i, len(i))
	i = i[2:3]
	fmt.Println(i, len(i))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var left, leftCur, right, rightCur *ListNode
	for head != nil {
		fmt.Println("head start")
		if head.Val < x {
			if left == nil {
				left = head
				leftCur = head
				fmt.Println("left init")
			} else {
				leftCur.Next = head
				leftCur = leftCur.Next
			}
		} else {
			if right == nil {
				right = head
				rightCur = head
				fmt.Println("right init")
			} else {
				rightCur.Next = head
				rightCur = rightCur.Next
			}
		}
		head = head.Next
	}
	if leftCur == nil {
		return right
	}
	leftCur.Next = right
	fmt.Println(left.Val)
	fmt.Println(leftCur.Val)
	fmt.Println(right.Val)
	fmt.Println(right.Next)
	return left
}
