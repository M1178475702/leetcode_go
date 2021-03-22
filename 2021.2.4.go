package leetcode_go

/**
 * 200. 岛屿数量
 * 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
 * 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
 * 此外，你可以假设该网格的四条边均被水包围。
 * 解法1：以出现的1为种子，用BFS搜索整个区域，将遇到的1变为0。每次搜索结束即为一个岛屿
 * 效率很低：时间：5.76%；空间：8.23%
 */



func numIslands(grid [][]byte) int {
	q := NewSQueue(len(grid[0]))
	m := len(grid)
	n := len(grid[0])
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				q.Push(&PairInt{
					first:  i,
					second: j,
				})
				for !q.Empty() {
					p := q.Pop().(*PairInt)
					for _, d := range direction(p.first, p.second, m, n) {
						if grid[d.first][d.second] == '1' {
							grid[d.first][d.second] = '0'
							q.Push(d)
						}
					}
				}
				ans++
			}
		}
	}
	return ans
}



//使用并查集合并
/**
	如果一个位置是1，则将旁边的1都合并到一起。最后统计连通分量（不同根个数）个数。
	合并过程要路径压缩。而对于两个树合并，要将另一个树上的节点全部指向另一个树的根节点，则需要知道该树的所有子节点。
	而对于数组形式，只能遍历全部数组（如果是树，也要双向树节点），相比BFS麻烦许多。
	单次操作的时间复杂度即为 O（MN）！
	所谓并查集，除了并，还有查！
	不写了。。
 */
func numIslands1(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	ans := 0
	un := make([][]bool, m)
	for i, _ := range un {
		un[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ds := direction(i, j, m, n)
				for _, d := range ds {
					if grid[d.first][d.second] == '1' && un[d.first][d.second] {
						un[i][j] = true
						break
					}
				}
				if !un[i][j] {
					un[i][j] = true
					ans++
				}
				for _, d := range ds {
					if grid[d.first][d.second] == '1' {
						un[d.first][d.second] = true
					}
				}
			}
		}
	}
	return ans
}