package ex_03

import (
	"aoc_2015/utils"
	"fmt"
)

type Position struct {
	x int
	y int
}

func Solve() {
	data := utils.LoadExercise("03")
	fmt.Println(solve032(data))
}

func solve031(data string) int {
	visited := make(map[Position]struct{})

	currentPosition := Position{}
	visited[currentPosition] = struct{}{}
	totalVisited := 1

	for _, char := range data {
		switch char {
		case '^':
			currentPosition.y += 1
		case '>':
			currentPosition.x += 1
		case 'v':
			currentPosition.y -= 1
		case '<':
			currentPosition.x -= 1
		}

		_, ok := visited[currentPosition]
		if !ok {
			totalVisited += 1
			visited[currentPosition] = struct{}{}
		}
	}

	return totalVisited
}

func solve032(data string) int {
	visited := make(map[Position]struct{})

	santaPosition := Position{}
	robotPosition := Position{}
	visited[santaPosition] = struct{}{}
	totalVisited := 1

	for index, char := range data {
		var position *Position
		if index%2 == 0 {
			position = &santaPosition
		} else {
			position = &robotPosition
		}

		switch char {
		case '^':
			position.y += 1
		case '>':
			position.x += 1
		case 'v':
			position.y -= 1
		case '<':
			position.x -= 1
		}

		_, ok := visited[*position]
		if !ok {
			totalVisited += 1
			visited[*position] = struct{}{}
		}
	}

	return totalVisited
}
