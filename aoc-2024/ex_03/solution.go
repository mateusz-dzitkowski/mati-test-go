package ex_03

import (
	"fmt"
	"main/utils"
	"strconv"
	"strings"
)

func Solve() {
	data := utils.LoadExercise("03")
	fmt.Println(solveFirst(data))
	fmt.Println(solveSecond(data))
}

func solveFirst(data string) int {
	return addUpMuls(data)
}

func solveSecond(data string) int {
	output := 0
	donts := strings.Split(data, "don't()")
	output += addUpMuls(donts[0]) // before the first don't()
	for _, dont := range donts[1:] {
		dos := strings.Split(dont, "do()")
		for _, do := range dos[1:] {
			output += addUpMuls(do)
		}
	}
	return output
}

func addUpMuls(data string) int {
	output := 0
	for _, restOfMul := range strings.Split(data, "mul(") {
		insidePar := strings.Split(restOfMul, ")")[0]
		numberStrings := strings.Split(insidePar, ",")
		if len(numberStrings) != 2 {
			continue
		}
		left, err := strconv.Atoi(numberStrings[0])
		if err != nil {
			continue
		}
		right, err := strconv.Atoi(numberStrings[1])
		if err != nil {
			continue
		}
		output += left * right
	}
	return output
}
