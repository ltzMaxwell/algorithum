package tree

import (
	"errors"
	"math"
)

type BTreeNode struct {
	data   interface{}
	parent *BTreeNode
	LChild *BTreeNode
	RChild *BTreeNode
	Height float64
	Size   float64
}

//New node
func NewNode(el interface{}) *BTreeNode {
	return &BTreeNode{
		data: el,
	}
}

//--------------------
func (b *BTreeNode) IsLeaf() bool {
	if !b.hasLeftChild() && !b.hasRightChild() {
		return true
	} else {
		return false
	}
}

func (b *BTreeNode) IsLChild() bool {
	if b.HasParent() && b.parent.LChild == b {
		return true
	} else {
		return false
	}
}

func (b *BTreeNode) IsRChild() bool {
	if b.HasParent() && b.parent.RChild == b {
		return true
	} else {
		return false
	}
}

//--------------------
//Get data
func (b *BTreeNode) GetData() (interface{}, error) {
	if b != nil {
		return b.data, nil
	} else {
		return nil, errors.New("node is nil")
	}
}

func (b *BTreeNode) SetData(el interface{}) error {
	if b != nil {
		b.data = el
		return nil
	} else {
		return errors.New("node is nil")
	}
}

//--------------------
func (b *BTreeNode) GetHeight() float64 {
	return b.Height
}

//set height of subtree belog to this node
//so height is 1+ max(left height , right height)
func (b *BTreeNode) SetHeight() {
	//init height
	var h float64
	h = 0
	//find max of left height and right height
	if b.hasLeftChild() {
		h = math.Max(0, (float64)(b.LChild.GetHeight()))
	} else {
		h = math.Max(0, (float64)(b.RChild.GetHeight()))
	}
	//plus b itself
	h++
	b.Height = h

	//set parent height
	if b.HasParent() {
		b.parent.SetHeight()
	}
}

//--------------------
func (b *BTreeNode) hasLeftChild() bool {
	if b.LChild != nil {
		return true
	} else {
		return false
	}
}

func (b *BTreeNode) setLChild(c *BTreeNode) {
	if b == nil {
		return
	}
	//set relation
	b.LChild = c
	c.parent = b
	//set height
	b.SetHeight()
}

//--------------------
func (b *BTreeNode) hasRightChild() bool {
	if b.RChild != nil {
		return true
	} else {
		return false
	}
}

func (b *BTreeNode) setRChild(c *BTreeNode) {
	if b == nil {
		return
	}
	//set relation
	b.RChild = c
	c.parent = b
	//set height
	b.SetHeight()
}

//--------------------

func (b *BTreeNode) SetParent(p *BTreeNode) {
	//set parent
	b.parent = p
	//update parent height
	b.parent.SetHeight()
}

func (b *BTreeNode) GetParent() *BTreeNode {
	if b.HasParent() {
		return b.parent
	} else {
		return nil
	}
}
func (b *BTreeNode) HasParent() bool {
	if b.parent != nil {
		return true
	} else {
		return false
	}
}
