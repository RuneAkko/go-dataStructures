package main

import "container/list"

// Graph is 无向图； 邻接表
type Graph struct {
	vNum          int
	eNum          int
	adjacencyList map[int]Vertex
}

// Vertex is
type Vertex struct {
	value    int
	edgeList *list.List
}

// NewGraph return
func NewGraph() *Graph {
	return &Graph{
		vNum:          0,
		eNum:          0,
		adjacencyList: make(map[int]Vertex, 0),
	}
}

func (g *Graph) addEdge(a, b int) {
	va, ok := g.adjacencyList[a]
	vb, ok := g.adjacencyList[b]
	if !ok {
		panic("v is not existed")
	}
	if !ok {
		panic("v is not existed")
	}

	ha := va.edgeList.Front()
	if ha != nil {
		if ha.Value == b {
			panic("added edge already existed")
		}
	}
	hb := vb.edgeList.Front()
	if hb != nil {
		if hb.Value == a {
			panic("added edge already existed")
		}
	}
	va.edgeList.PushBack(b)
	vb.edgeList.PushBack(a)
	g.eNum++
}

func (g *Graph) addVertex(a int) {
	if _, ok := g.adjacencyList[a]; ok {
		panic("v is existed")
	}
	g.vNum++
	g.adjacencyList[a] = Vertex{
		value:    a,
		edgeList: list.New(),
	}
}
