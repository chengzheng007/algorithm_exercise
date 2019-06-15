package main

import "fmt"

func main() {
	graph := newUndirectGraph(8)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 3)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 5)
	graph.AddEdge(4, 6)
	graph.AddEdge(5, 7)
	graph.AddEdge(6, 7)

	graph.PrintEdge()

	fmt.Println("bfs test:")
	graph.BFS(0, 6)
	graph.BFS(2, 6)
	fmt.Println("\ndfs test:")
	graph.DFS(0, 6)
	graph.DFS(2, 6)
	
}
