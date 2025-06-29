package validator

import (
	"math/rand"
	"time"
)

func NewRandomArray(length, max int) []int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = r.Intn(max)
	}
	return res
}
