package binarytree

import "fmt"

/**
二叉树基本操作
*/

// Node is
type Node struct {
	val    interface{}
	left   *Node
	right  *Node
	parent *Node
}

// getNodeHeight 返回输入节点的高度
func (n *Node) getNodeHeight() int {
	if n == nil {
		return 0
	}
	return max(n.left.getNodeHeight(), n.right.getNodeHeight()) + 1
}

// max return bigger number
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// getNodeDepth 返回输入节点的深度
func (n *Node) getNodeDepth() int {
	// solve-1 while
	d := 0
	for n.parent != nil {
		n = n.parent
		d++
	}
	return d
}

// preOrder 前序遍历二叉树 递归方式
func (n *Node) preOrder() {
	if n != nil {
		fmt.Println(n.val)
		n.left.preOrder()
		n.right.preOrder()
	}
}

// inOrder 中序遍历二叉树 递归方式
func (n *Node) inOrder() {
	if n != nil {
		n.left.inOrder()
		fmt.Println(n.val)
		n.right.inOrder()
	}
}

// postOrder 后序遍历二叉树 递归方式
func (n *Node) postOrder() {
	if n != nil {
		n.left.postOrder()
		n.right.postOrder()
		fmt.Println(n.val)
	}
}

// 层次遍历
