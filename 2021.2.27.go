package leetcode_go

/**
 * 395. 至少有K个重复字符的最长子串
 * 给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。
 * 返回这一子串的长度。
 * 解法：字符串长度为k*n+m，n为不重复字符数量，m为字符数减去k的综合
 * 考虑从前往后迭代：如果以i结尾的字符串中，最长子串li，若如果i字符数量不足k，则不能要，因此为0。
 * 求单个字符串的某个子串，可用滑动窗口（滑动窗口需要解决的问题）
 */


func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}
	ans := 0
	m := make(map[byte]*struct{v, i int})
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			m[s[i]].i = i
			m[s[i]].v++
		} else {
			m[s[i]] = &struct{ v, i int }{v: 1, i: i}
		}
	}
	b := -1
	for _, p := range m {
		if p.v < k {
			b = p.i
			break
		}
	}
	if b != -1 {
		ans = MaxInt(longestSubstring(s[:b], k), longestSubstring(s[b+1:], k))
	} else {
		return len(s)
	}
	return ans
}
