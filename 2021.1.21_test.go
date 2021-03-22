package leetcode_go

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	n := 11
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	fmt.Println(len(permute1(nums)))
}

func TestBitmap(t *testing.T) {
	b := NewBitMap(64)
	//for i := 0; i < 12; i++ {
	//	fmt.Println(b.Get(i))
	//}
	//b.Set(0)
	//fmt.Println(b.Get(0))
	b.Set(1)
	fmt.Println(b.Get(1))
	b.Set(2)
	fmt.Println(b.Get(2))
}


/**
type SliceHeader struct {
    Data uintptr
    Len  int
    Cap  int
}
append不会修改切片的长度，只会返回更新（了长度）的切片
但如果没有扩容，其底层数组应该是相通的
var e []int -> e == nil
切片只是数组的视图！切片的len只是摆设，是逻辑上，语法上的限制（len < cap时），而不是物理上的限制
append仍是在s1[len]上赋值/追加元素，但是对原切片来说，其len不变，所以视图上仍然看不到添加的元素
切片的“截取”可以超过当前len，但需小于cap

golang参数传递时值传递，是copy；因此：修改什么就传什么的指针

 */
func TestSliceCopy(t *testing.T) {
	arr := [...]int{1,2,3,4,5}
	s1 := arr[:1]
	//s2 := arr[:2]

	s1 = append(s1, 3)
	fmt.Println(arr)
	fmt.Println(s1[:len(arr)])

	sliceParam([][]int{arr[:]})
}

func sliceParam(cr [][]int) {
	cr[1] = []int{3}
	cr = append(cr, []int{2})
	fmt.Println(cr)
}
