package main

import "testing"

//	    1
//	2       3
// 4   5  6    7
// 二叉树的递归序列
// 1 2 4 4 4 2 5 5 5 2 1 3 6 6 6 3 7 7 7 3 1
func TestRecurrence(t *testing.T) {
	tree := createTest()
	Recurrence(tree.Root)
}

// 先序遍历 1 2 4 5 3 6 7
func TestFirstRecurrence(t *testing.T) {
	tree := createTest()
	FirstRecurrence(tree.Root)
}

// 中序遍历 4 2 5 1 6 3 7
func TestSecondRecurrence(t *testing.T) {
	tree := createTest()
	MiddleRecurrence(tree.Root)
}

// 后序遍历4 5 2 6 7 3 1
func TestLastRecurrence(t *testing.T) {
	tree := createTest()
	LastRecurrence(tree.Root)
}
