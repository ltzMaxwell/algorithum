package linkedlist

// import (
// 	"fmt"
// )

//Node ...
type Node struct {
	Value interface{}
	Pre   *Node
	Next  *Node
}

//NewNode ...
func NewNode() *Node {
	return &Node{
		Value: nil,
	}
}

//List ...
type List struct {
	len  int
	Head *Node
	Tail *Node
}

//NewList ...
func NewList() *List {
	return &List{
		len: 0,
	}
}

//Len ...
func (l *List) Len() int {
	return l.len
}

//IsEmpty ...
func (l *List) IsEmpty() bool {
	return l.len == 0
}

//Prepend ...
func (l *List) Prepend(el *Node) {
	//if a empty list
	if l.IsEmpty() {
		l.Head = el
		l.Tail = el
		l.len = 1
		return
	}
	//get current head
	currentHead := l.Head

	//combine head and older head
	currentHead.Pre = el
	el.Next = currentHead

	//set list head
	l.Head = el
	l.len++
	return
}

func (l *List) Append(el *Node) {
	if l.IsEmpty() {
		l.Head = el
		l.Tail = el
		l.len = 1
		return
	}
	currentTail := l.Tail
	currentTail.Next = el
	el.Pre = currentTail
	//set new tail
	l.Tail = el
	l.len++
	return
}

//Add element into list
func (l *List) Add(index int, el *Node) {
	//condition
	if (index < 0) || (index > l.len-1) {
		panic("out of range")
	}
	//head
	if index == 0 {
		l.Prepend(el)
		return
	}
	//add to tail - 1 , tail wouldn't change
	if index == (l.len - 1) {
		currentTail := l.Tail
		preTail := l.Tail.Pre

		//combine
		el.Pre = preTail
		preTail.Next = el

		el.Next = currentTail
		currentTail.Pre = el
		l.len++

		return
	}

	p := l.Get(index - 1)
	//n is the one in current index
	n := l.Get(index)
	//combine
	el.Pre = p
	p.Next = el

	el.Next = n
	n.Pre = el
	//len refresh
	l.len++

	return
}

func (l *List) Remove(index int) (cursor int) {
	//is empty
	if l.IsEmpty() {
		panic("list is empty")
	}
	//range restraint
	if (index < 0) || (index > (l.len - 1)) {
		panic("out of range")
	}
	//is head
	if index == 0 {
		l.Head = l.Head.Next
		l.len--
		return 0
	}
	//tail
	if index == (l.len - 1) {
		l.Tail = l.Tail.Pre
		l.len--
		return index
	}

	//in range
	p := l.Get(index - 1)
	n := l.Get(index + 1)
	//combine
	p.Next = n
	n.Pre = p
	l.len--

	return index
}

func (l *List) Each(f func(node *Node)) {
	for node := l.Head; node != nil; node = node.Next {
		f(node)
	}
}

//Get ...
func (l *List) Get(index int) (el *Node) {
	if l.len == 0 {
		if index == 0 {
			return l.Head
		}
	} else {
		//make sure index is in range
		if index <= (l.len - 1) {
			cursor := index
			el = l.Head
			for cursor > 0 {
				el = el.Next
				cursor--
			}
			return
		}
	}
	return
}
