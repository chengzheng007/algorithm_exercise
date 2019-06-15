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
	
	fmt.Println("\ntopological sortint test:")
	dg := newDirectedGraph(7)

	/*
	0 袜子
	1 鞋
	2 内裤
	3 长裤
	4 T恤
	5 外衣
	6 领带
	*/
	dg.AddEdge(4, 5)
	dg.AddEdge(5, 6)
	dg.AddEdge(2, 0)
	dg.AddEdge(2, 3)
	dg.AddEdge(0, 1)
	dg.AddEdge(3, 1)

	dg.TopoSortByKahn()
}
