package main

import "fmt"

//Morris遍历：利用空闲节点往回指，避免使用了额外的空间
//  1、让cur节点指向root节点，根据如下标准进行移动
//  2、cur为空节点 return
//  3、cur不为空但是没有左孩子，cur右移
//  4、有左孩子，找到左树上最右的节点 mostRight
//  5、mostRight的右孩子为空 右孩子指向cur cur左移
//  6、mostRight的右孩子指向cur，说明第二次来到 right置为nil cur 右移

func Morris(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	//cur指向头节点
	cur := node
	var mostRight *BinaryTreeNode = nil
	for cur != nil {
		//cur没有左孩子的话
		fmt.Printf("%d ", cur.Data)
		mostRight = cur.LeftChild
		if mostRight != nil {
			for mostRight.RightChild != nil && mostRight.RightChild != cur {
				mostRight = mostRight.RightChild
			}
			if mostRight.RightChild == nil {
				mostRight.RightChild = cur
				cur = cur.LeftChild
				continue
			} else {
				mostRight.RightChild = nil
			}
		}
		cur = cur.RightChild
	}
}
func FirstMorris(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	//cur指向头节点
	cur := node
	var mostRight *BinaryTreeNode = nil
	for cur != nil {
		//cur没有左孩子的话
		mostRight = cur.LeftChild
		if mostRight != nil {
			for mostRight.RightChild != nil && mostRight.RightChild != cur {
				mostRight = mostRight.RightChild
			}
			if mostRight.RightChild == nil {
				fmt.Printf("%d ", cur.Data)
				mostRight.RightChild = cur
				cur = cur.LeftChild
				continue
			} else {
				mostRight.RightChild = nil
			}
		} else {
			fmt.Printf("%d ", cur.Data)
		}
		cur = cur.RightChild
	}
}
func MiddleMorris(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	//cur指向头节点
	cur := node
	var mostRight *BinaryTreeNode = nil
	for cur != nil {
		//cur没有左孩子的话
		mostRight = cur.LeftChild
		if mostRight != nil {
			for mostRight.RightChild != nil && mostRight.RightChild != cur {
				mostRight = mostRight.RightChild
			}
			if mostRight.RightChild == nil {
				mostRight.RightChild = cur
				cur = cur.LeftChild
				continue
			} else {
				mostRight.RightChild = nil
			}
		}
		fmt.Printf("%d ", cur.Data)
		cur = cur.RightChild
	}
}
func LastMorris(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	//cur指向头节点
	cur := node
	var mostRight *BinaryTreeNode = nil
	for cur != nil {
		//cur没有左孩子的话
		mostRight = cur.LeftChild
		if mostRight != nil {
			for mostRight.RightChild != nil && mostRight.RightChild != cur {
				mostRight = mostRight.RightChild
			}
			if mostRight.RightChild == nil {
				mostRight.RightChild = cur
				cur = cur.LeftChild
				continue
			} else {
				mostRight.RightChild = nil
				printEdge(cur.LeftChild)
			}
		}

		cur = cur.RightChild
	}
	printEdge(node)
	fmt.Println()
}
func printEdge(node *BinaryTreeNode) {
	tail := reserve(node)
	for tail != nil {
		fmt.Printf("%d ", tail.Data)
		tail = tail.RightChild
	}
	reserve(tail)
}
func reserve(node *BinaryTreeNode) *BinaryTreeNode {
	var (
		pre  *BinaryTreeNode
		next *BinaryTreeNode
	)
	for node != nil {
		next = node.RightChild
		node.RightChild = pre
		pre = node
		node = next
	}
	return pre
}
