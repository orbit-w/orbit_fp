package common

/*
   @Author: orbit-w
   @File: common
   @2023 11月 周三 17:16
*/

const (
	Mask = 0xFFFFFFFF
)

type Pos struct {
	X, Y int32
}

func (p *Pos) Id() (id PosNum) {
	id = (id|PosNum(p.X))<<32 | PosNum(p.Y)
	return
}

func (p *Pos) Equal(t *Pos) bool {
	return p.X == t.X && p.Y == t.Y
}

type PosNum uint64

func (i *PosNum) XY() (x, y int32) {
	y = int32(*i & Mask)
	x = int32((*i >> 32) & Mask)
	return
}

func (i *PosNum) Equal(t *Pos) bool {
	y := int32(*i & Mask)
	x := int32((*i >> 32) & Mask)
	return x == t.X && y == t.Y
}

type Rect struct {
	Width, Height int32
}

func NewRect(m [][]uint8) Rect {
	maxY := len(m)
	maxX := len(m[0])
	return Rect{
		Width:  int32(maxX),
		Height: int32(maxY),
	}
}

func (r *Rect) Invalid(p *Pos) bool {
	if p.X < 0 || p.X >= r.Width || p.Y < 0 || p.Y >= r.Height {
		return true
	}
	return false
}
