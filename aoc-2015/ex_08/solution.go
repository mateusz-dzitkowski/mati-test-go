package ex_08

import (
	"aoc_2015/utils"
	"fmt"
	"strings"
)

func Solve() {
	data := utils.LoadExercise("08")
	fmt.Println(solveFirst(data))
	fmt.Println(solveSecond(data))
}

func solveFirst(data string) (total int) {
	for _, line := range strings.Split(data, "\n") {
		codeChars, memChars := countCharsFirst(line)
		total += codeChars - memChars
	}

	return total
}

func countCharsFirst(s string) (codeChars int, memChars int) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, 0
	}

	codeChars = len(s)
	for i := 0; i < codeChars; i++ {
		switch s[i] {
		case '\\':
			i++
			if s[i] == 'x' {
				i += 2
			}
			memChars++
		case '"':
		default:
			memChars++
		}
	}

	return codeChars, memChars
}

func solveSecond(data string) (total int) {
	for _, line := range strings.Split(data, "\n") {
		oldChars, newChars := countCharsSecond(line)
		total += newChars - oldChars
	}

	return total
}

func countCharsSecond(s string) (oldChars int, newChars int) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, 0
	}

	oldChars = len(s)
	newChars = oldChars
	for i := 0; i < oldChars; i++ {
		switch s[i] {
		case '\\':
			newChars++
		case '"':
			newChars++
		}
	}

	return oldChars, newChars + 2 // +2 for the beginning and ending quotation mark
}
