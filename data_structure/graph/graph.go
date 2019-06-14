package main

import (
	"container/list"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"strconv"
)

// 无向图
type UndirectGraph struct {
	// 顶点个数（约定顶点编号为[0,vertexNum-1]）
	vertexNum int
	// 邻接表，存储的是图中每一个顶点到其可访问的顶点列表
	adj [][]int
}

func newUndirectGraph(v int) *UndirectGraph {
	if v <= 0 {
		panic("invalid v")
	}
	graph := &UndirectGraph{
		vertexNum:v,
		adj:make([][]int, v),
	}
	return graph
}

func (g *UndirectGraph) AddEdge(s, t int) bool {
	if s >= g.vertexNum || t >= g.vertexNum {
		return false
	}
	// 无向图一条边添加两次
	g.adj[s] = append(g.adj[s], t)
	g.adj[t] = append(g.adj[t], s)
	return true
}

func (g UndirectGraph) PrintEdge() {
	str := ""
	for i, lst := range g.adj {
		data, err := json.Marshal(lst)
		if err != nil {
			fmt.Printf("PrintEdge() error(%v)\n", err)
			return
		}
		str += strconv.Itoa(i)+":"+string(data)+"\n"
	}
	fmt.Printf("adj list:\n%s\n", str)
}

// breadth first search，广度优先搜索，寻找两点直径路径
func (g *UndirectGraph) BFS(s, t int) {
	if s >= g.vertexNum || t >= g.vertexNum {
		fmt.Printf("one of the point(%v, %v) not in graph.\n", s, t)
		return
	}
	if s == t {
		fmt.Println("start point equal to end point.")
		return
	}
	// 存储访问路径
	prev := make([]int, g.vertexNum)
	for i := 0; i < g.vertexNum; i++ {
		prev[i] = -1
	}
	// 存储节点是否被访问过
	visited := make([]bool, g.vertexNum)
	// 广度搜索依赖的队列
	queue := list.New()

	visited[s] = true
	queue.PushBack(s)

	for queue.Len() > 0 {
		wElem := queue.Front()
		w := wElem.Value.(int)
		queue.Remove(wElem)
		for _, q := range g.adj[w] {
			if visited[q] {
				continue
			}
			prev[q] = w
			if q == t {
				fmt.Printf("find the path %d to %d:", s, t)
				printPath(prev, s, t)
				fmt.Println()
				return
			}
			visited[q] = true
			queue.PushBack(q)
		}
	}
	fmt.Println("no path found.")
}


func printPath(path []int, s, t int) {
	if t != s {
		printPath(path, s, path[t])
	}
	fmt.Printf("->%d ", t)
}

