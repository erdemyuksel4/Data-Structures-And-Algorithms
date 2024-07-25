package structures

import "fmt"

type BinaryTreeNode struct {
	value int
	right *BinaryTreeNode
	left  *BinaryTreeNode
}

func (node *BinaryTreeNode) Value() int             { return node.value }
func (node *BinaryTreeNode) Right() *BinaryTreeNode { return node.right }
func (node *BinaryTreeNode) Left() *BinaryTreeNode  { return node.left }
func (n *BinaryTreeNode) Insert(node *BinaryTreeNode) {
	if n.value < node.value {
		n.InsertRight(node)
	} else {
		n.InsertLeft(node)
	}
}
func (n *BinaryTreeNode) InsertRight(node *BinaryTreeNode) {
	if n.right != nil {
		node.right = n.right
	}
	n.right = node
}
func (n *BinaryTreeNode) InsertLeft(node *BinaryTreeNode) {
	if n.left != nil {
		node.left = n.left
	}
	n.left = node
}
func (n *BinaryTreeNode) Depth() int {
	if n == nil {
		return 0
	}
	leftDepth := n.left.Depth()
	rightDepth := n.right.Depth()

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}
func NewBinaryTreeNode(value int) *BinaryTreeNode {
	return &BinaryTreeNode{value: value}
}
func (node *BinaryTreeNode) PreOrderTraversal() {
	if node == nil {
		return
	}
	fmt.Printf("%d ,", node.value)
	node.left.PreOrderTraversal()
	node.right.PreOrderTraversal()
}
func (node *BinaryTreeNode) InOrderTraversal() {
	if node == nil {
		return
	}
	node.left.InOrderTraversal()
	fmt.Printf("%d ,", node.value)
	node.right.InOrderTraversal()
}
func (node *BinaryTreeNode) PostOrderTraversal() {
	if node == nil {
		return
	}
	node.left.PostOrderTraversal()
	node.right.PostOrderTraversal()
	fmt.Printf("%d ,", node.value)
}
