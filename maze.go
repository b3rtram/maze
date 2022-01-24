package maze

import (
	"fmt"
	"math/rand"
	"time"
)

type vec struct {
	x int
	y int
	d int
}

const (
	U = 2
	L = 4
	R = 8
	D = 16
)

var dir = []vec{
	{x: 1, y: 0, d: R},
	{x: -1, y: 0, d: L},
	{x: 0, y: -1, d: U},
	{x: 0, y: 1, d: D},
}

type Maze struct {
	w int
	h int
}

func (m *Maze) GenerateMaze(height int, width int) [][]int {

	m.w = width
	m.h = height

	var cells []vec
	grid := make([][]int, m.h)

	for i := range grid {
		grid[i] = make([]int, m.w)
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	pos := vec{x: r.Intn(m.w), y: r.Intn(m.h)}
	grid[pos.y][pos.x] = 1
	cells = append(cells, vec{x: pos.x, y: pos.y})
	i := 2

	for len(cells) > 0 {

		r.Shuffle(len(dir), func(i int, j int) {
			dir[i], dir[j] = dir[j], dir[i]
		})

		pos = cells[len(cells)-1]
		suc := false

		for _, d := range dir {

			//... movement is inside grid
			if pos.x+d.x >= 0 && pos.y+d.y >= 0 && pos.x+d.x < m.w && pos.y+d.y < m.h && grid[pos.y+d.y][pos.x+d.x] == 0 {

				grid[pos.y][pos.x] |= d.d

				pos.x += d.x
				pos.y += d.y
				grid[pos.y][pos.x] = 1
				cells = append(cells, vec{x: pos.x, y: pos.y})

				suc = true
				break
			}
		}

		if !suc {
			cells = cells[:len(cells)-1]
		} else {
			i++
		}
	}

	return grid
}

func (m *Maze) PrintConsole(grid [][]int) {

	for x := 0; x < m.w; x++ {

		fmt.Print("\033[4m")
		fmt.Print("  ")

	}

	fmt.Print(" ")
	fmt.Print("\033[0m")

	fmt.Println()
	for y := 0; y < m.h; y++ {
		fmt.Print("|")

		for x := 0; x < m.w; x++ {

			rBlocked := true
			dBlocked := true

			if grid[y][x]&R == R {
				rBlocked = false
			}

			if x+1 < m.w && grid[y][x+1]&L == L {
				rBlocked = false
			}

			if y+1 < m.h && grid[y+1][x]&U == U || grid[y][x]&D == D {
				dBlocked = false
			}

			if dBlocked {
				fmt.Print("\033[4m")
			}

			fmt.Print(" ")

			if x+1 < m.w {
				if rBlocked {

					fmt.Print("\033[0m")
					fmt.Print("|")
				} else {
					fmt.Print(" ")
				}
			}

			fmt.Print("\033[0m")
		}

		fmt.Println("|")
	}

}
