package leetcode_go

import (
	"math"
	"sort"
)

/**
 * 1300. 转变数组后最接近目标值的数组和
 * 给你一个整数数组arr 和一个目标值target ，请你返回一个整数value，使得将数组中所有大于value 的值变成value 后，数组的和最接近 target（最接近表示两者之差的绝对值最小）。
 * 如果有多种使得和最接近target的方案，请你返回这些整数中的最小值。
 * 请注意，答案不一定是arr 中的数字。
 * 解法：基本情况：首先从ave = t/len(arr)开始检查。若ave小于数组最小值，则初始ave即ans
 * 若ave大于最小值，则二分查找枚举ave，然后找出小于ave的元素个数，及大约ave的个数，计算和（nlogn+n）
 */

func findBestValue(arr []int, target int) int {
	ave := int(math.Floor(float64(target)/float64(len(arr)) + 0.5))
	sort.Ints(arr)
	if arr[0] >= ave {
		return ave
	}
	i := BinarySearchLeft(arr, ave, nil)
	pres := 0 //pre sum
	for j := 0; j < i; j++ {
		pres += arr[j]
	}
	if arr[i] < ave {
		pres += arr[i]
		i++
	}
	mindiff := target - (pres + (len(arr)-i)*ave)
	minave := ave
	for {
		s := pres + (len(arr)-i)*ave
		if s == target {
			return ave
		} else if s < target {
			if target-s < mindiff {
				mindiff = target - s
				minave = ave
			}
		} else {
			if s-target < mindiff {
				mindiff = s - target
				minave = ave
			}
		}
		//ave处于i与i+1之间时，pres不变，此时ave增加sum增加，可以直接计算ave=arr[i+1]时的sum
		//若sum小于t，则ave<=arr[i+1]时的mindiff即t-sum，可避免单独迭代
		if s < target && i + 1 < len(arr) && pres + arr[i] + (len(arr)-i-1) * arr[i+1] <= target {
			ave = arr[i+1]
		} else {
			ave++
		}
		if ave > arr[len(arr)-1] {
			break
		}
		if ave > arr[i] {
			pres += arr[i]
			i++
		}
	}
	return minave
}

/**
 * 56. 合并区间
 * 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
 * 解法：先对区间按start排序
 */

type interval [][]int
func (t interval) Len() int{
	return len(t)
}
func (t interval) Less(i, j int) bool{
	if t[i][0] < t[j][0] {
		return true
	}
	return false
}

func (t interval) Swap(i, j int) {
	tmp := t[i]
	t[i] = t[j]
	t[j] = tmp
}

func merge(intervals [][]int) [][]int {
	sort.Sort(interval(intervals))
	i, j := 0, 1
	var ans [][]int
	for len(intervals) > 1 {
		if intervals[i][1] >= intervals[j][0] {
			intervals[j][0] = intervals[i][0]
			if intervals[i][1] >  intervals[j][1] {
				intervals[j][1] = intervals[i][1]
			}
		} else {
			ans = append(ans, intervals[i])
		}
		intervals = intervals[1:]
	}
	ans = append(ans, intervals[i])
	return ans
}


