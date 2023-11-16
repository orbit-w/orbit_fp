package maze

import (
	"fmt"
	"math/rand"
	"time"
)

/*
   @Author: orbit-w
   @File: maze
   @2023 11月 周四 15:13
*/

const (
	Path = iota
	Wall
	Begin
	End
)

type Maze struct {
	rows    int
	columns int
	grid    [][]int8
}

func New(rows, cols int) *Maze {
	m := &Maze{
		rows:    rows,
		columns: cols,
		grid:    make([][]int8, rows),
	}

	for i := range m.grid {
		m.grid[i] = make([]int8, cols)
		for j := range m.grid[i] {
			m.grid[i][j] = Wall
		}
	}
	return m
}

func (m *Maze) Gen() {
	rand.Seed(time.Now().UnixNano())
	row := rand.Intn(m.rows)
	col := rand.Intn(m.columns)
	m.grid[row][col] = Begin

	m.genMazeRecursive(row, col)

	row = rand.Intn(m.rows)
	col = rand.Intn(m.columns)
	m.grid[row][col] = End
}

func (m *Maze) genMazeRecursive(row, col int) {
	directions := rand.Perm(4)
	step := 2
	for _, dir := range directions {
		switch dir {
		case 0:
			next := row - step
			if next < 0 || m.grid[next][col] != Wall {
				continue
			}
			m.grid[next][col] = Path
			m.grid[row-1][col] = Path
			m.genMazeRecursive(next, col)

		case 1:
			next := row + step
			if next >= m.rows || m.grid[next][col] != Wall {
				continue
			}
			m.grid[next][col] = Path
			m.grid[row+1][col] = Path
			m.genMazeRecursive(next, col)

		case 2:
			next := col - 2
			if next < 0 || m.grid[row][next] != Wall {
				continue
			}
			m.grid[row][next] = Path
			m.grid[row][col-1] = Path
			m.genMazeRecursive(row, next)

		case 3:
			next := col + 2
			if next >= m.columns || m.grid[row][next] != Wall {
				continue
			}
			m.grid[row][next] = Path
			m.grid[row][col+1] = Path
			m.genMazeRecursive(row, next)
		}
	}
}

func (m *Maze) Print() {
	for _, row := range m.grid {
		for _, cell := range row {
			switch cell {
			case Wall:
				fmt.Printf("%c ", '\u25A0')
			case Path:
				fmt.Print("  ")
			case Begin:
				fmt.Print("S ")
			case End:
				fmt.Print("E ")
			}
		}
		fmt.Println()
	}
}
