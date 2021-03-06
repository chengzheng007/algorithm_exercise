package main

import (
	"container/list"
	"fmt"
)

// 拓扑排序

// 有向图
type DirectedGraph struct {
	// 定点个数
	vertexNum int
	// 邻接表
	adj [][]int
}

func newDirectedGraph(num int) *DirectedGraph {
	if num <= 0 {
		panic("invalid num")
		return nil
	}
	dg := &DirectedGraph{
		vertexNum: num,
		adj:       make([][]int, num),
	}
	return dg
}

func (g *DirectedGraph) AddEdge(s, t int) bool {
	if s >= g.vertexNum || t >= g.vertexNum {
		return false
	}
	// 有向图，只加一边
	// s先于t执行，s->t
	g.adj[s] = append(g.adj[s], t)
	return true
}

func (g *DirectedGraph) TopoSortByKahn() {
	// 构建每个顶点的入度数统计（出度:某顶点指向其他顶点的总数，入度相反）
	inDgree := make([]int, g.vertexNum)
	for _, adjList := range g.adj {
		for _, w := range adjList {
			// s->w
			inDgree[w]++
		}
	}
	queue := list.New()

	// 先将入度为0的点加入队列，这些顶点不依赖于任何其他顶点
	for v, inDgreeNum := range inDgree {
		if inDgreeNum == 0 {
			queue.PushBack(v)
		}
	}

	for queue.Len() > 0 {
		vElem := queue.Front()
		v := vElem.Value.(int)
		// 获取入度为0的顶点后，需将其删除
		queue.Remove(vElem)
		fmt.Printf("-> %d ", v)
		for _, w := range g.adj[v] {
			inDgree[w]--
			if inDgree[w] == 0 {
				queue.PushBack(w)
			}
		}
	}
	fmt.Println()
}

// dfs实现拓扑排序
func (g *DirectedGraph) TopoSortByDFS() {
	// 构建逆邻接表
	inverseAdj := make([][]int, g.vertexNum)
	for i, adjList := range g.adj {
		for _, v := range adjList {
			//i->v
			inverseAdj[v] = append(inverseAdj[v], i) // v->i
		}
	}
	// 逆邻接表里，都是指向当前节点的顶点
	//fmt.Printf("inverse adj:%v\n", inverseAdj)
	visited := make([]bool, g.vertexNum)
	for i := 0; i < g.vertexNum; i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		g.topoDfs(i, inverseAdj, visited)
	}
	fmt.Println()
}

func (g *DirectedGraph) topoDfs(vertex int, inverseAdj [][]int, visited []bool) {
	for _, vid := range inverseAdj[vertex] {
		if visited[vid] {
			continue
		}
		visited[vid] = true
		g.topoDfs(vid, inverseAdj, visited)
	}
	// 先打印指向自己的顶点（自己依赖的顶点），最后再打印自己
	fmt.Printf("->%d ", vertex)
}
