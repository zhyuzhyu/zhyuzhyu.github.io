package graph

import (
	"fmt"
	"testing"
)

var g = NewAdjacencyMatrixGraph()

func TestDijkstra(t *testing.T) {
	start := 0
	dist, prev := g.Dijkstra(start)

	//输出最短距离
	fmt.Println("顶点\t最短距离\t路径")
	for i := 0; i < g.vertices; i++ {
		fmt.Printf("%d\t%d\t\t", i, dist[i])
		fmt.Println(getPath(prev, start, i))
	}
}

func TestBFS(t *testing.T) {
	start := 0
	fmt.Println("顶点\t路径")
	for i := -1; i < g.vertices; i++ {
		path := g.BFS(start, i)
		//输出路径
		fmt.Printf("%d\t", i)
		fmt.Println(path)
	}
}

func TestPrimMST(t *testing.T) {
	g := &AdjacencyMatrixGraph{
		vertices: 9,
		edges:    make(map[int][]Edge),
	}
	g.edges[0] = []Edge{{1, 3}, {5, 4}}
	g.edges[1] = []Edge{{0, 3}, {2, 8}, {6, 6}, {8, 5}}
	g.edges[2] = []Edge{{1, 8}, {3, 12}, {8, 2}}
	g.edges[3] = []Edge{{2, 11}, {4, 10}, {6, 14}, {7, 6}, {8, 11}}
	g.edges[4] = []Edge{{3, 10}, {5, 18}, {7, 1}}
	g.edges[5] = []Edge{{0, 4}, {4, 18}, {6, 7}}
	g.edges[6] = []Edge{{1, 6}, {3, 14}, {5, 7}, {7, 9}}
	g.edges[7] = []Edge{{3, 6}, {4, 1}, {6, 9}}
	g.edges[8] = []Edge{{1, 5}, {2, 2}, {3, 11}}

	mst := g.PrimMST(0)
	fmt.Println(mst)
}
