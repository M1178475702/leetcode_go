package leetcode_go

import (
	"container/heap"
)

/**
 * 215. 数组中的第 K 个最大元素
 * 在未排序的数组中找到第 k 个最大的元素。
 * 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 * 0 1 2 3 4 5
 * i = 3 k = 4
 */

func findKthLargest(nums []int, k int) int {
	k--
	for k >= 0 {
		i := setoOrder(nums)
		if i == k {
			return nums[i]
		} else if i < k {
			nums = nums[i+1:]
			k = k - i - 1
		} else {
			nums = nums[:i]
		}
	}
	return nums[0]
}

func setoOrder(nums []int) int {
	i, j := 0, len(nums)-1
	b := nums[0]
	for i < j {
		for j > i {
			if nums[j] > b {
				nums[i] = nums[j]
				i++
				break
			}
			j--
		}

		for j > i {
			if nums[i] < b {
				nums[j] = nums[i]
				j--
				break
			}
			i++
		}
	}
	nums[i] = b
	return i
}

func findKthLargest1(nums []int, k int) int {
	h := NewHeapC(Int2Inter(nums), func(x interface{}, y interface{}) bool {
		xi := x.(int)
		yi := y.(int)
		return xi > yi
	})
	var ans int
	for i := 0; i < k; i++ {
		ans = heap.Pop(h).(int)
	}
	return ans
}

/**
 * 347. 前 K 个高频元素
 * 给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
 */

func topKFrequent(nums []int, k int) []int {
	fm := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := fm[nums[i]]; ok {
			fm[nums[i]]++
		} else {
			fm[nums[i]] = 1
		}
	}
	//频率map（等同于乱序数组）-> 组织成数组->堆？

	h := NewHeap([]interface{}{}, func(x interface{}, y interface{}) bool {
		xi := x.(struct {
			k int
			v int
		})
		yi := y.(struct {
			k int
			v int
		})
		return xi.v > yi.v
	})
	for k, v := range fm {
		h.Push(struct {
			k int
			v int
		}{k:k, v:v})
	}
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		k := h.Pop().(struct {
			k int
			v int
		}).k
		ans[i] = k
	}
	return ans
}

/**
 * 25. K 个一组翻转链表
 *
 * 给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。
 * k是一个正整数，它的值小于或等于链表的长度。
 * 如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。(最后一组不满K的不逆转)
 *
 */

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	iter := head
	size := 0
	for iter != nil {
		iter = iter.Next
		size++
	}
	maxRound := size / k //得到遍历最大轮
	round := 0
	hh := &ListNode{
		Val:  0,
		Next: nil,
	} //头节点
	iter = head
	i := 0
	nhh := head  //下一个头节点
	var ansh *ListNode  //答案第一个节点
	for iter != nil {
		//带头逆转代码
		tmp := iter.Next
		iter.Next = hh.Next
		hh.Next = iter
		iter = tmp
		i++
		if i == k {
			i = 0
			if ansh == nil {
				ansh = hh.Next
			}
			//正常逆转完nhh.Next=nil
			nhh.Next = iter
			round++
			if round == maxRound {
				break
			}
			hh = nhh
			nhh = iter
		}
	}
	//nhh.Next = nil
	return ansh

}

