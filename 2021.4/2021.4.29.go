package _021_4

import (
	ds "leetcode_go"
)

/**
如果root是p,q的lca，则有三种情况：
p,q在root两端
p或q是另一个的祖先。
由此递归为：
首先看root是否p或q，如果是则直接返回。首先根是p，或q，则另一个节点必然在其子树中。故root即lca。
对于函数的返回值，如果在其左子树和右子树只发现一个目标节点，则返回这个目标节点，表示这个树只存在这一个目标节点。
如果左子树和右子树返回的都是目标节点（即找到两个），说明root为lca，返回root。
如果两个子树的返回值存在不是nil，也不是目标节点，说明返回值是lca，直接返回该返回值。
如果返回值是nil，说明该子树不存在目标节点，返回nil。

*/
func lowestCommonAncestor(root, p, q *ds.TreeNode) *ds.TreeNode {
	if root == nil {
		return nil
	}
	if isPorQ(root, p, q) {
		return root
	}
	lr := lowestCommonAncestor(root.Left, p, q)
	if lr != nil {
		if !isPorQ(lr, p, q) {
			return lr
		}
	}
	rr := lowestCommonAncestor(root.Right, p, q)
	if isPorQ(lr, p, q) {
		if rr == nil {
			return lr
		} else if isPorQ(rr, p, q) {
			return root
		} else {
			return rr
		}
	} else {
		//lr == nil
		if rr == nil {
			return nil
		} else if isPorQ(rr, p, q) {
			return rr
		} else {
			return rr
		}
	}
}

func isPorQ(x, p, q *ds.TreeNode) bool {
	return x == p || x == q
}

type TopK struct {
	c    []int
	size int
}

func NewTopK(cap int) *TopK{
	return &TopK{
		c:    make([]int, cap),
		size: 0,
	}
}

func (h *TopK) Push(val int) {
	if h.size == cap(h.c) {
		if val > h.c[0] {
			h.c[0] = val
			h.adjustDown()
			return
		}
	} else {
		h.c[h.size] = val
		h.size++
		h.adjustUp()
	}
}

func (h *TopK) adjustUp() {
	cur := h.size-1
	val := h.c[cur]
	for cur > 0 {
		p := (cur - 1) / 2
		if h.c[p] > val {
			ds.Swap(h.c, cur, p)
			cur = p
		} else {
			break
		}
	}
}

func (h *TopK) adjustDown() {
	cur := 0
	for cur < h.size {
		l, r := 2*cur+1, 2*cur+2
		if l < h.size && r < h.size && h.c[r] < h.c[l] {
			ds.Swap(h.c, l, r)
		}
		if l < h.size && h.c[l] < h.c[cur] {
			ds.Swap(h.c, l, cur)
			cur = l
		} else {
			break
		}
	}
}
