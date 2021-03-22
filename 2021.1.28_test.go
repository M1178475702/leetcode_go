package leetcode_go

import (
	"fmt"
	"testing"
)

func TestNumberOfSubarrays(t *testing.T) {
	fmt.Println(numberOfSubarrays([]int{1,1,2,1,1}, 3))
}


func TestLen(t *testing.T) {
	nums := make([]int, 1e9)
	cnt := 0
	//n := len(nums)
	for i := 0; i < len(nums); i++ {
		cnt++
	}

}