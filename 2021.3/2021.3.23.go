package _2021_3

import "leetcode_go"

/**
  341. 扁平化嵌套列表迭代器 中等
  将整个嵌套列表视作树，使用栈，构造中序迭代器
*/

type NestedInteger struct{}

func (this NestedInteger) IsInteger() bool { return true }

func (this NestedInteger) GetInteger() int { return 0 }

func (n *NestedInteger) SetInteger(value int) {}

func (this *NestedInteger) Add(elem NestedInteger) {}

func (this NestedInteger) GetList() []*NestedInteger { return nil }

func isLeaf(this *NestedInteger) bool {
	if this.IsInteger() || len(this.GetList()) == 0 {
		return true
	} else {
		return false
	}
}

type NestedIntWithIndex struct {
	ni   *NestedInteger
	nidx int //next index
}

func NewNestedIntWithIndex(ni *NestedInteger, nidx int) *NestedIntWithIndex {
	return &NestedIntWithIndex{
		ni:   ni,
		nidx: nidx,
	}
}

type NestedIterator struct {
	iter *NestedIntWithIndex
	s    *leetcode_go.Stack
	next *NestedInteger
	new  bool
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	var root *NestedInteger
	if len(nestedList) != 0 {
		root = new(NestedInteger)
		for _, v := range nestedList {
			root.Add(*v)
		}
	}

	nest := &NestedIterator{
		iter: NewNestedIntWithIndex(root, 0),
		s:    leetcode_go.NewStack(1),
	}
	return nest
}

func (this *NestedIterator) Next() int {
	if this.HasNext() {
		this.new = false
		return this.next.GetInteger()
	}
	panic("there is no next integer.")
}

func (this *NestedIterator) HasNext() bool {
	if this.new {
		return true
	} else {
		this.findNext()
		return this.new
	}
}

func (this *NestedIterator) findNext() {
	//已有新next
	if this.next != nil && this.new {
		return
	}
	//无了
	if !this.new && this.iter == nil && this.s.Empty() {
		return
	}

	for this.iter != nil || !this.s.Empty() {
		if this.iter != nil {
			//第一次遍历到该节点
			if !isLeaf(this.iter.ni) && this.iter.nidx < len(this.iter.ni.GetList()) {
				this.s.Push(NewNestedIntWithIndex(this.iter.ni, this.iter.nidx+1))
				this.iter = NewNestedIntWithIndex(this.iter.ni.GetList()[this.iter.nidx], 0)
			} else {
				tmp := this.iter
				this.iter = nil
				if tmp.ni.IsInteger() {
					this.next = tmp.ni
					this.new = true
					return
				}
			}

		} else {
			this.iter = this.s.Pop().(*NestedIntWithIndex)
			//遍历，检查是否是leaf
			if !isLeaf(this.iter.ni) && this.iter.nidx < len(this.iter.ni.GetList()) {
				//不是integer，且list len > 0
				//二叉树在这里直接进入右子树，且由于没有其他子树可遍历，则无需入栈
				this.s.Push(NewNestedIntWithIndex(this.iter.ni, this.iter.nidx+1))
				this.iter = NewNestedIntWithIndex(this.iter.ni.GetList()[this.iter.nidx], 0)

			} else {
				tmp := this.iter
				this.iter = nil
				if tmp.ni.IsInteger() {
					this.next = tmp.ni
					this.new = true
					return
				}
			}
		}
	}
}
