package leetcode_go

/**
剑指 Offer 58 - I. 翻转单词顺序
难度：简单
输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。
为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。
解法1：用递归，每层递归parse一个word（递归相较于迭代，占内存，且耗时（如果迭代比递归更简单的话）
*/
func firstWord(s string) (word string, i int) {
	n := len(s)
	for ; i < n; i++ {
		if s[i] == ' ' {
			break
		}
	}
	return s[:i], i
}

func isSpace(s string) bool {
	for _, v := range s {
		if v != ' ' {
			return false
		}
	}
	return true
}

func reverseWords(s string) string {
	//if isSpace(s) {
	//	return ""
	//}
	n := len(s) //len不耗时间（并不是遍历
	if n == 0 {
		return ""
	}
	//ts := strings.TrimSpace(s)
	var j int
	for ; j < n; j++ {
		if s[j] != ' ' {
			break
		}
	}
	ts := s[j:]
	word, i := firstWord(ts)
	r := reverseWords(ts[i:])
	if r == "" {
		return word
	}
	return r + " " + word
}

/**
剑指 Offer 58 - II. 左旋转字符串
难度：简单
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。
请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。
*/
func reverseLeftWords(s string, n int) string {
	s += s[:n]
	return s[n:]
}

/**
剑指 Offer 59 - I. 滑动窗口的最大值
难度：简单
给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。
1 ≤ k ≤ 输入数组的大小。
解法：k=1时，相当于findmax；使用单调队列，队首为最大元素；队列中存储的是“最大候选”，即已经是最大，及可能是最大的元素；
即：当前最大元素，及当前元素右边的（在窗口内的）递减序列；最大元素左边的元素不能是最大元素了（因为在cur max被划掉之前他们先被划掉）
push新元素j时，将队尾中小于j的删去；由于队列中只保存比j大的元素，故划掉的元素可能不在队列中，又或者是后来被划掉后，后来加入的同值元素
因此队列要保存每个元素的下标：对待删元素（被滑掉）i，if i >= max then delete else pass
问题：遍历结束条件
*/

func maxSlidingWindow(nums []int, k int) []int {
	if k <= 1 {
		return nums
	}
	res := make([]int, len(nums)-k+1)
	q := NewMonoDequeEleWithIndex(k)
	for i := 0; i < k; i++ {
		q.Push(nums[i], i)
	}
	res[0] = q.Front()
	h := 1
	//slide
	i, j := 0, k-1
	for j < len(nums)-1 {
		//TODO 不能根据值来删元素，要根据下标（可能i已经在队列中被删去，而在后续又有同值插入队列，在移动时，删除i而将后面的同值删掉了
		// 或者，外部留一个最左保留下标，以来判断是否需要删除元素
		//滑动：先删，再前移
		if nums[i] == q.Front() {
			q.Pop()
		} else {
			q.Delete(i)
		}
		i++
		j++
		q.Push(nums[j], j)
		res[h] = q.Front()
		h++
	}
	return res
}

/**
剑指 Offer 59 - II. 队列的最大值
难度：简单
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
若队列为空，pop_front 和 max_value 需要返回 -1
monotonous dequeue for int in c order
*/
type MaxQueue struct {
	c    []int
	size int
	m    *MonoDequeInt
}

func Constructor() MaxQueue {
	return MaxQueue{
		c:    make([]int, 8),
		size: 0,
		m: NewMonoDequeInt(8),
	}
}

func (this *MaxQueue) Max_value() int {
	if this.size == 0 {
		return -1
	}
	return this.m.Front()
}

func (this *MaxQueue) Push_back(value int) {
	if this.size < len(this.c) {
		this.c[this.size] = value
	} else {
		this.c = append(this.c, value)
	}
	this.size++
	this.m.Push(value)
}

func (this *MaxQueue) Pop_front() int {
	if this.size == 0 {
		return -1
	}
	r := this.c[0]
	this.c = this.c[1:]
	this.size--
	this.m.Delete(r)
	return r
}
