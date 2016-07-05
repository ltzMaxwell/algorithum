package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	fmt.Println(q)
	q.Push("a")
	fmt.Println(q)
	fmt.Println("len is :", q.Len())
	fmt.Println(q.Peek())
	fmt.Println(q.shift())
	fmt.Println(q)
	fmt.Println("len is :", q.Len())
}
