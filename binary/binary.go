package binary

import "fmt"

func PrintBinary(num int32) string {
	fmt.Printf("十进制: %d\n", num)
	var binaryNum string
	for i := 31; i >= 0; i-- {
		if num&(1<<i) == 0 {
			binaryNum += "0"
		} else {
			binaryNum += "1"
		}
	}
	fmt.Printf("二进制: %s\n", binaryNum)
	return binaryNum
}

func OppositeNum(num int32) bool {
	a1 := PrintBinary(-num)
	a2 := PrintBinary(^num + 1)
	if a1 == a2 {
		return true
	}
	return false
}
