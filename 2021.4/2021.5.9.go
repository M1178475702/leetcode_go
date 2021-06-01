package _021_4

import (
	ds "leetcode_go"
)

func zigzagLevelOrder(root *ds.TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	q := ds.NewCQueue(10)
	left := true
	b := new(ds.TreeNode)

	q.Push(root)
	q.Push(b)
	level := []int{}
	for !q.Empty() {
		node := q.Pop().(*ds.TreeNode)
		if node == b {
			tmp := make([]int, len(level))
			_copy(tmp, level, left)

			ans = append(ans, tmp)
			level = level[:0]

			if !q.Empty() {
				left = !left
				q.Push(b)
			} else {
				break
			}
		} else {
			level = append(level, node.Val)
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}

		}
	}
	return ans
}

func _copy(dst, src []int, reverse bool) {
	if !reverse {
		copy(dst, src)
	} else {
		for i := len(src)-1; i>=0; i-- {
			dst[i] = src[i]
		}
	}
}
