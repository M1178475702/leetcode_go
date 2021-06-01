package _021_4

import (
	"fmt"
	ds "leetcode_go"
	"testing"
)

func TestFindLCA(t *testing.T) {

	//tree := ds.CreateBiTree([]int{3,5,1,6,2,0,8,-1,-1,7,4}, 0)
	tree := ds.CreateBiTree([]int{2, 0, 3,-2,4,-1,-1,8}, 0)

	p := ds.FindTreeNode(tree, 8)
	q := ds.FindTreeNode(tree, 4)
	a := lowestCommonAncestor(tree, p, q)
	fmt.Println(a)
}

func TestTopk(t *testing.T) {
	nums := []int{1,2,3,4,5,6,7}
	tk := NewTopK(2)
	for _, v := range nums {
		tk.Push(v)
	}
	fmt.Println(tk.c)
}