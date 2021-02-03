package list

import (
	"errors"
)

/**
实现标准的双向链表
*/

// Node 是双向链表的一个节点
type node struct {
	next, prev *node
	val        interface{}
}

type list struct {
	len        int
	head, tail *node
}

func initList() *list {
	return &list{
		len:  0,
		head: nil,
		tail: nil,
	}
}

func (l *list) Insert(i int, ele interface{}) error {

	if i > l.len || i < 0 {
		return errors.New("index out of range")
	}
	newNode := &node{
		val:  ele,
		next: nil,
		prev: nil,
	}
	// 头部插入
	if i == 0 {
		newNode.next = l.head
		l.head = newNode
		l.head.prev = nil
		l.len++
		return nil
	}
	// 尾部插入
	if i == l.len {
		newNode.prev = l.tail
		l.tail = newNode
		l.tail.next = nil
		l.len++
		return nil
	}
	// 中间插入
	pre := l.head
	for j := 1; j < i; j++ {
		pre = pre.next
	}
	temp := pre.next
	pre.next = newNode
	newNode.prev = pre
	newNode.next = temp
	temp.prev = newNode
	l.len++
	return nil

}
