package dijkstra

/*
   @Author: orbit-w
   @File: base
   @2023 11月 周三 16:53
*/

type Result struct {
	Distance int64
	Path     []int
}

const (
	initializedSize = 1 << 3

	defaultBest = -1
)
