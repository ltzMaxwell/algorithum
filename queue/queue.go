package queue

import (
	"sync"
)

type Queue struct {
	lock sync.Mutex
	data []interface{}
	len  int
}

func NewQueue() *Queue {
	queue := &Queue{}
	queue.data = make([]interface{}, 0)
	return queue
}

func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.len
}

func (q *Queue) isEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.len == 0
}

func (q *Queue) Shift() (el interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	el, q.data = q.data[0], q.data[1:]
	q.len--
	return el
}

func (q *Queue) Push(el interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.data = append(q.data, el)
	q.len++
}

func (q *Queue) Peek() (el interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	el = q.data[0]
	return el
}
