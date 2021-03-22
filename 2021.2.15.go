package leetcode_go

/**
 * 两数相加
 * 给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
 * 请你将两个数相加，并以相同形式返回一个表示和的链表。
 * 你可以假设除了数字 0 之外，这两个数都不会以 0开头。

 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1i := l1
	l2i := l2
	var ansh, ansr *ListNode
	up := false
	for l2i != nil && l1i != nil {
		s := l1i.Val + l2i.Val
		if up {
			s += 1
			up = false
		}
		if s >= 10 {
			s = s - 10
			up = true
		}
		if ansr == nil {
			ansh = &ListNode{Val: s, Next: nil}
			ansr = ansh
		} else {
			ansr.Next = &ListNode{Val: s, Next: nil}
			ansr = ansr.Next
		}
		l1i = l1i.Next
		l2i = l2i.Next
	}
	ri := l1i
	if l2i != nil {
		ri = l2i
	}
	for ri != nil {
		s := ri.Val
		if up {
			s += 1
			up = false
		}
		if s >= 10 {
			s = s - 10
			up = true
		}
		if ansr == nil {
			ansh = &ListNode{Val: s, Next: nil}
			ansr = ansh
		} else {
			ansr.Next = &ListNode{Val: s, Next: nil}
			ansr = ansr.Next
		}
		ri = ri.Next
	}
	if up {
		ansr.Next = &ListNode{Val: 1, Next: nil}
		ansr = ansr.Next
	}
	return ansh
}

/**
 * 206. 反转链表
 */
//递归版
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	h := reverseList(head.Next)
	if h.Next == nil {
		h.Next = head
	}
	//cur.Next is tail; tail.Next = cur; cur.Next = nil(cur becomes new tail)
	head.Next.Next = head
	head.Next = nil
	return h
}

//迭代
func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	h, iter, next := &ListNode{}, head, head.Next
	for iter != nil {
		next = iter.Next
		iter.Next = h.Next
		h.Next = iter
		iter = next
	}
	return h.Next
}

/**
 * 21. 合并两个有序链表
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var h, r, ne *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			ne = l1
			l1 = l1.Next
		} else {
			ne = l2
			l2 = l2.Next
		}
		if r == nil {
			r = ne
			h = ne
		} else {
			r.Next = ne
			r = ne
		}
	}
	var rest *ListNode
	if l1 != nil {
		rest = l1
	} else {
		rest = l2
	}
	if r == nil {
		r = rest
		h = r
	} else {
		r.Next = rest
	}
	return h
}

/**
 * 124. 二叉树中的最大路径和
 */

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxSum := root.Val
	var subr func(root *TreeNode) int
	//TODO 没办法在声明时，写递归函数（在函数体引用自己），只能先声明，后赋值
	subr = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lps, rps := subr(root.Left), subr(root.Right)
		//作为根节点，其可以左右子树都走
		max := root.Val + MaxInt(0, lps) +  MaxInt(0, rps)
		//作为子树，其返回给父节点的路径和不应该左右子节点都走
		subMax := root.Val + MaxInt(0, MaxInt(lps, rps))
		//计算最大和时，两个都比较
		maxSum = MaxInt(maxSum, MaxInt(max, subMax))
		//返回作为子树时的值，由于结果是通过修改闭包变量maxSum而得，所以无所谓最后返回值
		return subMax
	}
	subr(root)
	return maxSum
}


//下面的代码时计算节点路径和（不是题目中的最大路径和）的代码
//func maxPathSum(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	s, _ := maxPathSumWithSize(root)
//	return s
//}
//
//func maxPathSumWithSize(root *TreeNode) (int, int) {
//	if root == nil {
//		return 0, 0
//	}
//	lp, ls := maxPathSumWithSize(root.Left)
//	rp, rs := maxPathSumWithSize(root.Right)
//	return lp + ls + rp + rs, ls + rs + 1
//}
