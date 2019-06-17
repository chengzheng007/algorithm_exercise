package main

import (
	"math"
	"container/heap"
	"fmt"
)

type edge struct {
	sid int
	tid int
	w int  // 点sid到tid的权重，此处为路径长度
}

// 为实现算法定义的类
type vertex struct {
	index int       // go heap需要的参数
	vid   int       // 顶点编号
	distance int	// 从起始顶点到该顶点的距离
}

type vertexData struct {
	// 顶点id=>顶点，用于O(1)时间获取顶点信息
	mp map[int]*vertex
	// 建堆需要的列表
	list []*vertex
}

// 有权有向图
type weightdeDirectedGraph struct {
	vertexNum int
	adj [][]*edge
}

func newWeightdeDirectedGraph(num int) *weightdeDirectedGraph {
	if num <= 0 {
		panic("invalid num")
	}
	wdg := &weightdeDirectedGraph{
		vertexNum:num,
		adj:make([][]*edge, num),
	}
	return wdg
}

// 添加一条s指向t的边
func (wdg *weightdeDirectedGraph) AddEdge(s, t, w int) {
	wdg.adj[s] = append(wdg.adj[s], &edge{sid:s, tid:t, w:w})
}

func (wdg *weightdeDirectedGraph) Dijkstra(s, t int) {
	if s >= wdg.vertexNum || t >= wdg.vertexNum {
		fmt.Printf("one of the point(%v, %v) not in graph.\n", s, t)
		return
	}
	if s == t {
		fmt.Println("start point equal to end point.")
		return
	}

	// 存储路径
	path := make([]int, wdg.vertexNum)
	
	// 标记是否进入过堆
	inheap := make([]bool, wdg.vertexNum)
	
	vertexsPQ := priorityQueue{
		mp:make(map[int]*vertex, wdg.vertexNum),
	}
	for i := 0; i < wdg.vertexNum; i++ {
		vertexsPQ.mp[i] = &vertex{index:i, vid:i, distance:math.MaxInt64}
	}
	
	// 初始设置
	vertexsPQ.list = make([]*vertex, 1)
	// 起点s先入堆
	vertexsPQ.list[0] = &vertex{index:0, vid:s, distance:0}
	// 初始化堆
	heap.Init(&vertexsPQ)
	
	// s为已进入过优先级队列
	inheap[s] = true

	for vertexsPQ.Len() > 0 {
		minVertex := heap.Pop(&vertexsPQ).(*vertex)
		if minVertex.vid == t {  // 最短路径产生
			break
		}
		for _, edge := range wdg.adj[minVertex.vid] {
			nextVertex := vertexsPQ.mp[edge.tid]
			if minVertex.distance+edge.w < nextVertex.distance {
				vertexsPQ.mp[edge.tid].distance = minVertex.distance+edge.w
				path[nextVertex.vid] = minVertex.vid
				if inheap[nextVertex.vid] {
					// 更新优先级队列中的值
					vertexsPQ.update(nextVertex, nextVertex.vid, minVertex.distance+edge.w)
				} else {
					heap.Push(&vertexsPQ, nextVertex)
					inheap[nextVertex.vid] = true
				}
			}
		}
	}
	fmt.Printf("find the path %d to %d: ", s, t)
	printPath(path, s, t)
	fmt.Println()
}

type priorityQueue vertexData
func (pq priorityQueue) Len() int { return len(pq.list) }
func (pq priorityQueue) Less(i, j int) bool {
	// 小顶堆
	return pq.list[i].distance < pq.list[j].distance
}
func (pq priorityQueue) Swap(i, j int) {
	pq.list[i], pq.list[j] = pq.list[j], pq.list[i]
	pq.list[i].index = i
	pq.list[j].index = j
}
// 特殊处理，因为如果已存在，只需替换，不存在才加入
// 值的index应事先指定
func (pq *priorityQueue) Push(x interface{}) {
	n := len((*pq).list)
	item := x.(*vertex)
	item.index = n
	(*pq).list = append((*pq).list, item)
	pq.mp[item.vid] = item
}

func (pq *priorityQueue) Pop() interface{} {
	old := (*pq).list
	n := len(old)
	item := old[n-1]
	item.index = -1
	(*pq).list = old[0:n-1]
	delete(pq.mp, item.vid)
	return item
}

func (pq *priorityQueue) update(v *vertex, newVid, newDist int) {
	v.vid = newVid
	v.distance = newDist
	heap.Fix(pq, v.index)
	pq.mp[newVid] = v
}

 