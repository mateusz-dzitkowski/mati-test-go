package ex_01

import (
	"fmt"
	"main/utils"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Solve() {
	data := utils.LoadExercise("01")
	fmt.Println(solveFirst(data))
	fmt.Println(solveSecond(data))
}

func solveFirst(data string) int {
	output := 0
	left, right := readNumbers(data)
	slices.Sort(left)
	slices.Sort(right)
	for i := range left {
		output += int(math.Abs(float64(left[i] - right[i])))
	}
	return output
}

func solveSecond(data string) int {
	output := 0
	left, right := readNumbers(data)
	scorer := createSimilarityScorer(right)
	for _, number := range left {
		output += scorer.getScore(number)
	}
	return output
}

func readNumbers(data string) ([]int, []int) {
	var left []int
	var right []int

	for _, line := range strings.Split(data, "\n") {
		numberStrings := strings.Split(line, "   ")
		leftNumber, _ := strconv.Atoi(numberStrings[0])
		rightNumber, _ := strconv.Atoi(numberStrings[1])
		left = append(left, leftNumber)
		right = append(right, rightNumber)
	}

	return left, right
}

type SimilarityScorer map[int]int

func createSimilarityScorer(list []int) SimilarityScorer {
	scorer := make(SimilarityScorer)
	for _, number := range list {
		_, found := scorer[number]
		if found {
			scorer[number] += 1
		} else {
			scorer[number] = 1
		}
	}
	return scorer
}

func (s SimilarityScorer) getScore(number int) int {
	value, found := s[number]
	if found {
		return value * number
	} else {
		return 0
	}
}
