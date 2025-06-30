package graph

import (
	"container/heap"
	"math"
)

type Edge struct {
	to, weight int
}

// 邻接表
type AdjacencyMatrixGraph struct {
	vertices int
	edges    map[int][]Edge
}

type AdjacencyListGraph struct {
}

type Node struct {
	vertex int
	dist   int
}

type PriorityQueue []*Node //pointer

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

// 有向图
func NewAdjacencyMatrixGraph() *AdjacencyMatrixGraph {
	g := &AdjacencyMatrixGraph{
		vertices: 6,
		edges:    make(map[int][]Edge),
	}
	g.edges[0] = []Edge{{1, 2}, {2, 4}}
	g.edges[1] = []Edge{{2, 1}, {3, 7}}
	g.edges[2] = []Edge{{4, 3}}
	g.edges[3] = []Edge{{5, 1}}
	g.edges[4] = []Edge{{3, 2}, {5, 5}}

	return g
}

func (g *AdjacencyMatrixGraph) Dijkstra(start int) ([]int, []int) {
	dist := make([]int, g.vertices)
	prev := make([]int, g.vertices)
	visited := make([]bool, g.vertices)

	for i := 0; i < g.vertices; i++ {
		dist[i] = math.MaxInt32
		prev[i] = -1
	}
	dist[start] = 0
	pq := &PriorityQueue{} //init
	heap.Push(pq, &Node{start, 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Node)
		if !visited[cur.vertex] {
			visited[cur.vertex] = true //优先队列在距离最小时标记，其他情况过滤
			for _, e := range g.edges[cur.vertex] {
				if newDist := dist[cur.vertex] + e.weight; newDist < dist[e.to] {
					dist[e.to] = newDist
					prev[e.to] = cur.vertex
				}
				heap.Push(pq, &Node{e.to, dist[e.to]})
			}
		}
	}

	return dist, prev
}

// Breadth First Search
func (g *AdjacencyMatrixGraph) BFS(start, end int) []int {
	prev := make([]int, g.vertices)
	visited := make([]bool, g.vertices)
	path := []int{start}
	found := false
	for i := 0; i < g.vertices; i++ {
		prev[i] = -1
	}
	visited[start] = true //队列需要入队前标记

	pq := []int{start}
	for len(pq) > 0 {
		cur := pq[0]
		if cur == end {
			found = true
			break
		}
		pq = pq[1:]

		for _, e := range g.edges[cur] {
			if !visited[e.to] {
				visited[e.to] = true
				path = append(path, e.to)
				prev[e.to] = cur
				pq = append(pq, e.to)
			}
		}
	}

	//fmt.Println(path)
	if found {
		path = getPath(prev, start, end)
	}

	return path
}

type PriorityQueueMST []*Edge

func (pq PriorityQueueMST) Len() int           { return len(pq) }
func (pq PriorityQueueMST) Less(i, j int) bool { return pq[i].weight < pq[j].weight }
func (pq PriorityQueueMST) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueueMST) Push(x interface{}) {
	*pq = append(*pq, x.(*Edge))
}
func (pq *PriorityQueueMST) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// prim mini span tree
func (g *AdjacencyMatrixGraph) PrimMST(start int) []Edge {
	visited := make([]bool, g.vertices)
	path := []Edge{}

	pq := &PriorityQueueMST{}
	heap.Push(pq, &Edge{start, 0})
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Edge)
		if visited[cur.to] {
			continue
		}
		visited[cur.to] = true
		path = append(path, *cur)
		for _, e := range g.edges[cur.to] {
			if !visited[e.to] {
				heap.Push(pq, &e)
			}
		}
	}

	return path
}

func getPath(prev []int, start, end int) []int {
	path := []int{end}
	for i := end; i != start; i = prev[i] {
		path = append([]int{prev[i]}, path...)
	}

	return path
}
