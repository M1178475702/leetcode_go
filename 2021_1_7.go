package leetcode_go


/**
	剑指 Offer 57. 和为s的两个数字
	输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。
	难度：简单
	解法：二分找到s/2，使a处于0位置，b处于s的位置，两个数，a，b，要么a=b，要么a<s/2, b>s/2；a向前，b向后遍历，直到s/2
 */

func twoSum(nums []int, target int) []int {
	n := len(nums)
	low, hi := 0, n - 1
	var m int
	for low < hi {
		m = (low + hi) / 2
		if nums[m] == target {
			break
		} else if nums[m] > target {
			hi = m - 1
		} else {
			low = m + 1
		}
	}
	//使nums[b] >= target的关系
	var a, b int
	if nums[m] >= target {
		b = m
	} else if m + 1 < n {
		b = m + 1
	}
	for a < b && nums[a] <= target / 2 {
		if nums[a] + nums[b] == target {
			return []int{nums[a], nums[b]}
		} else if nums[a] + nums[b] < target{
			a++
		} else {
			b--
		}
	}
	return nil

}



