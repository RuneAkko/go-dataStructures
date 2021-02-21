package main

func main() {
	g := NewGraph()
	g.addVertex(0)
	for _, ele := range []int{1, 2, 3, 4, 5, 6, 7} {
		g.addVertex(ele)
	}
	edges := [][]int{
		[]int{0, 1},
		[]int{0, 2},
		[]int{0, 3},
		[]int{1, 4},
		[]int{2, 5},
		[]int{3, 6},
		[]int{5, 6},
		[]int{4, 7},
	}
	for _, ele := range edges {
		g.addEdge(ele[0], ele[1])
	}
	Bfs(g, 0)
	DfsWithRecursion(g, 0)
	DfsWithIteration(g, 0)

}
