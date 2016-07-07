package linkedlist

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	fmt.Println("test linked list")
	l := NewList()

	l.Prepend(&Node{Value: "a"})
	l.Prepend(&Node{Value: "b"})
	l.Prepend(&Node{Value: "c"})
	l.Append(&Node{Value: "d"})
	l.Append(&Node{Value: "e"})

	//add
	l.Add(0, &Node{Value: "f"})
	l.Add(5, &Node{Value: "g"})
	l.Add(1, &Node{Value: "h"})

	l.Add(6, &Node{Value: "i"})
	l.Add(8, &Node{Value: "j"})

	fmt.Println("len of list is :", l.len)

	fmt.Println("get 0 is ", l.Get(0))
	fmt.Println("get 1 is ", l.Get(1))
	fmt.Println("get 2 is ", l.Get(2))
	fmt.Println("get 3 is ", l.Get(3))
	fmt.Println("get 4 is ", l.Get(4))
	fmt.Println("get 5 is ", l.Get(5))
	fmt.Println("get 6 is ", l.Get(6))
	fmt.Println("get 7 is ", l.Get(7))
	fmt.Println("get 8 is ", l.Get(8))
	fmt.Println("get 9 is ", l.Get(9))

	// fmt.Printf("%d removed \n", l.Remove(0))
	// fmt.Printf("%d removed \n", l.Remove(8))
	fmt.Printf("%d removed \n", l.Remove(1))
	fmt.Printf("%d removed \n", l.Remove(1))

	fmt.Println("get 0 is ", l.Get(0))
	fmt.Println("get 1 is ", l.Get(1))
	fmt.Println("get 2 is ", l.Get(2))
	fmt.Println("get 3 is ", l.Get(3))
	fmt.Println("get 4 is ", l.Get(4))
	fmt.Println("get 5 is ", l.Get(5))
	fmt.Println("get 6 is ", l.Get(6))
	fmt.Println("get 7 is ", l.Get(7))
	fmt.Println("get 8 is ", l.Get(8))
	fmt.Println("get 9 is ", l.Get(9))

	fmt.Println("head is ", l.Head)
	fmt.Println("tail is ", l.Tail)
	fmt.Println("--------------------")

	l.Each(func(node *Node) {
		fmt.Println(node)
	})
}
