package tree

import (
	"container/list"
)

type BTree struct {
	root   *BTreeNode
	Height float64
	Size   int
}

func NewBTree() *BTree {

	return &BTree{
		root:   nil,
		Height: 0,
		Size:   0,
	}
}

func (b *BTree) GetSize() int {
	return b.Size
}

//get height of root
func (b *BTree) GetHeight() float64 {
	return b.root.GetHeight()
}

//empty if root is nil
func (b *BTree) IsEmpty() bool {
	return b.root == nil
}

func (b *BTree) GetRoot() *BTreeNode {
	return b.root
}

//traverse
//-------------------------------------
func postOrder(root *BTreeNode, traverse *list.List) {
	if root.hasLeftChild() {
		postOrder(root.LChild, traverse)
	}
	if root.hasRightChild() {
		postOrder(root.RChild, traverse)
	}
	traverse.PushBack(root)
}

func (b *BTree) PostOrder(root *BTreeNode) *list.List {
	traverse := list.New()
	postOrder(root, traverse)
	return traverse
}

func inOrder(root *BTreeNode, traverse *list.List) {
	//left subtree first in inOrder
	if root.hasLeftChild() {
		inOrder(root.LChild, traverse)
	}
	//record
	traverse.PushBack(root)
	//right sub tree
	if root.hasRightChild() {
		inOrder(root.RChild, traverse)
	}
}

func (b *BTree) InOrder(root *BTreeNode) *list.List {
	traverse := list.New()
	inOrder(root, traverse)
	return traverse
}

func preOrder(root *BTreeNode, traverse *list.List) {
	//push root first in preOrder
	traverse.PushBack(root)
	if root.hasLeftChild() {
		preOrder(root.LChild, traverse)
	}
	if root.hasRightChild() {
		preOrder(root.RChild, traverse)
	}
}

func (b *BTree) PreOrder(root *BTreeNode) *list.List {
	traverse := list.New()
	preOrder(root, traverse)
	return traverse
}
