package list

import (
	"errors"
)

/**
实现一个单链表。
利用array实现单链表十分愚蠢，故不考虑这种实现。
*/

// Node 链表节点含有 数据域和指针域
type Node struct {
	val  interface{}
	next *Node
}

// LinkList 含有头结点指针和链表长度（含有的节点个数）
type LinkList struct {
	head *Node
	len  uint
}

// InitLinkList 返回一个新的链表
func InitLinkList() *LinkList {
	return &LinkList{
		head: &Node{
			val:  nil,
			next: nil,
		},
		len: 0,
	}
}

// IsEmpty 判断链表是否为空
func (l *LinkList) IsEmpty() bool {
	if l.len == 0 {
		return true
	}
	return false
}

// Insert 在链表位置i处插入新节点
func (l *LinkList) Insert(i uint, e interface{}) error {
	if i < 0 || i > l.len {
		return errors.New("index out of range")
	}
	newNode := &Node{
		val:  e,
		next: nil,
	}
	//头部插入
	if i == 0 {
		newNode.next = l.head
		l.head = newNode
		l.len++
		return nil
	}
	// 尾部插入
	pre := l.head
	if i == l.len {
		for pre.next != nil {
			pre = pre.next
		}
		pre.next = newNode
		newNode.next = nil
		l.len++
		return nil
	}
	// 中间插入
	for j := uint(1); j < i; j++ {
		pre = pre.next
	}
	newNode.next = pre.next
	pre.next = newNode
	l.len++
	return nil
}

// Delete 删除链表中位置为i的节点
func (l *LinkList) Delete(i uint) error {
	if l.IsEmpty() {
		return errors.New("list is empty")
	}
	if i < 0 || i >= l.len {
		return errors.New("index i is not valid")
	}
	// 删除头节点
	if i == 0 {
		l.head = l.head.next
		l.len--
		return nil
	}
	pre := l.head
	// 删除尾节点
	if i == l.len-1 {
		for j := uint(1); j < i; j++ {
			pre = pre.next
		}
		pre.next = nil
		l.len--
		return nil
	}
	// 删去中间节点
	for j := uint(1); j < i; j++ {
		pre = pre.next
	}
	pre.next = pre.next.next
	l.len--
	return nil
}
