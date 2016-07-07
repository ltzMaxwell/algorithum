package main

import (
	"basic/algorithum/linkedlist"
	"errors"
	"fmt"
	"math"
)

type Item struct {
	Value interface{}
	Key   string
}

type HashTable struct {
	Table map[int]*linkedlist.List
	Size  int
	Cap   int
}

func NewHashTable(cap int) *HashTable {
	return &HashTable{
		Table: make(map[int]*linkedlist.List, cap),
		Size:  0,
		Cap:   cap,
	}
}

func (h *HashTable) Put(item *Item) {
	fmt.Println("put")
	index := h.position(item.Key)
	fmt.Println("key ,index is ", item.Key, index)

	//find linkelist with index
	l, ok := h.Table[index]
	if !ok {
		fmt.Println("not find list")
		l = linkedlist.NewList()
		l.Append(&linkedlist.Node{Value: item})
		h.Table[index] = l
		return
	}
	fmt.Println("list exist, ", l)
	node, err := h.find(index, item.Key)
	//item not found in this list
	if err != nil {
		l.Append(&linkedlist.Node{Value: item})
	} else {
		node.Value = item
	}
}

func (h *HashTable) Get(key string) (interface{}, error) {
	fmt.Println("get key :", key)
	//get index
	index := h.position(key)
	fmt.Println("index is ", index)
	//find result using index(key of map) and key(key of node in linkedlist)

	node, err := h.find(index, key)
	if err != nil {
		fmt.Println("get nothing")
		return nil, errors.New("get nothing")
	}
	return node.Value.(*Item).Value, nil
}

//return type node in linkedlist
func (h *HashTable) find(index int, key string) (*linkedlist.Node, error) {
	//find linked list in the map
	l, ok := h.Table[index]
	fmt.Println("l is ", l)
	//find target in linkelist
	var result *linkedlist.Node
	if ok {
		//traverse linkedlist
		l.Each(func(node *linkedlist.Node) {
			fmt.Println("each result is : ", node)
			if node.Value.(*Item).Key == key {
				result = node
			}
		})
		//else not found
		if result == nil {
			return nil, errors.New("not found")
		} else {
			return result, nil
		}
	} else {
		return nil, errors.New("list is not found")
	}
}

func (h *HashTable) position(key string) int {
	code := hashCode(key)
	return code % h.Cap
}

// Horner's Method to hash string of length L (O(L))
func hashCode(s string) int {
	hash := int32(0)
	for i := 0; i < len(s); i++ {
		hash = int32(hash<<5-hash) + int32(s[i])
		hash &= hash
	}
	return int(math.Abs(float64(hash)))
}

func main() {
	ht := NewHashTable(10)
	ht.Put(&Item{Value: "tom", Key: "name"})
	r, _ := ht.Get("name")
	fmt.Println("result is :", r)
	r, _ = ht.Get("age")
	fmt.Println("result is :", r)
}
