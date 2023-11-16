package dijkstra

/*
   @Author: orbit-w
   @File: vertex
   @2023 11月 周四 16:38
*/

type Vertex struct {
	Id    int
	Best  int
	Dist  int64
	Edges map[int]int64 //维护了相关的所有边
}
