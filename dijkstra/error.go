package dijkstra

import "errors"

/*
   @Author: orbit-w
   @File: error
   @2023 11月 周四 13:05
*/

var (
	ErrVertexNotFound = errors.New("err_vertex_not_found")
	ErrVertexInvalid  = errors.New("err_vertex_invalid")
)
