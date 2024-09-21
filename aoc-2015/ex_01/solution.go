package ex_01

import (
	"aoc_2015/utils"
	"errors"
	"fmt"
)

func Solve() {
	data := utils.LoadExercise("01")
	fmt.Println(solve012(data))
}

func solve011(data string) int {
	total := 0
	for _, char := range data {
		switch char {
		case '(':
			total += 1
		case ')':
			total -= 1
		default:
			panic("wrong character detected")
		}
	}
	return total
}

func solve012(data string) (int, error) {
	total := 0
	for index, char := range data {
		switch char {
		case '(':
			total += 1
		case ')':
			total -= 1
			if total == -1 {
				return index + 1, nil
			}
		default:
			panic("wrong character detected")
		}
	}
	return 0, errors.New("didn't hit basement")
}
