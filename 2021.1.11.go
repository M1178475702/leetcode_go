package leetcode_go

/**
剑指 Offer 57 - II. 和为s的连续正数序列
难度：简单
解法1：暴力
解乏2：滑动窗口
*/

func findContinuousSequence(target int) [][]int {
	n := (target + 1) / 2
	var res [][]int

	for i := 1; i <= n; i++ {
		sum := 0
		var r []int
		for j := i; j <= n; j++ {
			sum += j
			if sum == target {
				r = append(r, j)
				res = append(res, r)
			} else if sum < target {
				r = append(r, j)
			} else {
				break
			}
		}
	}
	return res
}

//0ms
func findContinuousSequence1(target int) [][]int {
	n := (target + 1) / 2
	var res [][]int
	r := make([]int, 0, n)
	sum, j := 0, 1
	for i := 1; i <= n; i++ {
		sum += i
		r = append(r, i)
		if sum > target {
			for j < i {
				sum -= j
				r = r[1:]
				j++
				if sum == target {
					res = append(res, r)
					sum -= j
					r = r[1:]
					j++
					break
				} else if sum < target {
					break
				}
			}
		} else if sum == target {
			res = append(res, r)
			sum -= j
			j++
			r = r[1:]
		}
	}
	return res
}
