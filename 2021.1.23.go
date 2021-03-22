package leetcode_go

import (
	"strconv"
	"strings"
)

/**
 * 剑指 Offer 67. 把字符串转换成整数
 * 难度：中等
 */


func strToInt(str string) int {
	s := strings.TrimSpace(str)
	if len(s) == 0 || (!IsDigit(s[0]) &&  !IsSig(s[0])) {
		return 0
	}
	sig := true
	if IsSig(s[0]) {
		if s[0] == 45 { //minus
			sig = false
		}
		s = s[1:]
	}
	var i int
	for ; i < len(s); i++ {
		if !IsDigit(s[i]){
			break
		}
	}
	s = s[:i]
	num, _ := strconv.Atoi(s)

	if !sig {
		if num > 2147483648 {
			return -2147483648
		}
		return 0 - num
	} else {
		if num >= 2147483648 {
			return 2147483647
		}
		return num
	}
}


/**
 * 剑指 Offer 03. 数组中重复的数字
 * 在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
 * 解法1：使用map；解法2：使用bitmap
 */

func findRepeatNumber(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			return v
		} else {
			m[v] = struct{}{}
		}
	}
	return 0
}

func findRepeatNumber1(nums []int) int {
	b := NewBitMap(len(nums))
	for _, v := range nums {
		if b.Get(v) {
			return v
		} else {
			b.Set(v)
		}
	}
	return 0
}

/**
 * 剑指 Offer 04. 二维数组中的查找
 * 在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
 * 解法：从右上角开始，记当前元素为c，目标为t，行为i，列为j。若c == t,return true; 若 c > j，说明该列都大于t，j--；若c < j，说明该行都小于t，i++。
 * 或从左下角开始，算法与上相反。
 */

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	i, j := 0, len(matrix[0]) - 1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}



