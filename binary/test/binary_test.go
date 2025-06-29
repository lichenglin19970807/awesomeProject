package test

import (
	"awesomeProject/binary"
	"fmt"
	"math"
	"testing"
)

func TestBinary(t *testing.T) {
	// 打印二进制
	binary.PrintBinary(78)
	binary.PrintBinary(-78)

	// 相反数
	fmt.Println(binary.OppositeNum(78))

	// 最小数
	fmt.Println(math.MinInt)
	// fmt.Println(^math.MinInt + 1)
	// go编译时常量表达式严格检测
	// go运行时计算溢出静默回绕
	a := 9223372036854775807
	b := 111
	fmt.Println(a + b)
}

func TestBitOperation(t *testing.T) {
	// 短路特性
	a := func() bool {
		fmt.Println("a")
		return true
	}
	b := func() bool {
		fmt.Println("b")
		return true
	}
	fmt.Println(a() || b())

	// 左移右移
	binary.PrintBinary(78)
	binary.PrintBinary(78 << 1)
	binary.PrintBinary(78 >> 1)

	binary.PrintBinary(-78)
	binary.PrintBinary(-78 << 1)
	binary.PrintBinary(-78 >> 1)

	x := 9223372036854775807
	fmt.Println(x << 1)
}
