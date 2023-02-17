package main

type BinaryTreeNode struct {
	LeftChild  *BinaryTreeNode
	RightChild *BinaryTreeNode
	Data       any
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryNode(data any) (node *BinaryTreeNode) {
	node = &BinaryTreeNode{
		Data: data,
	}
	return
}
func NewBinaryTree(data any) (tree *BinaryTree) {
	tree = &BinaryTree{
		Root: NewBinaryNode(data),
	}
	return
}
func (b *BinaryTreeNode) InsertLeft(data any) *BinaryTreeNode {
	node := &BinaryTreeNode{
		Data: data,
	}
	b.LeftChild = node
	return node
}
func (b *BinaryTreeNode) InsertRight(data any) *BinaryTreeNode {
	node := &BinaryTreeNode{
		Data: data,
	}
	b.RightChild = node
	return node
}
func (b *BinaryTree) InsertLeft(e *BinaryTreeNode, data any) *BinaryTreeNode {
	return e.InsertLeft(data)
}
func (b *BinaryTree) InsertRight(e *BinaryTreeNode, data any) *BinaryTreeNode {
	return e.InsertRight(data)
}
func createTest() *BinaryTree {
	tree := NewBinaryTree(1)
	tree.InsertLeft(tree.Root, 2)
	tree.InsertRight(tree.Root, 3)
	tree.Root.LeftChild.InsertLeft(4)
	tree.Root.LeftChild.InsertRight(5)
	tree.Root.RightChild.InsertLeft(6)
	tree.Root.RightChild.InsertRight(7)
	return tree
}
