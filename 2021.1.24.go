package leetcode_go
/**
 * 674. 最长连续递增序列
 * 给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。
 * 连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，
 * 那么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。
 * 解法：遍历一遍，维护l,r， sl(seq len)就行了
 */

func findLengthOfLCIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	l, r := 0, 0
	sl := 0
	for i := 0; i + 1 < len(nums); i++ {
		if nums[i] < nums[i + 1] {
			r = i + 1
		} else {
			if r - l + 1 > sl {
				sl = r - l + 1
			}
			l = i + 1
			r = l
		}
	}
	if r - l + 1 > sl {
		sl = r - l + 1
	}
	return sl
}




