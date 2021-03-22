package leetcode_go

//the key point is how to find the LRU item to delete it in a variable param
/**
 * 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制 。
 * 实现 LRUCache 类：
 * LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
 * int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
 * void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
 * 解法：使用map存储k-item，item包括value及lastUseNode，即使用链表维护LRU关系。对该链表，使用头插法先将节点插入表头
 * 然后再将该结点移至表尾（lunr），lunr表示最新节点，lunh.next表示LRU节点。删去时删去lunh.next
 * 由于维护了尾指针，所以插入比普通头插复杂
 */

//将value整合到lun中，不再用item，map中存储node
type LastUseNode struct {
	key  int
	value int
	next *LastUseNode
	last *LastUseNode
}


//type LRUCacheItem struct {
//	value int
//	lun   *LastUseNode
//}

type LRUCache struct {
	m    map[int]*LastUseNode
	lunh *LastUseNode //with head node
	lunr *LastUseNode //rear of the link
	cap  int
	size int
}

func LRUConstructor(capacity int) LRUCache {
	c := LRUCache{
		m:    make(map[int]*LastUseNode, capacity),
		lunh: &LastUseNode{},
		lunr: nil,
		cap:  capacity,
		size: 0,
	}
	return c
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		this.moveToRear(v)
		return v.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.m[key]; ok {
		this.moveToRear(v)
		//update
		v.value = value
	} else {
		//check if the cache is full
		if this.size < this.cap {
			item := &LastUseNode{
				value: value,
				key:  key,
				next: this.lunh.next,
				last: this.lunh,
			}
			//第一个节点时，lunh.next为nil，不能访问last
			if this.lunh.next != nil {
				this.lunh.next.last = item
			}
			//头插
			this.lunh.next = item
			//第一个节点时，设置该节点为尾节点
			if this.lunr == nil {
				this.lunr = item
			}

			this.m[key] = item
			this.size++
			//为了实现上简单些，使用moveToRear，而不是直接插到表尾
			this.moveToRear(item)
		} else {
			//删除表头节点，实现上不是删除，而是修改key值（重复利用节点）
			hn := this.lunh.next
			//map的key值不能修改，只能先删除再插入
			delete(this.m, hn.key)
			hn.key = key
			hn.value = value
			this.m[key] = hn
			this.moveToRear(hn)
		}
	}
}

func (this *LRUCache) moveToRear(lun *LastUseNode) {
	//insert this lun to the fisrt node of lru link
	//if lun is rear
	if this.size == 1 || lun == this.lunr{
		return
	}
	lun.last.next = lun.next
	lun.next.last = lun.last
	lun.last = this.lunr
	lun.next = nil
	this.lunr.next = lun
	this.lunr = lun
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
