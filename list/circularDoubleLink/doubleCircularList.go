package list

/**
带有哨兵
*/
import (
	"errors"
)

type node struct {
	next, prev *node
	val        interface{}
}

type list struct {
	root node //哨兵，位于表头和表尾之间
	len  int
}

func (l *list) init() *list {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func (l *list) front() *node {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *list) back() *node {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// 在表i位置处插入新节点
func (l *list) insert(i int, v interface{}) error {
	if i < 0 || i > l.len {
		return errors.New("index i is valid")
	}
	newNode := &node{
		val:  v,
		next: nil,
		prev: nil,
	}
	// 找到新节点插入位置之前的节点 i-1
	pre := &l.root
	for j := 0; j < i; j++ {
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
