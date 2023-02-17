package main

import "fmt"

// 这个文件是二叉树最常规的算法，递归算法

func Recurrence(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Data)
	if node.LeftChild != nil {
		Recurrence(node.LeftChild)
	}
	fmt.Printf("%d ", node.Data)
	if node.RightChild != nil {
		Recurrence(node.RightChild)
	}
	fmt.Printf("%d ", node.Data)
}

func FirstRecurrence(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Data)
	if node.LeftChild != nil {
		FirstRecurrence(node.LeftChild)
	}
	if node.RightChild != nil {
		FirstRecurrence(node.RightChild)
	}
}

func MiddleRecurrence(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	if node.LeftChild != nil {
		MiddleRecurrence(node.LeftChild)
	}
	fmt.Printf("%d ", node.Data)
	if node.RightChild != nil {
		MiddleRecurrence(node.RightChild)
	}
}
func LastRecurrence(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	if node.LeftChild != nil {
		LastRecurrence(node.LeftChild)
	}
	if node.RightChild != nil {
		LastRecurrence(node.RightChild)
	}
	fmt.Printf("%d ", node.Data)
}
