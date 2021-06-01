package leetcode_go

import (
	"sort"
)

/**
 * 4. 寻找两个正序数组的中位数
 * 解法：转化成寻找两个有序数组中的第k小数，k = (m+n)/2
 * 一种可以用归并，效率为O(m+n)，题目要求O(log(m+n))，因此应该使用二分
 * 二分思想：第k小的数前，有k个小于该数的数。对AB两个数，开始比较A[K/2 - 1]与B[K/2 - 1]（二者前面包括二者刚好两个数）
 * 若A[K/2 - 1]<=B[K/2 - 1]，则A[0-k/2-1] <= B[k/2-1]，因此A[0-K/2-1]都没有K-1的数小于其，因此不可能是第K小的数
 */

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex + 1))
	} else {
		midIndex1, midIndex2 := totalLength/2 - 1, totalLength/2
		//偶数个情况，找第k和第k+1个
		return float64(getKthElement(nums1, nums2, midIndex1 + 1) + getKthElement(nums1, nums2, midIndex2 + 1)) / 2.0
	}

}
//专注于：寻找第K小的数
func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		//检查是否有数组为空
		if index1 == len(nums1) {
			return nums2[index2 + k - 1]
		}
		if index2 == len(nums2) {
			return nums1[index1 + k - 1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k/2
		newIndex1 := min(index1 + half, len(nums1)) - 1
		newIndex2 := min(index2 + half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}


/**
 * 15. 三数之和
 * 给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
 * 注意：答案中不可以包含重复的三元组。
 * 解法：暴力时三重循环，但此时还要去重。现在如何不去重上优化：a <= b <= c时，不会重复。因此，将原数组排序，然后三重循环。
 * 对于第二重和第三重，由于b增大，因而对应c减小。由此，b，c是-> <-的趋势。因此，由于a,b确定，先用二分找到c的位置，然后滑动b,c。此时将二三重的N^2优化成N
 * 结合第一重的N，最终复杂度为N^2
 * 此外，对于重复的元素，需保证a不重复，c不重复，b自然不重复。
 * 「双指针」，当我们需要枚举数组中的两个元素时，如果我们发现随着第一个元素的递增，第二个元素是递减的，那么就可以使用双指针的方法，将枚举的时间复杂度从 O(N^2)减少至O(N)。
 */

func threeSum(nums []int) [][]int {
	if len(nums) <= 2 {
		return [][]int{}
	}
	sort.Ints(nums)
	var ans [][]int
	for i := 0; i < len(nums) - 2; i++ {
		//开始，使用二分查找c的位置，及0-a-b；此时nums[k]要么等于c，要么小于c
		//之后的c的位置不再用二分查找，改用缩小的范围的双指针查找（也可以直接用缩小的范围的二分，不用双指针
		//先以Nums[i+1]为b，其为最小b，然后确定最大c位置，然后二分
		k := BinarySearchRight(nums, 0 - nums[i] - nums[i+1], nil)
		for j := i + 1; j < len(nums) && j < k; j++ {
			for k > j {
				if nums[i] + nums[j] + nums[k] == 0 {
					ans = append(ans, []int{nums[i], nums[j], nums[k]})
					//确保c不重复
					for ; k > j && nums[k - 1] == nums[k]; k-- {}
					k--
					break
				} else if nums[i] + nums[j] + nums[k] < 0 {
					//由于0...k-1小于k，若到k已经小于0,后面都小于0，可退出
					break
				} else {
					k--
				}
			}
		}
		//确保a不重复
		for ; i + 1 < len(nums) - 2 && nums[i + 1] == nums[i]; i++ {}
	}
	return ans
}


/**
 * 238. 除自身以外数组的乘积
 * 给你一个长度为n的整数数组nums，其中n > 1，返回输出数组output，其中 output[i]等于nums中除nums[i]之外其余各元素的乘积。
 * 不能使用乘法。时间限制O(n)；因此出i以外的乘积可表示为0...i-1的乘积与i+1...n-1的乘积之乘积，即前缀积与后缀积之积。
 * 空间优化：没必要同时保存pre和back。可以把output当做pre。而当计算出back[i]时，直接将其应用到output上。因此pre和back都省去了
 */

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	pre, back := make([]int, len(nums)), make([]int, len(nums))
	output := make([]int, len(nums))
	pre[0] = 1
	back[len(nums) - 1] = 1
	for i := 1; i < len(nums); i++ {
		pre[i] = pre[i - 1] * nums[i - 1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		back[i] = back[i + 1] * nums[i + 1]
	}
	for i := 0; i < len(nums); i++ {
		output[i] = pre[i] * back[i]
	}
	return output
}
//
func productExceptSelf1(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	output := make([]int, len(nums))
	output[0] = 1
	back := 1
	for i := 1; i < len(nums); i++ {
		output[i] = output[i - 1] * nums[i - 1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		back = back * nums[i + 1]
		output[i] *= back
	}
	return output
}

















