package binarysearchtree

import (
	"errors"
)

type node struct {
	val                 int
	left, right, parent *node
}

// find 二叉搜索树的查找操作 递归
func (n *node) find(ele int) *node {
	if n == nil {
		return nil
	}
	if n.val == ele {
		return n
	}
	if n.val > ele {
		return n.left.find(ele)
	}
	return n.right.find(ele)
}

// insert 二叉搜索树的插入操作
func (n *node) insert(ele int) error {
	if n == nil {
		n.val = ele
		n.left, n.right, n.parent = nil, nil, nil
		return nil
	}
	pre := n
	for n != nil {
		if ele > n.val {
			pre = n
			n = n.right
			continue
		}
		if ele < n.val {
			pre = n
			n = n.left
			continue
		}
		if ele == n.val {
			return errors.New("insert ele conflict ")
		}
	}
	n.val = ele
	n.left, n.right = nil, nil
	n.parent = pre
	return nil
}

// delete 二叉搜索树的删除操作
func (n *node) delete(ele int) error {
	p := n.find(ele)
	if p == nil {
		return errors.New("ele do not exist")
	}

	if p.left == nil && p.right == nil {
		p.parent = nil
		return nil
	}

	if p.left != nil && p.right != nil {
		min := p.getMinimum()
		p.val = min.val
		if min.right != nil {
			min.right.parent = min.parent
		}
		min.parent = nil
		return nil
	}

	if p.left != nil {
		p.left.parent = p.parent
		p.parent = nil
		return nil
	}
	// p.right != nil
	p.right.parent = p.parent
	p.parent = nil
	return nil
}

// getMinimum return minimum node
func (n *node) getMinimum() *node {
	if n == nil {
		return nil
	}
	for n.left != nil {
		n = n.left
	}
	return n
}

// getMaximum return maximum node
func (n *node) getMaximum() *node {
	if n == nil {
		return nil
	}
	for n.right != nil {
		n = n.right
	}
	return n
}

//
func (n *node) preNode() *node {
	if n == nil {
		return nil
	}
	if n.left != nil {
		return n.left.getMaximum()
	}
	if n.parent == nil {
		// has no prenode
		return nil
	}
	if n.parent.right == n {
		return n.parent
	}
	p1 := n.parent
	p2 := p1.parent
	for p2 != nil {
		if p2.right == p1 {
			return p2
		}
		p1 = p2
		p2 = p2.parent
	}
	return nil
}

//
func (n *node) postNode() *node {
	if n == nil {
		return nil
	}
	if n.right != nil {
		return n.right.getMinimum()
	}
	if n.parent == nil {
		return nil
	}
	if n.parent.left == n {
		return n.parent
	}
	p1 := n.parent
	p2 := p1.parent
	for p2 != nil {
		if p2.left == p1 {
			return p2
		}
		p1 = p2
		p2 = p2.parent
	}
	return nil
}
