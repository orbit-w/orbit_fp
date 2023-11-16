package a_star

/*
   @Author: orbit-w
   @File: a_star
   @2023 11月 周三 17:07
*/

import (
	"github.com/orbit-w/golib/bases/container/heap_list"
	"github.com/orbit-w/golib/bases/misc/number_utils"
	"github.com/orbit-w/orbit-fp/lib/inside/common"
)

type Rec struct {
	p common.Pos
	g int32
}

type FindingPath struct {
	open    *heap_list.HeapList[common.PosNum, common.Pos, int32]
	close   map[common.PosNum]struct{}
	gMap    map[common.PosNum]int32
	prevMap map[common.PosNum]common.PosNum
	rect    common.Rect
}

func (fp *FindingPath) Build(m [][]uint8) {
	fp.open = heap_list.New[common.PosNum, common.Pos, int32]()
	fp.gMap = make(map[common.PosNum]int32, 1<<3)
	fp.prevMap = make(map[common.PosNum]common.PosNum, 1<<3)
	fp.close = make(map[common.PosNum]struct{}, 1<<3)
	fp.rect = common.NewRect(m)
}

func (fp *FindingPath) Finding(m [][]uint8, s common.Pos, e common.Pos, collector func(x, y int32)) {
	fp.open.Push(s.Id(), s, 0)
	for !fp.open.Empty() {
		id, cur, _ := fp.open.Pop()
		cId := cur.Id()
		curG := fp.gMap[cId]
		fp.close[id] = zero

		for _, d := range directions {
			neighbor := common.Pos{X: cur.X + d.H, Y: cur.Y + d.V}
			if neighbor.Equal(&e) {
				fp.prevMap[e.Id()] = cId
				return
			}

			if fp.rect.Invalid(&neighbor) {
				continue
			}

			if m[neighbor.Y][neighbor.X] > 0 {
				continue
			}

			nId := neighbor.Id()
			if _, ok := fp.close[nId]; ok {
				continue
			}
			h := heuristic(neighbor, e)
			g := calcG(curG, &d)

			if exist := fp.open.Exist(nId); exist {
				if rG := fp.gMap[nId]; rG > g {
					fp.gMap[nId] = g
					fp.prevMap[nId] = cId
					fp.open.UpdatePriority(nId, h+g)
				}
			} else {
				fp.gMap[nId] = g
				fp.prevMap[nId] = cId
				fp.open.Push(nId, neighbor, h+g)
			}
		}
	}

	fp.backtrack(s, e, collector)
}

func (fp *FindingPath) backtrack(s, e common.Pos, collector func(x, y int32)) {
	index := e.Id()
	for {
		prev, ok := fp.prevMap[index]
		if !ok {
			break
		}
		if prev.Equal(&s) {
			break
		}

		collector(prev.XY())
	}
}

func heuristic(cur, e common.Pos) int32 {
	abs := number_utils.ABS[int32]
	return 10 * (abs(e.X-cur.X) + abs(e.Y-cur.Y))
}

func calcG(prevG int32, d *Direction) (g int32) {
	abs := number_utils.ABS[int32]
	g = prevG
	if abs(d.H) == abs(d.V) {
		g += 14
	} else {
		g += 10
	}
	return
}
