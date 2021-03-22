package leetcode_go

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * 121. 买卖股票的最佳时机
 * 解法：给出基本计算公式：r(i, j) = p[j] - p[i]
 * 对p[j]要想知道以其为卖出的那天的最大收益，则需要找到最小p[i]
 * 因此可以用pre[j-1]记录0-j-1中最小的一天。
 * 这类题，乍一看只能for i,j的暴力计算。因为有两个点需要分别选取。
 * 可以以其中一个点为基准（如j），然后分析基本计算公式，看有哪些状态可以重复利用
 */

func maxProfit1(prices []int) int {
	ans := 0
	pre := make([]int, len(prices))
	pre[0] = prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < pre[i-1] {
			pre[i] = prices[i]
		} else {
			pre[i] = pre[i-1]
		}
	}
	for i := 1; i < len(prices); i++ {
		if prices[i]-pre[i-1] > ans {
			ans = prices[i] - pre[i-1]
		}
	}
	return ans
}

//优化：由于只访问pre[i-1]，所以可以边计算pre，边计算ans
//同时只保存pre[i-1]
func maxProfit2(prices []int) int {
	ans := 0
	pre := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < pre {
			pre = prices[i]
		}
		if prices[i]-pre > ans {
			ans = prices[i] - pre
		}
	}
	return ans
}

/**
 * 买卖股票的最佳时机 II
 */
//暴力，回溯法，每个节点有三中操作
//回溯法可用于：每个节点有多种操作，每一步操作对后面的操作有影响，即有多种结果（或操作序列）
func maxProfitTwo(prices []int) int {
	//买
	p1 := mpdfs(prices, 0, 1, 0, true)
	//不操作
	p2 := mpdfs(prices, -1, 1, 0, false)
	if p1 < p2 {
		return p2
	}
	return p1
}

func mpdfs(prices []int, buy int, cur int, curp int, stock bool) int {
	if cur == len(prices) {
		//最后返回收入
		return curp
	}
	var p1, p2 int
	if stock {
		//不操作
		p1 = mpdfs(prices, buy, cur+1, curp, true)
		//卖
		curp += prices[cur] - prices[buy]
		p2 = mpdfs(prices, -1, cur+1, curp, false)
	} else {
		//不操作
		p1 = mpdfs(prices, -1, cur+1, curp, false)
		//买
		p2 = mpdfs(prices, cur, cur+1, curp, true)
	}
	if p1 < p2 {
		return p2
	}
	return p1
}

//动态规划：dp[i][j]表示第i天，j状态的最大现金数（初始为0）（0未买入（cash）状态，1表示买入（stock））
//dp[i][0]为第i-1天卖出stock->cash（s2c）与不操作cash->cash之间的最大值
//dp[i][0]为第i-1天买入cash->stock（c2s）与不操作stock->stock之间的最大值

func maxProfitTwo1(prices []int) int {
	dp := make([][2]int, len(prices))
	dp[0][0], dp[0][1] = 0, 0-prices[0]
	for i := 1; i < len(prices); i++ {
		s2c := dp[i-1][1] + prices[i]
		if s2c > dp[i-1][0] {
			dp[i][0] = s2c
		} else {
			dp[i][0] = dp[i-1][0]
		}
		c2s := dp[i-1][0] - prices[i]
		if c2s > dp[i-1][1] {
			dp[i][1] = c2s
		} else {
			dp[i][1] = dp[i-1][1]
		}
	}
	if dp[len(prices)-1][0] > dp[len(prices)-1][1] {
		return dp[len(prices)-1][0]
	}
	return dp[len(prices)-1][1]
}

//for the sake of only using the dp[i-1], so we just record dp[i-1]
func maxProfitTwo2(prices []int) int {
	//dp := make([][2]int, len(prices))
	last := make([]int, 2)
	last[0], last[1] = 0, 0-prices[0]
	for i := 1; i < len(prices); i++ {
		s2c := last[1] + prices[i]
		if s2c > last[0] {
			last[0] = s2c
		} else {
			last[0] = last[0]
		}
		c2s := last[0] - prices[i]
		if c2s > last[1] {
			last[1] = c2s
		} else {
			last[1] = last[1]
		}
	}
	if last[0] > last[1] {
		return last[0]
	}
	return last[1]
}

/**
 * 199. 二叉树的右视图
 * 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
 * 解法；如果是深搜，则需要先尽可能访问右节点，同时记录已访问过的层数。之后访问到的之前未访问到的层的第一个节点就是最右节点
 */

func rightSideView(root *TreeNode) []int {
	//若为数组，按层找其最右节点，找映射公式
	//对链表，层次遍历
	if root == nil {
		return []int{}
	}
	var ans []int
	barrier := &TreeNode{} //nil 不能转指针，因此必须显式创建一个barrier
	q := NewSQueue(10)
	q.Push(barrier)
	q.Push(root)
	for !q.Empty() {
		node := q.Pop().(*TreeNode)
		if node == barrier {
			if !q.Empty() {
				q.Push(barrier)
				node = q.Pop().(*TreeNode)
				ans = append(ans, node.Val)
			} else {
				break
			}
		}
		if node.Right != nil {
			q.Push(node.Right)
		}
		if node.Left != nil {
			q.Push(node.Left)
		}

	}
	return ans
}

/**
 * 79. 单词搜索
 * 给定一个二维网格和一个单词，找出该单词是否存在于网格中。
 * 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
 */

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	v := NewBitMap(m * n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				v.Set(i*n + j)
				if existDfs(&board, i, j, v, word[1:]) {
					return true
				}
				v.Set(i*n + j)
			}
		}
	}
	return false
}

func existDfs(board *[][]byte, i, j int, v *Bitmap, word string) bool {
	if len(word) == 0 {
		return true
	}
	m, n := len(*board), len((*board)[0])
	ps := direction(i, j, m, n)
	for _, p := range ps {
		if !v.Get(p.first*n+p.second) && (*board)[p.first][p.second] == word[0] {
			v.Set(p.first*n + p.second)
			if existDfs(board, p.first, p.second, v, word[1:]) {
				return true
			} else {
				v.Set(p.first*n + p.second)
			}
		}
	}
	return false
}

/**
 * 297. 二叉树的序列化与反序列化
 * 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
 * 解法：
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func CodecConstructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
/**
 * 使用层次遍历同时使用下标标记，需要记录对应的i。从队列取出节点时，需要知道其对应下标，才能计算出对应的子节点下标
 */
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "x"
	}
	return fmt.Sprintf("%v,%v,%v", root.Val, this.serialize(root.Left), this.serialize(root.Right))
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	return this.createTree(&vals)
}

func (this *Codec) createTree(vals *[]string) *TreeNode {
	if len(*vals) == 0 {
		return nil
	}
	if (*vals)[0] == "x" {
		//将遍历过的节点消除掉，这样剩下的节点自然就是剩下子树的节点
		//要使用指针修改
		*vals = (*vals)[1:]
		return nil
	}
	v, _ := strconv.Atoi((*vals)[0])
	node := &TreeNode{Val: v}
	*vals = (*vals)[1:]
	node.Left = this.createTree(vals)
	node.Right = this.createTree(vals)
	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
