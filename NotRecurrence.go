package main

import (
	"binTree/mystack"
	"fmt"
)

// 先序遍历
func FirstNotRecurrence(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	stack := mystack.NewStack()
	stack.Push(node)
	for !stack.IsEmpty() {
		cur := stack.Pop().(*BinaryTreeNode)
		fmt.Printf("%d ", cur.Data)
		if cur.RightChild != nil {
			stack.Push(cur.RightChild)
		}
		if cur.LeftChild != nil {
			stack.Push(cur.LeftChild)
		}
	}

}

func MiddleNotRecurrence(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	stack := mystack.NewStack()
	cur := node
	for !stack.IsEmpty() || cur != nil {
		if cur != nil {
			stack.Push(cur)
			cur = cur.LeftChild
		} else {
			cur = stack.Pop().(*BinaryTreeNode)
			fmt.Printf("%d ", cur.Data)
			cur = cur.RightChild
		}
	}
}

func LastNotRecurrence(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	stack1 := mystack.NewStack()
	stack2 := mystack.NewStack()
	stack1.Push(node)
	for !stack1.IsEmpty() {
		cur := stack1.Pop().(*BinaryTreeNode)
		if cur.LeftChild != nil {
			stack1.Push(cur.LeftChild)
		}
		if cur.RightChild != nil {
			stack1.Push(cur.RightChild)
		}
		stack2.Push(cur)
	}
	for !stack2.IsEmpty() {
		fmt.Printf("%d ", stack2.Pop().(*BinaryTreeNode).Data)
	}
}
