/**
@Title 2021_1_6
@Description leetcode
@Author	Bakatora
@CreateAt 2021.1.16
*/

package leetcode_go

import (
	"math"
)

/**
 * 剑指 Offer 52. 两个链表的第一个公共节点（没有头结点）
 * 难度：简单
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func LenOfList(head *ListNode) int {
	iter := head
	i := 0
	for iter != nil {
		iter = iter.Next
		i++
	}
	return i
}

//先遍历长度，长的先走，然后再一起走
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	lenA, lenB := LenOfList(headA), LenOfList(headB)
	var iter1, iter2 *ListNode
	var diff int
	if lenA >= lenB {
		iter1 = headA
		iter2 = headB
		diff = lenA - lenB
	} else {
		iter1 = headB
		iter2 = headA
		diff = lenB - lenA
	}

	for diff > 0 {
		iter1 = iter1.Next
		diff--
	}
	for iter1 != nil {
		if iter1 == iter2 {
			return iter1
		}
		iter1 = iter1.Next
		iter2 = iter2.Next
	}
	return nil
}

//双指针法，A走完再走一遍B，B走完再走一遍A，可以看成两个指针将AB过了一遍，若存在公共节点，则一定会相遇
//将两个链表分别分成两段：Aa和Ba，其中a表示公共段
//两个指针分别遍历AaBa和BaAa，可以看到AaB和BaA是同长，因而二者不然同时走到第二个a
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	iterA, iterB := headA, headB
	var aInB, bInA bool
	for iterA != nil && iterB != nil {
		if iterA == iterB {
			return iterA
		}
		if iterA.Next == nil && !aInB {
			iterA = headB
			aInB = true
		} else {
			iterA = iterA.Next
		}
		if iterB.Next == nil && !bInA {
			iterB = headA
			bInA = true
		} else {
			iterB = iterB.Next
		}
	}
	return nil
}

/**
	剑指 Offer 53 - I. 在排序数组中查找数字 I
	难度：简单
	先二分找到位置，再排查左右
*/

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	var low, hi = 0, n - 1
	var m int
	for low <= hi {
		m = (low + hi) / 2
		if nums[m] == target {
			break
		} else if nums[m] <= target {
			low = m + 1
		} else {
			hi = m - 1
		}
	}
	if nums[m] != target {
		return 0
	}
	var c = 1
	if m > 0 {
		for i := m - 1; i >= 0; i-- {
			if nums[i] == target {
				c++
			} else {
				break
			}
		}
	}
	if m < len(nums) {
		for i := m + 1; i < n; i++ {
			if nums[i] == target {
				c++
			} else {
				break
			}
		}
	}
	return c
}

/**
	剑指 Offer 53 - II. 0～n-1中缺失的数字
	难度：简单
	首先，如果nums[low]正常，则一定是右边的“坏了（下标和值不一样）”
	通过二分确定“坏段”的左边界
	1. 如果nums[low]不正常，则low即为缺少的数（如nums[0] != 0一定是少0）
	2. 如果nums[hi] = hi，则说明hi为“正常段”的右边界，少的数为hi + 1
	3. 判断中间点m。若nums[m] = m，说明m及m左边为正常段，“坏段”在m右边，low = m + 1
 	   若nums[m] != m，说明m处于“坏段”中，“坏段”左边界在m左边，hi = m - 1
*/

func missingNumber(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	low, hi := 0, n-1
	for low <= hi {
		if nums[low] != low {
			return low
		}
		if nums[hi] == hi {
			break
		}
		m := (low + hi) / 2
		if nums[m] == m {
			low = m + 1
		} else {
			hi = m - 1
		}
	}
	return hi + 1
}

/**
	剑指 Offer 54. 二叉搜索树的第k大节点
	难度：简单
	解法1：用递归，将问题简化
	先求右子树结点数，判断与k大小
	解法2：二叉搜索树的中序遍历是元素的递增序列，将中序的左右子树遍历顺序反过来（即右-中-左），便是逆序序列
		1. 用 go walk解决
		2. 用迭代中序遍历
*/

func walk(root *TreeNode, ch chan int) {
	if root == nil {
		return
	}
	walk(root.Right, ch)
	ch <- root.Val
	walk(root.Left, ch)
}
//12ms, 6.6mb 使用ch + 递归比用栈费空间
func kthLargest(root *TreeNode, k int) int {
	ch := make(chan int)
	go walk(root, ch)
	var r int
	for i := 0; i < k; i++ {
		r = <-ch
	}
	//TODO close没法通知后续接收的通道，后续接收视为堵塞？
	//close(ch)
	return r
}

//12ms, 6.3mb
func kthLargest1(root *TreeNode, k int) int {
	iter := root
	s := NewStackTreeNode(k)
	var c = 0
	for iter != nil || !s.Empty() {
		if iter != nil {
			s.Push(iter)
			iter = iter.Right
		} else {
			c++
			iter = s.Pop()
			if c == k {
				break
			}
			iter = iter.Left
		}
	}
	//由于k <= n，所以iter一定不为nil
	return iter.Val
}

/**
	剑指 Offer 55 - I. 二叉树的深度
	难度：简单
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l <= r {
		return r + 1
	}
	return l + 1
}


/**
	剑指 Offer 55 - II. 平衡二叉树，判断是否平衡
	难度：简单
	解法1. （较低效）后序遍历 + 对每个结点求左右深度 先判断左右子树是否平衡，若都平衡，再判断左右子树高度之差是否大于1
	解法2. 解法1每次对结点求深度，都要重新对已求过深度的子树再求一次。使用_isBalanced函数，在返回是否平衡同时，返回深度，避免重复计算深度
 */

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l - r > 1 || r - l > 1 {
		return false
	}
	return true
}

// 0ms（打败100%用户！） 6.1mb
func isBalanced1(root *TreeNode) bool {
	b, _ := _isBalanced(root)
	return b
}
//返回是否平衡，及深度
func _isBalanced(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	lb, ld := _isBalanced(root.Left)
	if !lb {
		return false, 0 //此时depth的值已经没有意义，随便返回
	}
	rb, rd := _isBalanced(root.Right)
	if !rb {
		return false, 0
	}
	if ld - rd > 1 || rd - ld > 1 {
		return false, 0
	}
	return true, int(math.Max(float64(ld), float64(rd)) + 1)
}

/**
	剑指 Offer 56 - I. 数组中数字出现的次数
	一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
	难度：中等
	解法：1. 异或（^）可以记住操作数一次（类似状态叠加），当与曾经异或过一次的操作数再次异或时，会抵消
		    由此，对一个数列（只有一个数是单次，其他都是双次）进行异或遍历，最后的值即是该单次数
		 2. 由1，将数列分成两部分，每部分只有一个单次数，其他都是双次；分别对两部分再次异或遍历，即得到两个单次数
		 3. 分成两部分：1）对整个序列异或遍历一次，得到两个单次数的异或结果，记作x
                      2）任取x的不为0的一位，记作第i位。
					  3）对整个序列遍历，按第i位是否为1分组，如元素的第i位为0则放到左边，为1放到右边，由此分成两部分
					  4）这两部分的每部分保证仅包括一个单次，其他都是双次（即刚好把成对的放在一遍）
		 4. 3原理：对所有元素进行异或，对第i位，所有的双次的两个数之间相同，都抵消为0；对两个单次，由于选取的第i位为1，说明两个单次的第i位不同
			从而被分到两部分
 */

func singleNumbers(nums []int) []int {
	n := len(nums)
	if n == 2 {
		return nums
	}
	x := nums[0]
	for _, v := range nums[1:] {
		x = x ^ v
	}
	off := 0
	for k := x; k % 2 == 0; off++ {
		k >>= 1
	}
	l, r := 0, n - 1
	//将第off位为1的移到右边
	for l < r  {
		if nums[l] >> off % 2 == 1 {
			Swap(nums, l, r)
			r--
		} else {
			l++
		}
	}
	// l == r
	var lb int //确保lb是左半区的最后一个元素
	if nums[l] >> off % 2 == 0 {
		lb = l
	} else {
		lb = l - 1
	}
	res := []int{nums[0], nums[lb + 1]}
	for i := 1; i <= lb; i++ {
		res[0] ^= nums[i]
	}
	//lb + 1已经赋给res[1]，直接从lb + 2开始遍历
	for i := lb + 2; i < n; i++ {
		res[1] ^= nums[i]
	}
	return res
}









