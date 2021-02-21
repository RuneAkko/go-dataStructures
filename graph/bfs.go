package main

/**
对无向图(全联通图）的bfs
*/

import (
	"fmt"
)

// Bfs print Graph node with bfs order, with distance to starter
func Bfs(g *Graph, starter int) {
	if g.eNum == 0 {
		return
	}
	fmt.Printf("BFS begin with node %d\n", starter)
	isVisited := make(map[int]bool, g.vNum)
	distance := make(map[int]int, g.vNum)
	isVisited[starter] = true
	distance[starter] = 0
	var queue []Vertex
	queue = append(queue, g.adjacencyList[starter])
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:len(queue)]
		currentDis := distance[currentNode.value]
		fmt.Printf("node %d from %d (starter)'s distance is: %d\n", currentNode.value, starter, currentDis)
		adjNode := currentNode.edgeList.Front()
		for adjNode != nil {
			adjNodeVal := adjNode.Value.(int)
			if is, ok := isVisited[adjNodeVal]; !is || !ok {
				queue = append(queue, g.adjacencyList[adjNodeVal])
				isVisited[adjNodeVal] = true
				distance[adjNodeVal] = currentDis + 1
			}
			adjNode = adjNode.Next()
		}
	}
}
