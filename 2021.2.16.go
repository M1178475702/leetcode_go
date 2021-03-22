package leetcode_go

import (
	"math"
	"strconv"
	"strings"
)

/**
 * 22. 括号生成
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 */

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	var r []string
	gpdfs(&r, "", true, 0, 0, 2*n)
	gpdfs(&r, "", false, 0, 0, 2*n)
	return r
}

//if there is MODIFiCATION in func block, the ptr is needed
func gpdfs(r *[]string, cur string, left bool, lc, rc, max int) {
	if left {
		cur += "("
		lc++
	} else {
		cur += ")"
		rc++
	}
	if len(cur) == max {
		if isBalance(cur) {
			*r = append(*r, cur)
			return
		}
	}

	if lc < max/2 {
		gpdfs(r, cur, true, lc, rc, max)
	}
	if rc < max/2 {
		gpdfs(r, cur, false, lc, rc, max)
	}
}

func isBalance(str string) bool {
	s := NewStack(len(str) / 2)
	for _, c := range str {
		if c == '(' {
			s.Push(c)
		} else {
			if s.Empty() {
				return false
			} else {
				s.Pop()
			}
		}
	}
	return s.Empty()
}

/**
 * 394. 字符串解码
 */
func decodeString(s string) string {
	var ans strings.Builder
	for i := 0; i < len(s); i++ {
		if IsDigit(s[i]) {
			b := i
			for i++; i < len(s) && IsDigit(s[i]); i++ {
			}
			cnt, _ := strconv.Atoi(s[b:i])
			//使用栈来判定括号表达式的结束
			st := NewStack(1)
			st.Push(s[i])
			var j int
			for j = i + 2; j < len(s); j++ {
				if s[j] == '[' {
					st.Push(s[j])
				} else if s[j] == ']' {
					st.Pop()
					if st.Empty() {
						break
					}
				}
			}
			subs := decodeString(s[i+1 : j])
			for k := 0; k < cnt; k++ {
				ans.WriteString(subs)
			}
			i = j
		} else {
			ans.WriteByte(s[i])
		}
	}
	return ans.String()
}

/**
 * 236. 二叉树的最近公共祖先
 * 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
 * 解法：先用迭代后续得到路径，然后比较
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	p2p := PathToTreeNode(root, p)
	p2q := PathToTreeNode(root, q)

	p2pi := 0
	var ln, pi, qi *TreeNode
	for p2pi < len(p2p) && p2pi < len(p2q) {
		pi, qi = p2p[p2pi], p2q[p2pi]
		if pi != qi {
			break
		}
		ln = pi
		p2pi++
	}
	return ln

}

func PathToTreeNode(root, node *TreeNode) []*TreeNode {
	s := NewStack(1)
	iter := root
	var ans []*TreeNode
	v := make(map[*TreeNode]struct{})
	for !s.Empty() || iter != nil {
		if iter != nil {
			if iter == node {
				for !s.Empty() {
					ans = append(ans, s.Shift().(*TreeNode))
				}
				ans = append(ans, iter)
				break
			}
			s.Push(iter)
			iter = iter.Left
		} else {
			iter = s.Pop().(*TreeNode)
			if _, ok := v[iter]; ok {
				//do
				if iter == node {
					for !s.Empty() {
						ans = append(ans, s.Shift().(*TreeNode))
					}
					ans = append(ans, iter)
					break
				}
				iter = nil
			} else {
				v[iter] = struct{}{}
				s.Push(iter)
				iter = iter.Right
			}
		}
	}
	return ans
}

/**
 * 23. 合并 K 个升序链表
 * 给你一个链表数组，每个链表都已经按升序排列。
 * 请你将所有链表合并到一个升序链表中，返回合并后的链表。
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		lists[1] = mergeTwoLists(lists[0], lists[1])
		lists = lists[1:]
	}
	return lists[0]
}

/**
 * 剑指 Offer 47. 礼物的最大价值
 * 在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，
 * 并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
 * 解法：动态规划：记dp[i][j]为到达(i,j)时的最大值，其取决于是从左还是从上走过来的，即
 * dp[i][j] = grid[i][j] + max(dp[i][j-1], dp[i-1][j])
 * 空间优化：由于grid只访问当前(i,j)，已访问过的不会再访问，因此dp的值可以记录在grid上，空间降到O(1)
 * 如果不想修改原来数组，由于从迭代的顺序上，计算结果时只需访问上一行和左边的一个值，因此空间可以降到O(n+1)
 */
func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	dirc := func(i, j int) [][]int {
		var ps [][]int
		if i-1 >= 0 {
			ps = append(ps, []int{i - 1, j})
		}
		if j-1 >= 0 {
			ps = append(ps, []int{i, j - 1})
		}
		return ps
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ds := dirc(i, j)
			dp[i][j] = grid[i][j]
			for _, v := range ds {
				if grid[i][j]+dp[v[0]][v[1]] > dp[i][j] {
					dp[i][j] = grid[i][j] + dp[v[0]][v[1]]
				}
			}
		}
	}
	return dp[m-1][n-1]
}

/**
 * 474. 一和零
 * 给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
 * 请你找出并返回 strs 的最大子集的大小，该子集中 最多 有 m 个 0 和 n 个 1 。
 * 如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集
 * dp[k][i][j]表示考虑是否放入第k个字符串，i个0，j个1时的最大集合元素数
 * 因此dp[k][i][j] = max(1+dp[i-c0][j-len(s)+c0], dp[i][j])
 * 空间优化思路同背包，更新当前值时，使用的总是上一行，之前（<i,j)的值，因此可以自右向左遍历，因此可以只保留一个i*j的数组
 * 背包题：公式为dp[k][res]；k代表放第k个物品，res表示资源（如容量（这种题总是从最终结果向前推，如第n个结果如何从第n-1个结果推得
 */
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	zeroc := func(s string) int {
		cnt := 0
		for _, v := range s {
			if v == '0' {
				cnt++
			}
		}
		return cnt
	}
	for k := 1; k <= len(strs); k++ {
		c0 := zeroc(strs[k-1]) //count 0
		c1 := len(strs[k-1])-c0 //count 1
		for i := m; i >= c0; i-- {   //只管装的下的情况，提前结束
			for j := n; j >= c1; j-- {
				dp[i][j] = MaxInt(1+dp[i-c0][j-c1], dp[i][j])
				//不需要管装不下的情况，因此装不下默认是原来的值，由于数组时原来的数组，所以不需要修改
			}
		}
	}
	return dp[m][n]
}

/**
 * 300. 最长递增子序列
 * 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
 * 子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
 * 例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
 * ml[i]记录以nums[i]结尾的子数组的最长子序列长度
 * ml[i]的长度为，在0~i-1中小于nums[i]的最长子序列+1
 * 因此需要遍历0~i-1查找以比nums[i]小的数结尾的序列长度。这里可以优化
 * 用tail[k]记录k+1长度的子序列的结尾数。如果保证该序列时递增的，则上述查找过程，可以用二分查找tail中比nums[i]小的数
 * 然后返回的下标+1即对应最长序列长度，切由于是二分，其必为最大长度。
 * 维护tail：时tail[i]的元素尽量小，即如果有相同长度的序列，切结尾数比当前小，则更新
 * 证明tail是递增序列：（反证）如果d[j]≥d[i]且j<i，我们考虑从长度为i的最长上升子序列的末尾（从右边）删除i−j 个元素，
 * 那么这个序列长度变为j，且第j个元素x（末尾元素）必然小于d[i]，也就小于d[j]。那么此时我们得到了一个以x结尾，j长度的子序列、
 * 而x<d[j]，此时d[j]应该是x。因而矛盾。
 *
 */
func lengthOfLIS(nums []int) int {
	ml, tail := make([]int, len(nums)), make([]int, len(nums))//max len
	ml[0] = 1
	max := 1
	tail[0] = nums[0]
	for i := 1; i < len(tail); i++ {
		tail[i] = math.MaxInt32
	}
	for i := 1; i < len(nums); i++ {
		ml[i] = 1
		j := BinarySearchLeft(tail[:max], nums[i], nil)
		//tail[j]可能大于（tail中没有比nums[i]小的,j为0）等于（之前出现过nums[i]）nums[i]
		if tail[j] >= nums[i] {
			j--
		}
		ml[i] = j+2
		if tail[ml[i]-1] > nums[i] {
			tail[ml[i]-1] = nums[i]
		}
		max = MaxInt(max, ml[i])
	}
	return max
}

















