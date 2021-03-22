package leetcode_go

/**
 * 3. 无重复字符的最长子串
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 * 解法：子数组问题。摒弃暴力解法思路。选定一个维度（为什么是这个维度？尚未得解，目前只是碰）。就本题，选择以s[i]结尾的字符串。
 * 记max[i]为以s[i]结尾的字符串中，不重复的最大字符数
 * 迭代公式推导：即推导max[i]与max[i-1]的关系：若以是s[i-1]结尾的最大不重复字符串不包含s[i]，则max[i] = max[i-1]+1
 * 若重复，则max[i] = i - charIdx[s[i]]，即[charIdx[s[i]] + 1, ...i]是以s[i]结尾的最大不重复字符串
 * 实现中只需维护max[i-1]，记作last，不用维护max数组
 * 下面尝试推理下为什么可以使用这个角度：表层原因，以i结尾的子串，其总是“紧贴右边的”，以此可以推导max[i]与max[i-1]的关系
 * 此外，子串所要符合的关系并不复杂。因此max[i]与max[i-1]的关系比较好推导
 * 若更复杂，最好还是滑动窗口
 */

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	//记录最近的一次字符出现的下标
	charIdx := make([]int, 256)
	//初始化为-1，默认值0会与下标0冲突 //可以换用map
	for i := 0; i < 256; i++ {
		charIdx[i] = -1
	}
	last := 0 //以i-1结尾的字符串中不重复的最大字符数
	ans := 0
	for i := 0; i < len(s); i++ {
		if i-charIdx[s[i]] <= last {
			//重复
			last = i - charIdx[s[i]]
		} else {
			//不重复
			last = last + 1
		}
		charIdx[s[i]] = i
		if last > ans {
			ans = last
		}
	}
	return ans

}

//使用sliding window枚举所有以k开头的最大字符串
//最大的麻烦是，如何调整起始和开始遍历的第一个条件
func lengthOfLongestSubstring1(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	ans := 0
	k, rk := 0, 0
	charIdx := make([]int, 256)
	for i := 0; i < 256; i++ {
		charIdx[i] = -1
	}
	charIdx[s[0]] = 0
	for ; k < len(s) && rk+1 < len(s); k++ {
		for rk+1 < len(s) {
			//non-dup
			if v := charIdx[s[rk+1]]; v < k {
				rk++
				charIdx[s[rk]] = rk
			} else {
				break
			}

		}
		if rk-k+1 > ans {
			ans = rk - k + 1
		}
	}
	return ans
}
