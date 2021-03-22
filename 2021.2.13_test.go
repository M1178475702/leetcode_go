package leetcode_go

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	fmt.Println(findKthLargest1([]int{3, 2, 1, 5, 6, 4}, 2))
}

func TestTopKFrequent(t *testing.T) {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

func TestReverseKGroup(t *testing.T) {
	fmt.Println(reverseKGroup(CreateListInt([]int{1, 2, 3, 4, 5}), 3))
}

func TestReverseList(t *testing.T) {
	fmt.Println(reverseList(CreateListInt([]int{1, 2, 3, 4, 5})))
}

func TestMergeTwoLists(t *testing.T) {
	fmt.Println(mergeTwoLists(CreateListInt([]int{1, 2, 4}), CreateListInt([]int{1, 3, 4})))
}

func TestMaxPathSum(t *testing.T) {
	p := maxPathSum(CreateBiTree([]int{-10, 9, 20, -1, -1, 15, 7}, 0))
	fmt.Println(p)
}

func TestGenerateParenthesis(t *testing.T) {
	fmt.Println(generateParenthesis(4))
}

func TestPathToTreeNode(t *testing.T) {
	tr := CreateBiTree([]int{1,2,-1,3,6}, 0)
	fmt.Println(PathToTreeNode(tr,tr.Left.Right))
}

func TestLowestCommonAncestor(t *testing.T) {
	tr := CreateBiTree([]int{1,2,-1,3,6}, 0)
	fmt.Println(lowestCommonAncestor(tr, tr.Left.Left, tr.Left.Right))
}

func TestMaxValue(t *testing.T) {
	fmt.Println(maxValue([][]int{
		{1,3,1},
		{1,5,1},
		{4,2,1},
	}))
}

func TestFindMaxForm(t *testing.T) {
	fmt.Println(findMaxForm([]string{"10","0", "1"}, 1, 1))
}

func TestLengthOfLIS(t *testing.T) {
	fmt.Println(lengthOfLIS([]int{10,9,2,5,3,7,101,18}))
}

