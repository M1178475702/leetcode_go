package _2021_3

import (
	"fmt"
	"testing"
)

func TestQsort(t *testing.T) {
	nums := []int{8,4,2,1,9,3,6, 9}
	qsort(nums)
	fmt.Println(nums)
}