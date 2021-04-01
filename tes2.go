package main

import "fmt"

type node struct {
	val int
}

func main() {
	m := make(map[int][]*node)
	m[1] = []*node{&node{val: 1}, &node{val: 2}}
	m[1][0] = &node{val: 3}
	fmt.Println(m[1][0])
}
