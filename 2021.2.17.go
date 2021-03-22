package leetcode_go

import (
	"strings"
)

/**
 * 978. 最长湍流子数组
 * 当 A的子数组A[i], A[i+1], ..., A[j]满足下列条件时，我们称其为湍流子数组：
 * 若i <= k < j，当 k为奇数时，A[k] > A[k+1]，且当 k 为偶数时，A[k] < A[k+1]；
 * 或 若i <= k < j，当 k 为偶数时，A[k] > A[k+1]，且当 k为奇数时，A[k] < A[k+1]。
 * 也就是说，如果比较符号在子数组中的每个相邻元素对之间翻转，则该子数组是湍流子数组。
 * 返回 A 的最大湍流子数组的长度。
 * 解法：ml[i][2]记作以i结尾，以<(0)还是>(1)arr[i-1]结尾的最大子数组长度
 * 当arr[i-1]>arr[i]，ml[i][1] = ml[i][0] + 1
 * 当arr[i-1]>arr[i]，ml[i][0] = ml[i][1] + 1
 */

func maxTurbulenceSize(arr []int) int {
	ml := make([]int, 2)
	max := 1
	ml[0], ml[1] = 1, 1
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			ml[1] = ml[0] + 1
			ml[0] = 1
		} else if arr[i-1] < arr[i] {
			ml[0] = ml[1] + 1
			ml[1] = 1
		} else {
			ml[1] = 1
			ml[0] = 1
		}
		max = MaxInt(max, MaxInt(ml[0], ml[1]))
	}
	return max
}

/**
 * 1143. 最长公共子序列
 */

//暴力
//枚举short所有子序列，然后判断该子序列是否属于long
//怎么枚举所有子序列？回溯？按起始字符和长度？
func longestCommonSubsequence(text1 string, text2 string) int {
	long, short := text1, text2
	if len(text1) < len(text2) {
		long, short = text2, text1
	}
	max := 0
	subSeq(short, "", 0, func(seq string) {
		if isSubSeq(long, seq) {
			max = MaxInt(max, len(seq))
		}
	})
	return max
}

func isSubSeq(src, seq string) bool {
	if len(seq) > len(src) {
		return false
	}
	j := 0
	for i := 0; i < len(seq); i++ {
		k := j
		for ; k < len(src); k++ {
			if seq[i] == src[k] {
				j = k + 1
				break
			}
		}
		if k == len(src) {
			return false
		}
	}
	return true
}

func subSeq(src, cur string, i int, do func(seq string)) {
	if i >= len(src) {
		return
	}
	do(cur + src[i:i+1])
	//不可能往回回溯（即走过的路的状态是固定的，不会再变），所以不需要visited数组
	subSeq(src, cur+src[i:i+1], i+1, do)
	subSeq(src, cur, i+1, do)
}

//dp
//两个字符串的dp，通常是dp[i][j]代表长度为i的s1和长度为j的s2时的状态（且状态一般为最终所求结果）
//dp推导公式，一般由后向前推导，即dp[i][j]的值取决于前面的哪些值，与前面的值的关系
//想通这个，状态转移公式即可得出
func longestCommonSubsequence1(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = MaxInt(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

/**
 * 51. N 皇后
 * n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 * 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
 * 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
 */

func solveNQueens(n int) [][]string {
	ans := make([][]string, 0, 1)
	//每行决策为：在哪个位置放皇后
	grid := make([]string, 0, n)
	nqdfs(&ans, &grid, n)
	return ans
}

//结果集，状态集（当前位置，决策因素）
func nqdfs(ans *[][]string, grid *[]string, n int) {
	cur := len(*grid) //将要置放的行号（刚好等于长度）
	if cur >= n {     //当前行
		//copy grid
		ret := make([]string, n)
		copy(ret, *grid)
		*ans = append(*ans, ret)
		return
	}
	pos := nqGetSettablePos(*grid, n, cur)
	for _, v := range pos {
		*grid = append(*grid, setQ(n, v))
		nqdfs(ans, grid, n)
		*grid = (*grid)[:cur]
	}
}

func setQ(n, i int) string {
	b := strings.Builder{}
	for k := 0; k < n; k++ {
		if k != i {
			b.WriteByte('.')
		} else {
			b.WriteByte('Q')
		}
	}
	return b.String()
}

//还可以优化，使用空间记录已经放过的行，列，然后快速得到可以置放的位置
func nqGetSettablePos(grid []string, n, cur int) []int {
	var ret []int
	for i := 0; i < n; i++ {
		//检测上方
		canset := true
		for j := cur - 1; j >= 0; j-- {
			if grid[j][i] == 'Q' {
				canset = false
				break
			}
		}
		if !canset {
			continue
		}
		//左上
		for j, k := cur-1, i-1; j >= 0 && k >= 0; j-- {
			if grid[j][k] == 'Q' {
				canset = false
				break
			}
			k--
		}
		if !canset {
			continue
		}
		//右上
		for j, k := cur-1, i+1; j >= 0 && k < n; j-- {
			if grid[j][k] == 'Q' {
				canset = false
				break
			}
			k++
		}
		if canset {
			ret = append(ret, i)
		}
	}
	return ret
}

/**
 * 17. 电话号码的字母组合
 * 给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 */
//construct digit-char map
var dcm = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	if len(digits) == 1 {
		return strings.Split(dcm[Ctoi(digits[0])], "")
	}
	ret := letterCombinations(digits[1:])
	ans := make([]string, len(ret)*len(dcm[Ctoi(digits[0])]))
	for _, r := range ret {
		cs := dcm[Ctoi(digits[0])]
		for i, _ := range cs {
			b := strings.Builder{}
			b.WriteByte(cs[i])
			b.WriteString(r)
			ans = append(ans, b.String())
		}
	}
	return ans
}

/**
 * 37. 解数独
 * 编写一个程序，通过填充空格来解决数独问题。
 * 一个数独的解法需遵循如下规则：
 * 数字1-9在每一行只能出现一次。
 * 数字1-9在每一列只能出现一次。
 * 数字1-9在每一个以粗实线分隔的3x3宫内只能出现一次。
 * 空白格用'.'表示。
 * 行/列各9格；
 */

func solveSudoku(board [][]byte) {
	pos := sdNextPos(board, 0, -1)
	sddfs(board, pos[0], pos[1])
}

//回溯时可以通过bool来提前终结递归的
//pos, ds会重复计算，且会重复浪费空间。
//可以提前把所有要填的空挑出来，相当于将所有可能遍历的空及其可能状态（数字）挑出来
//然后只遍历状态即可。不需要费力找下一个状态。
//对ds，也可以先遍历当前可能的状态再决定之后，没必要先将所有可能的状态枚举出来再遍历（尽管代码上看起来很简洁）
func sddfs(board [][]byte, i, j int) bool {
	ds := sdGetSettableDigit(board, i, j)
	pos := sdNextPos(board, i, j)
	if len(ds) == 0 {
		return len(pos) == 0
	}
	for _, d := range ds {
		board[i][j] = d
		if len(pos) != 0 {
			if sddfs(board, pos[0], pos[1]){
				return true
			} else {
				board[i][j] = '.'
			}
		} else {
			return true
		}
	}
	return false
}

func sdNextPos(board [][]byte, i, j int) []int {
	for {
		if j == 8 {
			if i == 8 {
				return []int{}
			}
			i++
			j = 0
		} else {
			j++
		}
		if board[i][j] == '.' {
			break
		}
	}
	return []int{i,j}
}

func sdGetSettableDigit(board [][]byte, i, j int) (ds []byte) {
	ds = make([]byte, 0, 4) //cap总是会被扩张到8
	gl, gr, gt, gb := (j/3)*3, (j/3+1)*3, (i/3)*3, (i/3+1)*3
	//不声明byte，'1'会被认定为int32
	for d := byte('1'); d <= '9'; d++ {
		canset := true
		//check row and column
		for k := 0; k < 9; k++ {
			if board[i][k] == d || board[k][j] == d {
				canset = false
				break
			}
		}
		if !canset{
			continue
		}
		//check grid
		//left, right, top, bottom
		//TODO 一个很蛋疼的bug就是gt,gb,与gl,gr整反了=。=
		for k := gt; k < gb; k++ {
			for h := gl; h < gr; h++ {
				if board[k][h] == d {
					canset = false
					break
				}
			}
		}
		if canset {
			ds = append(ds, d)
		}
	}
	return ds
}
