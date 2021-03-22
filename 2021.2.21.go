package leetcode_go

/**
 * 169. 多数元素
 * 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于⌊ n/2 ⌋的元素。
 * 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
 */

func majorityElement(nums []int) int {
	candNum := nums[0]
	cnt := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == candNum {
			cnt++
		} else {
			cnt--
			if cnt == 0 {
				candNum = nums[i]
				cnt = 1
			}
		}
	}
	return candNum
}

/**
 * 31. 下一个排列
 * 实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
 * 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
 * 必须 原地 修改，只允许使用额外常数空间。
 * 解法：从后向前，若为逆序，一直遍历。假设最后一个逆序下标为ri
 * 将逆序列[ri:]逆转，然后找到序列中刚好大于ri-1的数i，交换ri-1与i即可
 * 排列的规律。。没什么窍门
 */

func nextPermutation(nums []int) {
	//3 2 5 => 3 5 2
	//3 2 5 1 =? 3 5 2 1
	//5 3 2 7 9 8 6 => 5 3 2 8 6 7 9  => 7 6 8 9
	ri := len(nums) - 1
	for ; ri >= 1 && nums[ri-1] > nums[ri]; ri-- {}
	//找到逆序前一个数，在后面序列的正序序列中的下一个数
	//原地修改成升序

	reverseSlice(nums[ri:])
	if ri > 0 {
		i := ri
		ri--
		for ; i < len(nums) && nums[i] <= nums[ri]; i++ {}
		if i < len(nums) {
			Swap(nums, ri, i)
		}
	}
}

func reverseSlice(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		tmp := nums[i]
		nums[i] = nums[j]
		nums[j] = tmp
		i++
		j--
	}
}
