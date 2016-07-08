package tree

import (
	"fmt"
	"testing"
)

func TestBTree(t *testing.T) {
	//new tree
	tree := NewBTree()

	//new root node
	na := NewNode("a")
	fmt.Println(na)

	fmt.Println("is leaf ?", na.IsLeaf())
	fmt.Println("is left child ?", na.IsLChild())
	fmt.Println("is right child  ?", na.IsRChild())

	//add root
	tree.root = na
	//get root
	fmt.Println("root is ", tree.GetRoot())

	//get data
	val, _ := na.GetData()
	fmt.Println("get data ", val)

	//set data
	// na.SetData("b")
	// bval, _ := na.GetData()
	// fmt.Println("set  data")
	// fmt.Println("get data after set data", bval)

	//get height
	fmt.Println("get height")
	fmt.Println(na.GetHeight())

	//new left child node
	nb := NewNode("b")
	//set as left child
	na.setLChild(nb)
	fmt.Println("--------------------------------------")
	fmt.Println("height of na is :", na.GetHeight())
	fmt.Println("height of nb is :", nb.GetHeight())
	fmt.Println("--------------------------------------")

	//new left child node
	nd := NewNode("d")
	//set as left child
	nb.setLChild(nd)

	fmt.Println("--------------------------------------")
	fmt.Println("height of na is :", na.GetHeight())
	fmt.Println("height of nb is :", nb.GetHeight())
	fmt.Println("height of nd is :", nd.GetHeight())
	fmt.Println("--------------------------------------")

	//new left child node
	nh := NewNode("h")
	//set as left child
	nd.setLChild(nh)

	fmt.Println("--------------------------------------")
	fmt.Println("height of na is :", na.GetHeight())
	fmt.Println("height of nb is :", nb.GetHeight())
	fmt.Println("height of nd is :", nd.GetHeight())
	fmt.Println("height of nh is :", nh.GetHeight())
	fmt.Println("--------------------------------------")

	//new right child node
	nc := NewNode("c")
	na.setRChild(nc)

	fmt.Println("--------------------------------------")
	fmt.Println("height of na is :", na.GetHeight())
	fmt.Println("height of nb is :", nb.GetHeight())
	fmt.Println("height of nc is :", nc.GetHeight())
	fmt.Println("height of nd is :", nd.GetHeight())
	fmt.Println("height of nh is :", nh.GetHeight())

	fmt.Println("--------------------------------------")

	ne := NewNode("e")
	nc.setLChild(ne)

	fmt.Println("--------------------------------------")
	fmt.Println("height of na is :", na.GetHeight())
	fmt.Println("height of nb is :", nb.GetHeight())
	fmt.Println("height of nc is :", nc.GetHeight())
	fmt.Println("height of nd is :", nd.GetHeight())
	fmt.Println("height of ne is :", ne.GetHeight())
	fmt.Println("height of nh is :", nh.GetHeight())

	fmt.Println("--------------------------------------")

	nf := NewNode("f")
	ne.setLChild(nf)

	fmt.Println("--------------------------------------")
	fmt.Println("height of na is :", na.GetHeight())
	fmt.Println("height of nb is :", nb.GetHeight())
	fmt.Println("height of nc is :", nc.GetHeight())
	fmt.Println("height of nd is :", nd.GetHeight())
	fmt.Println("height of ne is :", ne.GetHeight())
	fmt.Println("height of nf is :", nf.GetHeight())
	fmt.Println("height of nh is :", nh.GetHeight())

	fmt.Println("--------------------------------------")

	//-------------------------------
	//traverse

	fmt.Println("pre order")
	//preOrder
	l := tree.PreOrder(na)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*BTreeNode).data.(string))
	}

	fmt.Println("--------------------------------------")

	fmt.Println("in order")
	//inOrder
	l = tree.InOrder(na)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*BTreeNode).data.(string))
	}

	fmt.Println("--------------------------------------")
	fmt.Println("post order")
	//postOrder
	l = tree.PostOrder(na)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*BTreeNode).data.(string))
	}

	fmt.Println("tree height is ", tree.GetHeight())

}
