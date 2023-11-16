package a_star

import (
	"testing"
)

/*
   @Author: orbit-w
   @File: astar_test
   @2023 11月 周四 12:21
*/

func TestFindingPath_Finding(t *testing.T) {
	matrix := [][]uint8{
		{0, 1, 0, 0, 0}, // here node {0,2} is a forced neighbour for  {1,1}
		{0, 0, 1, 0, 0},
		{1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}

	fp := new(FindingPath)
	fp.Build(matrix)
}
