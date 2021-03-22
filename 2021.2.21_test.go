package leetcode_go

import (
	"fmt"
	"testing"
)

func TestCircularQueue(t *testing.T) {
	k := 5
	q := NewCQueue(k)
	for i := 0; i < k; i++ {
		q.Push(i)
	}
	q.Push(10)
	q.Pop()
	q.Pop()
	q.Push(6)
	fmt.Println(q)

}