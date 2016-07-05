package priority_queue

import (
	"basic/algorithum/heap"
	"basic/algorithum/queue"
	"fmt"
)

type Item struct {
	Priority int
	value    interface{}
}

func (i Item) Less(than heap.Sortable) bool {
	return i.Priority < than.(Item).Priority
}

type PQ struct {
	pq *heap.Heap
}

func NewPQ() *PQ {
	return &PQ{
		pq: heap.NewHeap(),
	}
}

func (p *PQ) Len() int {
	return p.pq.Len()
}

func (p *PQ) Insert(el Item) {
	p.pq.Insert(heap.Sortable(el))
}

func (p *PQ) Extract() (el Item) {
	raw := p.pq.Extract()
	if raw != nil {
		el = raw.(Item)
	}
	return
}

func (p *PQ) ChangePriority(el Item, prior int) {
	fmt.Println("priority ", el)
	//cache, cache els until given el found
	q := queue.NewQueue()
	//pop from pq
	pop := p.Extract()
	//find el
	for el != pop {
		if p.Len() == 0 {
			panic("item not found")
		}
		fmt.Println("push pop", pop)
		q.Push(pop)
		pop = p.Extract()
		fmt.Println("pop ...", pop)
	}
	fmt.Println("found el ,q is :", q)
	//el found
	el.Priority = prior
	fmt.Println("el to push is :", el)
	fmt.Println("q len is ", q.Len())
	//insert this value
	p.Insert(el)
	//q.len is changing , so not right for this way
	for q.Len() > 0 {
		fmt.Println("insert cache")
		p.Insert(q.Shift().(Item))
	}
}
