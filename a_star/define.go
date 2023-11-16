package a_star

/*
   @Author: orbit-w
   @File: define
   @2023 11月 周三 18:30
*/

type Direction struct {
	H, V int32
}

var (
	zero = struct{}{}

	directions = []Direction{
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}
)
