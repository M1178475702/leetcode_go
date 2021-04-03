package leetcode_go

/**
 * 剑指 Offer 60. n个骰子的点数
 * 难度：中等
 * 把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。
 * 你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。
 * 解法：使用dp：记Cn(s)为n个骰子时，点数和为s的次数
 * Cn(s) = Σ(i=1~6)(Cn-1(s-i))
 * 可以优化成一维dp（自后向前），此外
 */

func dicesProbability(n int) []float64 {
	dp := make([][]int, 2)
	last, cur := 1, 0
	dp[0] = make([]int, 6*n+1)
	dp[1] = make([]int, 6*n+1)
	for i := 1; i <= 6; i++ {
		dp[cur][i] = 1
	}
	sum := 6 //共有多少种组合
	for i := 2; i <= n; i++ {
		//i ~ 6 * i
		tmp := last
		last = cur
		cur = tmp
		sum = 0
		for j := i; j <= 6*i; j++ {
			dp[cur][j] = 0
			for k := 1; k <= 6; k++ {
				if j-k >= i-1 { //到i-1个骰子的首个s
					dp[cur][j] += dp[last][j-k]
					sum += dp[last][j-k]
				} else {
					break
				}
			}
		}
	}
	h := float64(6*n - n + 1)
	r := make([]float64, int(h))
	k := n

	for i := 0; i < int(h); i++ {
		r[i] = float64(dp[cur][k]) / float64(sum)
		k++
	}
	return r
}
