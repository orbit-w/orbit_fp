package dijkstra

/*
   @Author: orbit-w
   @File: graph
   @2023 11月 周二 16:22
*/

import (
	"github.com/orbit-w/golib/bases/container/heap"
	"math"
)

type IDijkstra interface {
	//AddVertex 添加顶点
	AddVertex(id int)
	GetVertex(id int) (int, bool)

	//AddEdge 添加边
	AddEdge(src, dest int, weight int64) error
	//GetEdge 获取边
	GetEdge(src, dest int) (w int64, exist bool)

	//ShortestPath 是从其中一个顶点到目标顶点的最短路径
	ShortestPath(start, dest int) (Result, bool)
}

//Graph 加权无向图
type Graph struct {
	vertices []Vertex
	visited  []bool
	pq       *heap.Heap[int, int64]
}

func New() IDijkstra {
	return &Graph{
		vertices: make([]Vertex, 0, initializedSize),
		visited:  make([]bool, 0, initializedSize),
		pq:       &heap.Heap[int, int64]{},
	}
}

func (g *Graph) AddVertex(id int) {
	v := Vertex{
		Id:   id,
		Best: defaultBest,
	}
	if v.Id >= len(g.vertices) {
		g.malloc(v.Id + 1 - len(g.vertices))
	}
	g.vertices[v.Id] = v
}

func (g *Graph) GetVertex(id int) (vId int, exist bool) {
	if id >= len(g.vertices) {
		return
	}
	vId = g.vertices[id].Id
	return vId, vId == id
}

//AddEdge 无向图中添加两个结点的边
func (g *Graph) AddEdge(src, dest int, weight int64) error {
	if !g.Exist(src) || g.Exist(dest) {
		return ErrVertexNotFound
	}

	if src == dest {
		return ErrVertexInvalid
	}

	vertex := &g.vertices[src]
	if vertex.Edges == nil {
		vertex.Edges = make(map[int]int64)
	}
	vertex.Edges[dest] = weight
	return nil
}

func (g *Graph) GetEdge(src, dest int) (w int64, exist bool) {
	if !g.Exist(src) || g.Exist(dest) {
		return
	}

	if src == dest {
		return
	}

	w, exist = g.vertices[src].Edges[dest]
	return
}

func (g *Graph) ShortestPath(start, dest int) (Result, bool) {
	g.prepare()
	g.pq.Push(&heap.Item[int, int64]{Value: start, Priority: 0})
	g.vertices[start].Dist = 0

	var (
		found bool
		bd    int64
	)

	for g.pq.Len() > 0 {
		head := g.pq.Pop()
		cId := head.Value
		if g.visited[cId] {
			continue
		}
		cur := g.vertices[cId]
		g.visited[cId] = true
		for id, weight := range cur.Edges {
			vertex := &g.vertices[id]
			u := cur.Dist + weight
			if u < vertex.Dist {
				vertex.Dist = u
				vertex.Best = cId
			}
			if id == dest {
				found = true
				bd = u
			}
		}
	}

	if found {
		return g.pathing(start, dest, bd), found
	}
	return Result{}, found
}

func (g *Graph) Exist(id int) bool {
	if id >= len(g.vertices) {
		return false
	}
	v := g.vertices[id]
	return v.Id == id
}

func (g *Graph) prepare() {
	for i := range g.vertices {
		v := &g.vertices[i]
		v.Dist = math.MaxInt64
		v.Best = defaultBest
		g.visited[i] = false
	}

	for {
		if g.pq.Len() == 0 {
			break
		}
		g.pq.Pop()
	}
}

func (g *Graph) pathing(start, dest int, bd int64) Result {
	r := Result{
		Distance: bd,
		Path:     make([]int, 0, 1<<3),
	}

	prev := g.vertices[dest].Best
	for ; prev != start; prev = g.vertices[prev].Best {
		r.Path = append(r.Path, prev)
	}
	r.Path = append(r.Path, start)
	reverse(r.Path)
	return r
}

func (g *Graph) malloc(size int) {
	newVertices := make([]Vertex, size)
	g.vertices = append(g.vertices, newVertices...)

	newVisited := make([]bool, size)
	g.visited = append(g.visited, newVisited...)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
