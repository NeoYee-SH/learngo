package main

import (
	"fmt"
	"github.com/yihuaiyuan/learngo/tree"
)

//组合扩展node包
//type myTreeNode struct {
//	node *tree.Node //放指针，避免需要拷贝
//}

//func (myNode *myTreeNode) myTraverse() {
//
//	if myNode == nil || myNode.node == nil {
//		return
//	}
//
//	left := myTreeNode{myNode.node.Left}
//	right := myTreeNode{myNode.node.Right}
//
//	left.myTraverse()
//	right.myTraverse()
//	myNode.node.Print()
//}


func main() {

	root := tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left = new(tree.Node)
	root.Right.Left.SetValue(4)

	nodeCount := 0
	root.TraverseFunc( func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)

	c := root.TraverseWithChannel()

	maxValue := 0
	for node := range c {
		if node.Value > maxValue {
			maxValue = node.Value
		}
	}
	fmt.Println("Max value is :", maxValue)


	//myRoot := myTreeNode{&root}
	//myRoot.myTraverse()
}
