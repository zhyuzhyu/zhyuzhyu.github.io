package main

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
