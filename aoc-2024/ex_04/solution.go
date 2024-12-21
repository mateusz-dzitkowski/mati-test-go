package ex_04

import (
	"fmt"
	"main/utils"
	"strings"
)

var Solution1 = []rune{'X', 'M', 'A', 'S'}
var Directions1 = []Point{
	{
		x: 1,
		y: 0,
	},
	{
		x: 1,
		y: 1,
	},
	{
		x: 0,
		y: 1,
	},
	{
		x: -1,
		y: 1,
	},
	{
		x: -1,
		y: 0,
	},
	{
		x: -1,
		y: -1,
	},
	{
		x: 0,
		y: -1,
	},
	{
		x: 1,
		y: -1,
	},
}

func Solve() {
	data := utils.LoadExercise("04")
	fmt.Println(solveFirst(data))
	fmt.Println(solveSecond(data))
}

func solveFirst(data string) int {
	output := 0
	grid := createGrid(data)
	for i, row := range grid {
		for j := range row {
			output += grid.countXmasAt1(i, j)
		}
	}
	return output
}

func solveSecond(data string) int {
	output := 0
	grid := createGrid(data)
	for i, row := range grid {
		for j := range row {
			if grid.isXmasAt(i, j) {
				output += 1
			}
		}
	}
	return output
}

type Grid [][]rune

func createGrid(data string) Grid {
	var grid Grid

	for _, stringRow := range strings.Split(data, "\n") {
		gridRow := []rune(stringRow)
		grid = append(grid, gridRow)
	}

	return grid
}

func (g Grid) countXmasAt1(i, j int) int {
	output := 0
	for _, direction := range Directions1 {
		isXmas := true
		for k, solutionLetter := range Solution1 {
			p := Point{i, j}.add(direction.times(k))
			if p.x < 0 || p.x >= len(g) {
				isXmas = false
				break
			}
			row := g[p.x]
			if p.y < 0 || p.y >= len(row) {
				isXmas = false
				break
			}
			if row[p.y] != solutionLetter {
				isXmas = false
				break
			}
		}
		if isXmas {
			output += 1
		}
	}
	return output
}

func (g Grid) isXmasAt(i, j int) bool {
	if g[i][j] != 'A' {
		return false
	}

	isXmas := true
	for _, directions := range [][]Point{
		{
			{
				x: 1,
				y: 1,
			},
			{
				x: -1,
				y: -1,
			},
		},
		{
			{
				x: -1,
				y: 1,
			},
			{
				x: 1,
				y: -1,
			},
		},
	} {
		forward := Point{i, j}.add(directions[0])
		backward := Point{i, j}.add(directions[1])
		if forward.x < 0 || forward.x >= len(g) || backward.x < 0 || backward.x >= len(g) {
			isXmas = false
			break
		}
		forwardRow := g[forward.x]
		backwardRow := g[backward.x]
		if forward.y < 0 || forward.y >= len(forwardRow) || backward.y < 0 || backward.y >= len(backwardRow) {
			isXmas = false
			break
		}
		if !(forwardRow[forward.y] == 'M' && backwardRow[backward.y] == 'S' || forwardRow[forward.y] == 'S' && backwardRow[backward.y] == 'M') {
			isXmas = false
			break
		}
	}
	return isXmas
}

type Point struct {
	x, y int
}

func (p Point) add(other Point) Point {
	return Point{
		x: p.x + other.x,
		y: p.y + other.y,
	}
}

func (p Point) times(a int) Point {
	return Point{
		x: p.x * a,
		y: p.y * a,
	}
}
