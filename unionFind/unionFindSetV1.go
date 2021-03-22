package main

/**
简单的并查集
*/

// Ele is
type Ele struct {
	key    int
	parent *Ele
}

// UnionFindSet is
type UnionFindSet struct {
	data []Ele
}

// NewUnionFindSet return
func NewUnionFindSet(elements []int) *UnionFindSet {
	var data []Ele
	for _, ele := range elements {
		e := Ele{
			key: ele,
		}
		e.parent = &e
		data = append(data, e)
	}
	return &UnionFindSet{
		data: data,
	}
}

func (ele *Ele) find() *Ele {
	if ele.parent.key == ele.key {
		return ele
	}
	return ele.parent.find()
}

func union(ele1, ele2 *Ele) {
	ele1.find().parent = ele2.find()
}
