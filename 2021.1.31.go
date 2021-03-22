package leetcode_go

/**
 * 415. 字符串相加
 * 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
 * num1 和num2的长度都小于 5100
 * num1 和num2 都只包含数字0-9
 * num1 和num2 都不包含任何前导零
 * 你不能使用任何內建 BigInteger 库，也不能直接将输入的字符串转换为整数形式
 */

func addStrings(num1 string, num2 string) string {
	//to guranteen len(num1) > len(num2)
	if len(num1) < len(num2) {
		tmp := num1
		num1 = num2
		num2 = tmp
	}
	ans := make([]byte, len(num1)+1)
	up := false
	diff := len(num1) - len(num2)
	for i := len(num1); i > diff; i-- {
		//i - 1是由于ans比num1多一位，以ans为基准时，对num1访问需-1
		sum := Ctoi(num1[i-1]) + Ctoi(num2[i-diff-1])
		if up {
			sum += 1
		}
		if sum >= 10 {
			up = true
			sum = sum - 10
		} else {
			up = false
		}
		ans[i] = Itoc(sum)
	}

	for i := diff - 1; i >= 0; i-- {
		sum := Ctoi(num1[i])
		if up {
			sum += 1
		}
		if sum >= 10 {
			up = true
			sum = sum - 10
		} else {
			up = false
		}
		ans[i+1] = Itoc(sum)
	}
	if up {
		ans[0] = Itoc(1)
	} else {
		ans = ans[1:]
	}
	return string(ans)

}


/**
* 560. 和为K的子数组
* 给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
首先要有一种快速计算子数组和的方法（而不是针对每个子数组都进行累加），由此可以用前缀和（前缀和的作用就是这个）
记sum(i,j)=pre[j]-pre[i-1]，j>i-1。当遍历到j，以j为结尾的子数组中，符合sum(i,j)=k的情况，
即符合pre[j]-k的pre[i-1]的次数（且i-1<j，因此不需要考虑j后面的情况）。
因此，只需要维护一个preCnt数组记录出现的前缀和的次数即可。
由于数据范围较大，且索引可能出现负数，因此使用map。

统计次数，便不能只是简单的单向迭代（如求最大值只需要简单迭代就好了），还需要维护必要的状态
*/

func subarraySum(nums []int, k int) int {

	pre := 0
	ans := 0
	preCnt := make(map[int]int)
	preCnt[0] = 1
	for i := 0; i < len(nums); i++ {
		pre = pre + nums[i]
		if v, ok := preCnt[pre-k]; ok {
			ans += v
		}
		preCnt[pre]++
	}
	return ans
}

/**
 * 974. 和可被 K 整除的子数组
 * 给定一个整数数组 A，返回其中元素之和可被 K 整除的（连续、非空）子数组的数目。
 * 解法：首先由前缀和，sum(i,j)=pre[j]-pre[i-1]
 * 又sum(i,j)%k = ((pre[j]%k) - pre[i-1]))%k，因此，要求符合模k的和的数量，
 * 可转而求：对mod=pre[j]%k，若mod>=0，找对应的pre[i-1]%k=mod，k-mod的次数
 *						 若mod<0，找对应的pre[i-1]%k=mod，k+mod的次数
  * 这种题的关键是要在前缀和公式的基础上，由题目关系，找到要统计的状态，并统计其次数。
*/

func subarraysDivByK(A []int, K int) int {
	modCnt := make(map[int]int)
	ans := 0
	pre := 0
	modCnt[0] = 1
	for i := 0; i < len(A); i++ {
		pre += A[i]
		mod := pre % K
		//K=5, mod=4，需要减去（注意不是加！）4或者-1
		if v, ok := modCnt[mod]; ok {
			ans += v
		}
		if mod > 0 {
			if v, ok := modCnt[mod-K]; ok {
				ans += v
			}
		} else {
			if v, ok := modCnt[K+mod]; ok {
				ans += v
			}
		}
		modCnt[mod]++
	}
	return ans
}
