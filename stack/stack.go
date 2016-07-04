package stack

import (
	"sync"
)

type Stack struct {
	lock  sync.Mutex
	Stack []interface{}
	len   uint64
}

func NewStack() *Stack {
	s := &Stack{}
	s.Stack = make([]interface{}, 0)
	s.len = 0
	return s
}

func (s *Stack) Len() uint64 {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.len
}

func (s *Stack) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.len == 0
}

func (s *Stack) Push(el interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	prepend := make([]interface{}, 1)
	prepend[0] = el
	s.Stack = append(prepend, s.Stack...)
	s.len++
}

func (s *Stack) Pop() (el interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	el, s.Stack = s.Stack[0], s.Stack[1:]
	s.len--
	return el
}

func (s *Stack) Peek() (el interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	el = s.Stack[0]
	return el
}
