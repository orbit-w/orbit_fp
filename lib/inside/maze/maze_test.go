package maze

import "testing"

/*
   @Author: orbit-w
   @File: maze_test
   @2023 11月 周五 19:23
*/

func TestMaze_Gen(t *testing.T) {
	m := New(20, 20)
	m.Gen()
	m.Print()
}
