package main

import (
	"container/list"
	"fmt"
)

/**
对无向图(全联通图）的dfs
*/
const (
	DOING    = "DOING"
	DONE     = "DONE"
	DOINGNOW = "DOINGNOW"
	TODO     = "TODO"
)

// DfsWithRecursion will print node from starter's distance with dfs order
func DfsWithRecursion(g *Graph, starter int) {
	if g.eNum == 0 {
		return
	}
	fmt.Printf("DFS begin with node %d\n", starter)
	nodeState := make(map[int]string, 0)
	distance := make(map[int]int, 0)
	distance[starter] = 0
	dfs(g, starter, nodeState, distance, starter)
}

func dfs(g *Graph, current int, nodeState map[int]string, distance map[int]int, starter int) {
	nodeState[current] = DOING
	currentNode := g.adjacencyList[current]
	fmt.Printf("node %d from node %d 's distance is: %d\n", current, starter, distance[current])
	adjNode := currentNode.edgeList.Front()
	for adjNode != nil {
		adjNodeVal := adjNode.Value.(int)
		state, ok := nodeState[adjNodeVal]
		if state == TODO || !ok {
			distance[adjNodeVal] = distance[current] + 1
			dfs(g, adjNodeVal, nodeState, distance, starter)
		}
		adjNode = adjNode.Next()
	}
	nodeState[current] = DONE
}

// DfsWithIteration will print node from starter's distance with dfs order; 使用栈实现
func DfsWithIteration(g *Graph, starter int) {
	if g.eNum == 0 {
		return
	}
	fmt.Printf("DFS begin with node %d\n", starter)
	nodeState := make(map[int]string, 0)
	distance := make(map[int]int, 0)
	stack := make([]int, 0)

	distance[starter] = 0
	nodeState[starter] = DOING
	stack = append(stack, starter)

	for len(stack) != 0 {
		curNodeVal := stack[len(stack)-1]
		adjNode := nodeNextAdj(g, curNodeVal, nodeState)
		if adjNode != nil {
			adjVal := adjNode.Value.(int)
			nodeState[adjVal] = DOING
			distance[adjVal] = distance[curNodeVal] + 1
			stack = append(stack, adjVal)
		} else {
			stack = stack[0 : len(stack)-1]
			nodeState[curNodeVal] = DONE
			fmt.Printf("node %d from node(starter) %d 's distance is : %d\n", curNodeVal, starter, distance[curNodeVal])
		}
	}

}

func nodeNextAdj(g *Graph, node int, nodeState map[int]string) *list.Element {
	adj := g.adjacencyList[node].edgeList.Front()
	for adj != nil {
		val := adj.Value.(int)
		if state, ok := nodeState[val]; state == TODO || !ok {
			return adj
		}
		adj = adj.Next()
	}
	return nil
}
