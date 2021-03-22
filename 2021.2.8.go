package leetcode_go

/**
 * 5. 最长回文子串
 * 解法1：暴力
 */

//暴力；优化：只检查比当前最长回文串长的串
func longestPalindrome(s string) string {
	maxb := s[:1]
	for i := 0; i < len(s); i++ {
		for j := i + len(maxb); j < len(s); j++ {
			if isBackWords(s[i : j+1]) {
				if j+1-i > len(maxb) {
					maxb = s[i : j+1]
				}
			}
		}
	}
	return maxb
}

//palindromic substring
func isBackWords(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

//动态规划
//暴力中，对每个子串都要遍历一边以检测是否为回文串；引入状态矩阵快速判断是否为回文串
//对于长度大于2的每个回文串s[i...j]，去掉两边的字符s[i],s[j]仍为回文串
//因此记录p[i][j]为s[i,j]是否为回文串；p[i][j] = p[i+1][j-1] && (s[i] == s[j])
//对于长度等于2的串，若两个字符相等，则为回文；对于一个字符的串必为回文
//实现中，迭代时引用的状态为p[i+1][j-1]，因而，若使用for i,j循环，引用的状态是还没有更新的状态！
//仔细考虑，迭代过程为，检查比自己更短的字符串是否为回文，而非ij循环。
//因此，以长度，和起始位置为基准迭代。从长度0开始迭代。由此可确保p[i+1][j-1]是被更新过的状态。
//此外，由于该法需要之前的状态，即使该状态不可能为最长回文串（长度小于当前最长回文串），也要计算。
//因此无法省略哪些长度小于当前最长串的串。因此效率可能还不如暴力，且还费空间
func longestPalindrome1(s string) string {
	p := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		p[i] = make([]bool, len(s))
	}
	maxb := s[:1]
	//长度l
	for l := 0; l < len(s); l++ {
		//起始位置i（以i为开头的字符串）
		for i := 0; i+l < len(s); i++ {
			j := i + l
			if i == j {
				p[i][j] = true
			} else if j-i == 1 {
				if s[i] == s[j] {
					p[i][j] = true
				}
			} else {
				//对状态的访问是i+1, j-1，因此更新状态应该
				p[i][j] = p[i+1][j-1] && (s[i] == s[j])
			}
			if p[i][j] && j+1-i > len(maxb) {
				maxb = s[i : j+1]
			}
		}
	}
	return maxb
}

//中心轴扩散
func longestPalindrome2(s string) string {
	maxb := s[:1]
	for i := 1; i < len(s); i++ {
		{ //以0,m,1从m开始
			{
				l, r := i-1, i
				for l >= 0 && r < len(s) && s[l] == s[r] {
					if r+1-l > len(maxb) {
						maxb = s[l : r+1]
					}
					l--
					r++
				}
			}
			{ //以0,1,2从1开始
				l, r := i, i
				for l >= 0 && r < len(s) && s[l] == s[r] {
					if r+1-l > len(maxb) {
						maxb = s[l : r+1]
					}
					l--
					r++
				}
				if r == len(s) {
					break
				}
			}
		}
	}
	return maxb
}

/**
 * 42. 接雨水
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 */

//暴力
//先找到最高点，以最高点为分割，向左或向右，0:max或max:len-1中的第二高点，然后以第二高点为准计算区间内积水量
//优化，使用前缀和，后缀和优化max查找 -> 时间：100%!
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	ans := 0
	var lmi, rmi int
	si, sv := -1, 0
	//初始化前缀后缀
	pre, suf := make([]int, len(height)), make([]int, len(height))
	pre[0], suf[len(height)-1] = 0, len(height)-1
	for i := 1; i < len(height); i++ {
		if height[i] > height[pre[i-1]] {
			pre[i] = i
		} else {
			pre[i] = pre[i-1]
		}
	}
	for i := len(height) - 2; i >= 0; i-- {
		if height[i] > height[suf[i+1]] {
			suf[i] = i
		} else {
			suf[i] = suf[i+1]
		}
	}

	lmi = pre[len(height)-1]
	rmi = lmi

	//向左查找
	for lmi > 0 {
		//si = max(height[:lmi])
		si = pre[lmi-1]
		sv = height[si]
		//计算存水量
		ans += cal(height[si:lmi+1], sv)
		lmi = si
	}
	//向右查找
	for rmi < len(height)-1 {
		//si = rmi + max(height[rmi+1:]) + 1
		si = suf[rmi+1]
		sv = height[si]
		//计算存水量
		ans += cal(height[rmi:si+1], sv)
		rmi = si
	}
	return ans
}

func cal(h []int, sv int) int {
	ans := 0
	for i := 1; i < len(h)-1; i++ {
		ans += sv - h[i]
	}
	return ans
}

//func max(height []int) int {
//	m := 0
//	for i, v := range height {
//		if v > height[m] {
//			m = i
//		}
//	}
//	return m
//}

/**
 * 1014. 最佳观光组合
 */
//暴力：超时
//使用后缀和优化：记录后缀最大值（下标），如果i到j的距离大于A[suf[j]]，即：即使A[j]时后续数组中的最大值
//也不能弥补距离，那么后续肯定不存在最优解（只是稍微减少了范围）但是勉强通过了=。=（8%）
func maxScoreSightseeingPair(A []int) int {
	max := 0
	suf := make([]int, len(A))
	suf[len(A)-1] = len(A) - 1
	for i := len(A) - 2; i >= 0; i-- {
		if A[i] > A[suf[i+1]] {
			suf[i] = i
		} else {
			suf[i] = suf[i+1]
		}
	}


	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if j - i > A[suf[i]]  {
				break
			}
			if A[i]+A[j]+i-j > max {
				max = A[i] + A[j] + i - j
			}
		}
	}
	return max
}
//动态规划
//枚举j，到j的最大评分。暴力的话时对每个j枚举0-j-1 ->n^2
//优化：分数计算公式A[i] + A[j] + i - j可以写成 A[i] + i + A[j] - j
//对每个j，A[j] - j是固定的，因此只需知道最大的A[i] + i 即可。
//而由于A[i] + i中只有一个变量A[i]是非规律变化，因此可以在迭代时用前缀和计算。
//计算每个j的最大A[i]+i只需O（1），因此时间降为O（N）

//对动态规划的公式，一是迭代角度，而是对公式的再研究，包括找辅助变量等，以产生可以迭代的状态！需要多练习
func maxScoreSightseeingPair1(A []int) int {
	maxs := 0
	maxi := A[0]
	for j := 1; j < len(A); j++ {
		if A[j-1] + j-1 > maxi {
			maxi = A[j-1]+j-1
		}
		if maxi + A[j] - j > maxs {
			maxs = maxi + A[j] - j
		}
	}
	return maxs
}