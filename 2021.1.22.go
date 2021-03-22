package leetcode_go

/**
 * 剑指 Offer 63. 股票的最大利润
 * 难度：中等
 * 假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？
 * 解法1：使用单调队列，第i天买入的收益为p[i + 1:]中最大值减去p[i]。使用单调队列维护[i+1:]中最大值
 */

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	var max int
	q := NewMonoDequeEleWithIndex(len(prices))
	for i := 1; i < len(prices); i++ {
		q.Push(prices[i], i)
	}

	for i, v := range prices {
		q.Delete(i)
		if q.Empty() {
			break
		}
		diff := q.Front() - v
		if diff > max {
			max = diff
		}
	}
	return max
}
