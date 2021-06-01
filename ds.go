/**
@Title ds
@Description leetcode data structure
@Author	Bakatora
@CreateAt 2021.1.16
*/

package leetcode_go

import (
	"container/heap"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type StackTreeNode struct {
	c    []*TreeNode
	size int
}

func NewStackTreeNode(cap int) *StackTreeNode {
	return &StackTreeNode{
		c: make([]*TreeNode, cap),
	}
}

func (s *StackTreeNode) Push(e *TreeNode) {
	n := len(s.c)
	if s.size < n {
		s.c[s.size] = e
	} else {
		s.c = append(s.c, e)
	}
	s.size++
}

func (s *StackTreeNode) Pop() *TreeNode {
	if s.Empty() {
		panic("stack is empty")
	}
	s.size--
	return s.c[s.size]
}

func (s *StackTreeNode) Empty() bool {
	return s.size == 0
}


type Stack struct {
	c    []interface{}
	size int
}

func NewStack(cap int) *Stack {
	return &Stack{
		c: make([]interface{}, cap),
	}
}

func (s *Stack) Push(e interface{}) {
	n := len(s.c)
	if s.size < n {
		s.c[s.size] = e
	} else {
		s.c = append(s.c, e)
	}
	s.size++
}

func (s *Stack) Pop() interface{} {
	if s.Empty() {
		panic("stack is empty")
	}
	s.size--
	return s.c[s.size]
}

func (s *Stack) Shift() interface{} {
	if s.Empty() {
		panic("stack is empty")
	}
	s.size--
	r := s.c[0]
	s.c = s.c[1:]
	return r
}

func (s *Stack) Empty() bool {
	return s.size == 0
}


/**
单调双向队列，队列内保持递减（从队首到尾）
To adapt to the question of offer 59, the elements have an index in the slice nums
*/

type EleWithIndex struct {
	ele   int
	index int
}

type MonoDequeEleWithIndex struct {
	c    []EleWithIndex
	size int
}

func NewMonoDequeEleWithIndex(cap int) *MonoDequeEleWithIndex {
	return &MonoDequeEleWithIndex{
		c:    make([]EleWithIndex, cap),
		size: 0,
	}
}

func (m *MonoDequeEleWithIndex) Push(e, index int) {
	ele := EleWithIndex{
		ele:   e,
		index: index,
	}

	if m.Empty() {
		m.c[0] = ele
		m.size++
		return
	}
	for i := m.size - 1; i >= 0; i-- {
		if m.c[i].ele < e {
			m.size--
		} else {
			break
		}
	}
	if m.size < len(m.c) {
		m.c[m.size] = ele
	} else {
		m.c = append(m.c, ele)
	}
	m.size++
}

func (m *MonoDequeEleWithIndex) Front() int {
	return m.c[0].ele
}

func (m *MonoDequeEleWithIndex) Pop() int {
	m.size--
	r := m.c[0]
	m.c = m.c[1:]
	return r.ele
}

//删除性能比较差，可以考虑使用链表（经常增删）
func (m *MonoDequeEleWithIndex) Delete(index int) {
	if m.c[0].index > index {
		return
	}
	for i, v := range m.c {
		if v.index > index {
			return
		}
		if v.index == index {
			//copy(m.c[:i], m.c[i+1:])
			//m.c = m.c[: m.size-1]
			m.c = append(m.c[:i], m.c[i+1:]...)
			m.size--
			return
		}
	}
}

func (m *MonoDequeEleWithIndex) Size() int {
	return m.size
}

func (m *MonoDequeEleWithIndex) Empty() bool {
	return m.size == 0
}

var defaultCmp = func(i, j int) int {
	if i < j {
		return -1
	} else if i == j {
		return 0
	} else {
		return 1
	}
}

func BinarySearch(nums []int, value int, cmp func(i, j int) int) int {
	if len(nums) == 0 {
		return 0
	}

	low, hi, m := 0, len(nums)-1, 0
	for low <= hi {
		m = (low + hi) / 2
		leb := cmp(nums[m], value)
		if leb == 0 {
			break
		} else if leb == -1 {
			low = m + 1
		} else {
			hi = m - 1
		}
	}
	return m
}

//若值为重复的，返回最左边的下标;若值不存在，返回小于该值的第一个数的下标
func BinarySearchLeft(nums []int, value int, cmp func(i, j int) int) int {
	if cmp == nil {
		cmp = defaultCmp
	}
	i := BinarySearch(nums, value, cmp)
	if nums[i] != value && cmp(nums[i], value) == 1 && i-1 >= 0 {
		i--
	} else {
		for i >= 0 && i+1 < len(nums) && cmp(nums[i+1], value) == 0{i--}
	}
	return i
}

//若值为重复的，返回最右边的下标；若值不存在，返回小于该值的第一个数的下标
func BinarySearchRight(nums []int, value int, cmp func(i, j int) int) int {
	if cmp == nil {
		cmp = defaultCmp
	}
	i := BinarySearch(nums, value, cmp)
	//当t小于最小值时，返回的值是大于t的
	if nums[i] != value && cmp(nums[i], value) == 1 && i-1 >= 0 {
		i--
	} else {
		for i < len(nums)-1 {
			if cmp(nums[i+1], value) == 0 {
				i++
			} else {
				break
			}
		}
	}
	return i
}

type MonoDequeInt struct {
	c    []int
	size int
}

func NewMonoDequeInt(cap int) *MonoDequeInt {
	return &MonoDequeInt{
		c:    make([]int, cap),
		size: 0,
	}
}

func (m *MonoDequeInt) Push(e int) {
	if m.size == 0 {
		if len(m.c) == 0 {
			m.c = append(m.c, e)
		} else {
			m.c[0] = e
		}
		m.size++
		return
	}
	//Using binary search to find where e should be inserted at
	//for i := m.size - 1; i >= 0; i-- {
	//	if m.c[i] < e {
	//		m.size--
	//	} else {
	//		break
	//	}
	//}
	i := BinarySearchRight(m.c[:m.size], e, func(i, j int) int {
		if i < j {
			return 1
		} else if i == j {
			return 0
		} else {
			return -1
		}
	})
	//当新插入值最大时，下标应该为0，而不是1（i为0没错，不再需要+1，单独处理下
	//i = 0的情况：新插入值最大；new < max, size=1，

	if i == 0 && e > m.c[0] {
		m.c[0] = e
		m.size = 1
		return
	} else {
		m.size = i + 2
	}

	if m.size-1 < len(m.c) {
		m.c[m.size-1] = e
	} else {
		m.c = append(m.c, e)
	}

}

func (m *MonoDequeInt) Pop() int {
	r := m.c[0]
	m.c = m.c[1:]
	m.size--
	return r
}

func (m *MonoDequeInt) Delete(e int) {
	if e == m.c[0] {
		m.Pop()
		return
	}
	//已知队列内递减，若删去值不是max，max不变，不需要删去该值
	//仅当max出队，且此时该值还没有被后来值顶掉时，其可能成为max；而按进队顺序，max一定比小于它的值先出队
	//当出队的值不是max时，
	i := BinarySearchRight(m.c[:m.size], e, func(i, j int) int {
		if i < j {
			return 1
		} else if i == j {
			return 0
		} else {
			return -1
		}
	})
	if m.c[i] == e {
		if i+1 < m.size {
			copy(m.c[i:], m.c[i+1:])
		}
		m.size--
	}
}

func (m *MonoDequeInt) Front() int {
	return m.c[0]
}

func (m MonoDequeInt) Empty() bool {
	return m.size == 0
}

type Bitmap struct {
	c []byte
}

func NewBitMap(cap int) *Bitmap {
	return &Bitmap{c: make([]byte, cap/8+1)}
}

func (b *Bitmap) Set(i int) {
	//ci := i / 8
	//bi := i % 8
	b.c[i/8] = b.c[i/8] ^ (1 << (8 - i%8 - 1))
}

func (b *Bitmap) Get(i int) bool {
	return (b.c[i/8]>>(8-i%8-1))%2 == 1
}


type Queue interface {
	Push(e interface{})
	Pop() interface{}
	Front() interface{}
	Empty() bool
}


type SimpleQueue struct {
	c []interface{}
}

//Simple queue
func NewSQueue(cap int) Queue {
	return &SimpleQueue{
		c: make([]interface{}, 0, cap),
	}
}

func (q *SimpleQueue) Push(e interface{}) {
	q.c = append(q.c, e)
}

func (q *SimpleQueue) Pop() interface{} {
	r := q.c[0]
	q.c = q.c[1:]
	return r
}

func (q *SimpleQueue) Front() interface{} {
	return q.c[0]
}

func (q *SimpleQueue) Empty() bool {
	return len(q.c) == 0
}

type PairInt struct {
	first  int
	second int
}

func direction(i, j, m, n int) (ps []*PairInt) {
	if i-1 >= 0 {
		ps = append(ps, &PairInt{first: i - 1, second: j})
	}
	if i+1 < m {
		ps = append(ps, &PairInt{first: i + 1, second: j})
	}
	if j-1 >= 0 {
		ps = append(ps, &PairInt{first: i, second: j - 1})
	}
	if j+1 < n {
		ps = append(ps, &PairInt{first: i, second: j + 1})
	}
	return ps
}

func CreateBiTree(nums []int, i int) *TreeNode {
	if len(nums) == 0 || i >= len(nums) || nums[i] == -1 {
		return nil
	}
	node := &TreeNode{Val: nums[i]}
	node.Left = CreateBiTree(nums, 2*i+1)
	node.Right = CreateBiTree(nums, 2*i+2)
	return node
}

func FindTreeNode(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Val == val {
		return node
	}
	l := FindTreeNode(node.Left, val)
	if l != nil {
		return l
	}
	r := FindTreeNode(node.Right, val)
	if r != nil {
		return r
	}
	return nil
}

//四舍五入
func Round(f float64) float64 {
	return math.Floor(f + 0.5)
}

type HeapContainer struct {
	c    []interface{}
	size int
	less func(x, y interface{}) bool
}

func NewHeapC(c []interface{}, less func(x, y interface{}) bool) *HeapContainer {
	p := &HeapContainer{
		c:    c,
		size: len(c),
		less: less,
	}
	heap.Init(p)
	return p
}

func (p *HeapContainer) Len() int {
	return p.size
}

func (p *HeapContainer) Less(i, j int) bool {
	return p.less(p.c[i], p.c[j])
}

func (p *HeapContainer) Swap(i, j int) {
	/*这算法实现的还会有-1出现（即当数组元素为0时）*/
	if i < 0 || j < 0 {
		return
	}
	tmp := p.c[i]
	p.c[i] = p.c[j]
	p.c[j] = tmp
}

func (p *HeapContainer) Push(x interface{}) {
	if p.size < len(p.c) {
		p.c[p.size] = x
	} else {
		p.c = append(p.c, x)
	}
	p.size++
}

func (p *HeapContainer) Pop() interface{} {
	if p.size == 0 {
		return nil
	}
	p.size--
	t := p.c[p.size]
	return t
}

func Int2Inter(nums []int) []interface{} {
	r := make([]interface{}, len(nums))
	for i := 0; i < len(nums); i++ {
		r[i] = nums[i]
	}
	return r
}

type Heap struct {
	heapc *HeapContainer
}

func NewHeap(c []interface{}, less func(x, y interface{}) bool) *Heap {
	return &Heap{
		heapc: NewHeapC(c, less),
	}
}

func (p *Heap) Len() int {
	return p.heapc.Len()
}

func (p *Heap) Push(x interface{}) {
	heap.Push(p.heapc, x)
}

func (p *Heap) Pop() interface{} {
	return heap.Pop(p.heapc)
}

func NewHeapInt(nums []int, less func(x interface{}, y interface{}) bool) *Heap {
	return NewHeap(Int2Inter(nums), less)
}

func CreateListInt(nums []int) *ListNode{
	if len(nums) == 0 {
		return nil
	}
	var head, iter *ListNode
	for _, v := range nums {
		if iter == nil {
			iter = &ListNode{Val: v, Next: nil}
			head = iter
		} else {
			iter.Next = &ListNode{
				Val:  v,
				Next: nil,
			}
			iter = iter.Next
		}
	}
	return head
}

func MaxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func IsDigit(c byte) bool {
	if c >= 48 && c <= 57 {
		return true
	}
	return false
}

func Ctoi(b byte) int {
	return int(b - 48)
}

func Itoc(i int) byte {
	return byte(i + 48)
}

func IsSig(c byte) bool {
	return c == 45 || c == 43
}




type CircularQueue struct {
	c []interface{}
	head, tail int
}


func NewCQueue(k int) Queue {
	return &CircularQueue{
		c:    make([]interface{}, k+1),
		head: 0,
		tail: 0,
	}
}


func (q *CircularQueue) Push(e interface{}) {
	if q.Full() {
		size := len(q.c)
		newsize := 2*size
		if size >= 1024 {
			newsize = size + size / 2
		}
		newc := make([]interface{}, newsize)
		n := copy(newc, q.c[q.head:])
		copy(newc[n:], q.c[:q.head])
		q.c = newc
		q.head = 0
		q.tail = size
	}
	q.tail = (q.tail + 1) % len(q.c)
	q.c[q.tail] = e
}


func (q *CircularQueue) Pop() interface{} {
	if q.Empty() {
		panic("The queue is empty")
	}
	i := (q.head+1)%len(q.c)
	q.head = i
	return q.c[i]
}


func (q *CircularQueue) Front() interface{} {
	if q.Empty() {
		panic("The queue is empty")
	}
	return q.c[(q.head + 1) % len(q.c)]
}


func (q *CircularQueue) Rear() interface{} {
	if q.Empty() {
		panic("The queue is empty")
	}
	return q.c[(q.tail) % len(q.c)]
}


func (q *CircularQueue) Empty() bool {
	return q.head == q.tail
}


func (q *CircularQueue) Full() bool {
	return (q.tail + 1) % len(q.c) == q.head
}

func Swap(nums []int, x, y int) {
	z := nums[x]
	nums[x] = nums[y]
	nums[y] = z
}


