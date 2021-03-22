package leetcode_go

/**
 * 1004. 最大连续1的个数 III
 * 给定一个由若干 0 和 1 组成的数组 A，我们最多可以将 K 个值从 0 变成 1 。
 * 返回仅包含 1 的最长（连续）子数组的长度。
 */

func longestOnes(A []int, K int) int {
	ans, i, j, k := 0, 0, 0, 0
	if K == 0 {
		for j < len(A)-1 {
			for i < len(A) {
				if A[i] == 0 {
					i++
				} else {
					break
				}
			}
			if i == len(A) {
				break
			}
			j = i
			for j+1 < len(A) {
				if A[j+1] == 1 {
					j++
				}else {
					break
				}
			}
			ans = MaxInt(ans, j - i + 1)
			i = j+1
		}
	} else {
		q := NewSQueue(K)
		for j < len(A)-1 {
			//丢弃第一个0
			if !q.Empty() {
				h := q.Pop().(int)
				i = h + 1
				k--
				j++
			}
			for j < len(A) && k < K {
				if A[j] == 0 {
					q.Push(j)
					A[j] = 1
					k++
				}
				j++
			}
			j--
			for ; j+1 < len(A) && A[j+1] == 1; j++ {}
			ans = MaxInt(ans, j-i+1)
		}

	}
	return ans
}
