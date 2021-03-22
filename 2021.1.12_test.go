package leetcode_go

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	s := "the sky is blue"
	fmt.Println(reverseWords(s))
}

func TestMaxSlidingWindow(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{9, 10, 9, -7, -4, -8, 2, -6}, 5))

}

func TestMaxQueue(t *testing.T) {
	q := Constructor()
	q.Push_back(1)
	q.Push_back(2)
	fmt.Println(q.Max_value())
	fmt.Println(q.Pop_front())
	fmt.Println(q.Max_value())
}

func TestMonoDequeInt(t *testing.T) {
	m := NewMonoDequeInt(8)
	m.Push(1)
	fmt.Println(m.Front())
	m.Push(2)
	fmt.Println(m.Front())
	m.Push(2)
	fmt.Println(m.Front())
	m.Push(1)
	m.Delete(2)
	fmt.Println(m.Front())
	m.Push(3)
	fmt.Println(m.Front())
	m.Push(2)
	m.Push(1)
	fmt.Println(m.Front())
	m.Delete(2)
	fmt.Println(m.Front())
	m.Push(2)
	m.Push(1)
	m.Delete(1)
	fmt.Println(m.Front())
}


func _BinarySearch(nums []int, value int) int {
	low, hi, m := 0, len(nums)-1, 0
	for low <= hi {
		m = (low + hi) / 2
		if nums[m] == value {
			break
		} else if nums[m] < value {
			low = m + 1
		} else {
			hi = m - 1
		}
	}
	return m
}

/**
 * To bianry search, if target <= max(nums) and there is no duplicate element in nums, then nums[m] >= target else m == len(nums) - 1
 * then m + 1 is the index inserted by e
 * if there are duplicate elements the inserting index of target can't be determined(may be front or back)
 */
func TestBinarySearch(t *testing.T) {
	nums := []int{0, 2, 4, 6,6, 14,14, 16, 18}
	values := []int{6, 7, 14, 17, 18, 19}
	for _, v := range values {
		index := BinarySearch(nums, v, func(i, j int) int {
			if i < j {
				return -1
			} else if i == j {
				return 0
			} else {
				return 1
			}
		})
		fmt.Println(index)
	}
}

/**
 * 测试与slice相关的下标操作
 */
func TestSlice(t *testing.T) {
	//nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	/**
	nums[start:end]
	start: the index of the first element; end: the index of the next element of last element
	end - start: the number of slice's elements
	scene1: divided into two parts: nums[:m], nums[m:]; m is the first element of the second part
	*/



}
