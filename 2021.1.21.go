package leetcode_go

import "sync"

/**
 * 46. 全排列
 * 给定一个 没有重复 数字的序列，返回其所有可能的全排列。
 * [1,2,3]
 * [
 * 	[1,2,3],
 * 	[1,3,2],
 * 	[2,1,3],
 * 	[2,3,1],
 * 	[3,1,2],
 * 	[3,2,1]
 * ]
 */

//效率很差（非dfs）
func permute(nums []int) [][]int {
	var r [][]int
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return append(r, nums)
	}
	for i := 0; i < len(nums); i++ {
		subr := permute(except(nums, i))
		for j := 0; j < len(subr); j++ {
			p := append([]int{nums[i]}, subr[j]...)
			r = append(r, p)
		}
	}
	return r
}

func except(nums []int, i int) []int {
	r := make([]int, len(nums) - 1)
	//copy不会自动增加切片大小，copy的数量为 min{len(dst), len(src)}
	n := copy(r, nums[:i])
	if i + 1 < len(nums) {
		copy(r[n:], nums[i + 1:])
	}
	return r
}

/**
 * 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗：2.5 MB, 在所有 Go 提交中击败了99.07%的用户
 */
func permute1(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return [][]int{nums}
	}
	c := 1
	for i := 2; i <= len(nums); i++{
		c *= i
	}
	r := make([][]int, 0, c)
	wg := sync.WaitGroup{}
	for i, v := range nums {
		wg.Add(1)
		go func(i, v int) {
			visited := NewBitMap(len(nums))
			visited.Set(i)
			p := make([]int, 0, len(nums))
			p = append(p, v)
			//TODO 切片传值不是引用，而是值传递？
			dfs(nums, p, &r, visited)
			wg.Done()
		}(i, v)
	}
	wg.Wait()
	return r
}

func dfs(nums []int, p []int, r *[][]int, visited *Bitmap) {
	for i, v := range nums {
		if !visited.Get(i) {
			visited.Set(i)
			p = append(p, v)
			if len(p) == cap(p) {
				pp := make([]int, cap(p))
				copy(pp, p)
				*r = append(*r, pp)
				p = p[:len(p) - 1]
				visited.Set(i)
				return
			} else {
				dfs(nums, p, r, visited)
				p = p[:len(p) - 1]
				visited.Set(i)
			}

		}
	}
}





