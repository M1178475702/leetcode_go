package leetcode_go

import (
	"fmt"
	"testing"
)

func TestAppend(t *testing.T) {
	r := make([]int,0, 3)
	for i := 0; i < 3; i++{
		r = append(r, i)
	}
	fmt.Println(r)
}

func TestFindContinousSequence(t *testing.T) {
	r := findContinuousSequence1(15)
	fmt.Println(r)
}
