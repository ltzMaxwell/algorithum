package priority_queue

import (
	"fmt"
	"testing"
)

func TestPq(t *testing.T) {
	pq := NewPQ()
	fmt.Println(pq.Len())

	pq.Insert(Item{6, "e"})
	pq.Insert(Item{1, "b"})
	pq.Insert(Item{2, "a"})

	pq.ChangePriority(Item{2, "a"}, 8)

	fmt.Println(pq.Extract())
	fmt.Println(pq.Extract())
	fmt.Println(pq.Extract())
	fmt.Println(pq.Extract())
}
