package main

import (
	"fmt"
)

type TreeNode struct {
	Data string
	Left *TreeNode
	Right *TreeNode
}

func PreOrder (tree *TreeNode) {
	if tree == nil {
		return
	}

	//打印根结点
	fmt.Println(tree.Data)
	PreOrder(tree.Left)
	PreOrder(tree.Right)
}

func MidOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	MidOrder(tree.Left)
	//打印根结点
	fmt.Println(tree.Data)
	MidOrder(tree.Right)
}

func PostOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	PostOrder(tree.Left)
	PostOrder(tree.Right)
	//打印根结点
	fmt.Println(tree.Data)
}

func Depth(tree *TreeNode) (i int){
	if tree == nil {
		return -1
	}
	m := Depth(tree.Left)
	n := Depth(tree.Right)
	if m>n {
		i = m+1
	}
	if n>m {
		i = n+1
	}
	return i
}

func ExChange(root *TreeNode) {
	root.Left, root.Right = root.Right, root.Left
}
func MirrorTree(root *TreeNode){
	ExChange(root)
	ExChange(root.Left)
	ExChange(root.Right)
	PreOrder(root)
}
func main () {
	t := &TreeNode{Data: "A"}
	fmt.Println(t)
	t.Left = &TreeNode{Data: "B"}
	t.Right = &TreeNode{Data: "C"}
	t.Left.Left = &TreeNode{Data: "D"}
	t.Left.Right = &TreeNode{Data: "E"}
	fmt.Println("先序遍历：")
	PreOrder(t)
	fmt.Println("中序遍历：")
	MidOrder(t)
	fmt.Println("后序遍历：")
	PostOrder(t)
	fmt.Printf("树的深度为：%d",Depth(t))
	fmt.Println("镜像：")
	MirrorTree(t)

}