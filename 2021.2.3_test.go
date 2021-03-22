package leetcode_go

import (
	"fmt"
	"testing"
)

func TestLruCache(t *testing.T) {
	c := LRUConstructor(3)
	c.Put(1,1)
	c.Put(2,2)
	c.Put(3,3)
	c.Put(4,4)
	fmt.Println(c.Get(4))
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(2))
	fmt.Println(c.Get(1))
	c.Put(5,5)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(2))
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(4))
	fmt.Println(c.Get(5))
}
