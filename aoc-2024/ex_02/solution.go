package ex_02

import (
	"fmt"
	"main/utils"
	"math"
	"strconv"
	"strings"
)

func Solve() {
	data := utils.LoadExercise("02")
	fmt.Println(solveFirst(data))
	fmt.Println(solveSecond(data))
}

func solveFirst(data string) int {
	output := 0
	for _, report := range readNumbers(data) {
		if isSafe(report) {
			output += 1
		}
	}
	return output
}

func solveSecond(data string) int {
	output := 0
	for _, report := range readNumbers(data) {
		if isSafeWithDampener(report) {
			output += 1
		}
	}
	return output
}

func readNumbers(data string) [][]int {
	var output [][]int

	for _, line := range strings.Split(data, "\n") {
		numbersStrings := strings.Split(line, " ")
		var numbers []int
		for _, s := range numbersStrings {
			number, _ := strconv.Atoi(s)
			numbers = append(numbers, number)
		}

		output = append(output, numbers)
	}

	return output
}

func isSafe(report []int) bool {
	var increasing bool
	for i := range report {
		if i == 0 {
			continue
		}
		diff := report[i] - report[i-1]
		currentlyIncreasing := diff > 0
		if i == 1 {
			increasing = currentlyIncreasing
		} else {
			if increasing != currentlyIncreasing {
				return false
			}
		}
		absDiff := int(math.Abs(float64(diff)))
		if absDiff < 1 || absDiff > 3 {
			return false
		}
	}
	return true
}

func isSafeWithDampener(report []int) bool {
	for indexToExclude := range report {
		dampenedReport := make([]int, len(report)-1)
		originalReportTracker := 0
		for i := range len(report) - 1 {
			if i == indexToExclude {
				originalReportTracker += 1
			}
			dampenedReport[i] = report[originalReportTracker]
			originalReportTracker += 1
		}
		if isSafe(dampenedReport) {
			return true
		}
	}
	return false
}
