package _2021_3

import (
	"fmt"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	fmt.Println(hammingWeight(00000000000000000000000000001011))
}

func DeleteSlice(s []int, i int) []int {
	s = append(s[:i], s[i+1:]...)
	return s
}

func TestDeleteSliceWhenIterating(t *testing.T) {
	src := []int{1, 2, 3, 4,5,6}
	for i, v := range src {
		if v == 2 || v == 3 {
			src = DeleteSlice(src, i)
		}
	}
	//i = 0, v = 1; i = 1, v = 2; delete(i=1) （1,3,4,5,6）; i = 2, v = 4
	fmt.Println(src)

}
