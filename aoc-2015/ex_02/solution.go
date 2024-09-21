package ex_02

import (
	"aoc_2015/utils"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	data := utils.LoadExercise("02")
	fmt.Println(solveFirst(data))
	fmt.Println(solveSecond(data))
}

func solveFirst(data string) int {
	total := 0

	for _, dimensions := range strings.Split(data, "\n") {
		lwh := strings.Split(dimensions, "x")

		l, _ := strconv.Atoi(lwh[0])
		w, _ := strconv.Atoi(lwh[1])
		h, _ := strconv.Atoi(lwh[2])

		top := l * w
		front := w * h
		side := h * l

		smallestSide := min(top, front, side)

		total += 2*top + 2*front + 2*side + smallestSide
	}

	return total
}

func solveSecond(data string) int {
	total := 0

	for _, dimensions := range strings.Split(data, "\n") {
		lwh := strings.Split(dimensions, "x")

		l, _ := strconv.Atoi(lwh[0])
		w, _ := strconv.Atoi(lwh[1])
		h, _ := strconv.Atoi(lwh[2])

		top := 2*l + 2*w
		front := 2*w + 2*h
		side := 2*h + 2*l

		smallestSide := min(top, front, side)
		volume := l * w * h

		total += smallestSide + volume
	}

	return total
}
