package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {

	var arr = []Int{10, 40, 30, 60, 90, 70, 20, 50, 80}

	h := NewHeap()
	fmt.Println(h)

	for i := 0; i < len(arr); i++ {
		h.Insert(arr[i])
	}
	fmt.Println(h)

	fmt.Println("extract result :", h.Extract())
	fmt.Println("extract result :", h.Extract())

	//insert 85
	add := Int(85)
	h.Insert(add)
	fmt.Println(h)
	fmt.Println("extract result :", h.Extract())
	fmt.Println("extract result :", h.Extract())
	fmt.Println("extract result :", h.Extract())
	fmt.Println("extract result :", h.Extract())
	fmt.Println("extract result :", h.Extract())
	fmt.Println("extract result :", h.Extract())
	fmt.Println("extract result :", h.Extract())
	fmt.Println("extract result :", h.Extract())
	fmt.Println("extract result :", h.Extract())
	fmt.Println("extract result :", h.Extract())
	fmt.Println("extract result :", h.Extract())

	//remove 90
	h.Remove(arr[4])
	fmt.Println(h)
}
