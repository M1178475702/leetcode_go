package leetcode_go

import (
	"fmt"
	"testing"
)

func TestExist(t *testing.T) {
	fmt.Println(exist([][]byte{
		{'A','B','C','E'},
		{'S','F','C','S'},
		{'A','D','E','E'},
	}, "ABCCED"))
}

func TestSerializeTreeNode(t *testing.T) {
	root := CreateBiTree([]int{1,2,3,-1,-1,4,5}, 0)
	ser := CodecConstructor()
	deser := CodecConstructor()
	data := ser.serialize(root)
	ans := deser.deserialize(data)
	fmt.Println(ans)
}


