package leetcode_go


/**
 * 给你一个整数数组 nums 和一个整数 k。
 * 如果某个连续子数组中恰好有 k 个奇数数字，我们就认为这个子数组是「优美子数组」。
 * 请返回这个数组中「优美子数组」的数目。
 * 暴力 => 通不过
 * 解法：前缀和：pre[i]表示0-i中有几个奇数 + 频次数组
 * 则子数组[l, r]是否为优美子数组，取决于pre[r]-pre[l]==k；
 * 交换下项：pre[l] = sum[r] - k，即对于pre[r]，以r结尾的优美子数组的数量为奇数数量为为pre[l]的频次
 * 因此，ans = preCnt[sum[r] - k]。此外，实现中无需维护pre[]，只需维护odd变量即可。preCnt[0] = 1.
 *
 * 前缀和特征：统计前缀数组某种特征的个数，和等信息
 */

func numberOfSubarrays(nums []int, k int) int {

	//没必要为sum建立一个数组保存遍历过的奇数数量
	odd := 0
	preCnt := make([]int, len(nums) + 1)
	preCnt[0] = 1
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] % 2 != 0 {
			odd++
		}
		//记录频次
		preCnt[odd]++
		if odd - k >= 0 {
			ans += preCnt[odd - k]
		}
	}
	return ans
}

/**
 * 剑指 Offer 42. 连续子数组的最大和
 * 输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。要求时间复杂度为O(n)。
 * 解法：若暴力，则是统计所有子数组的和。这里尝试简化：尝试寻找纬度，将问题变成单向迭代问题。
 * 如，考虑以i结尾的子数组中的最大值max[]，这样就将问题变成了单向迭代。
 * 对max[i]，以i结尾，那么唯一的问题就是该子数组是否包括以i-1为结尾的子数组（不用考虑右边）
 * 而是否要包括i-1，则考虑max[i-1]是否为负数，若为负数则不包括i-1。
 * 因此推导公式为：max[i] = max[i - 1] < 0 ? max[i - 1]
 * 实现中，由于只访问上一次迭代的结果，不会随机访问其他位置的结果，因此只需用变量记录上一次结果的值（max[i-1]）即可，无需声明数组
 */

func maxSubArray(nums []int) int {
	last := nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if last < 0 {
			last = nums[i]
		} else {
			last = last + nums[i]
		}
		if last > ans {
			ans = last
		}

	}
	return ans
}