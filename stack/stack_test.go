package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
	fmt.Println(s)
	fmt.Println("is empty?", s.IsEmpty())

	s.Push("a")

	fmt.Println(s)
	fmt.Println(s.Peek())
	fmt.Println("is empty?", s.IsEmpty())
	fmt.Println("length is :", s.Len())
	fmt.Println(s.Pop())
	fmt.Println("length is :", s.Len())
	fmt.Println(s)
}
