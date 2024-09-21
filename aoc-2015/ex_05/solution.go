package ex_05

import (
	"aoc_2015/utils"
	"fmt"
	"strings"
)

func Solve() {
	data := utils.LoadExercise("05")
	fmt.Println(solveFirst(data))
	fmt.Println(solveSecond(data))
}

func solveFirst(data string) int {
	var (
		vowels         = []string{"a", "e", "i", "o", "u"}
		naughtyStrings = []string{"ab", "cd", "pq", "xy"}
		repeats        = []string{"aa", "bb", "cc", "dd", "ee", "ff", "gg", "hh", "ii", "jj", "kk", "ll", "mm", "nn", "oo", "pp", "qq", "rr", "ss", "tt", "uu", "vv", "ww", "xx", "yy", "zz"}
	)

	total := 0

	for _, line := range strings.Split(data, "\n") {
		if countContains(line, vowels) >= 3 && countContains(line, repeats) >= 1 && countContains(line, naughtyStrings) < 1 {
			total += 1
		}
	}

	return total
}

func solveSecond(data string) int {
	total := 0

	alphabet := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	betweens := make([]string, len(alphabet)*len(alphabet))
	index := 0
	for _, let := range alphabet {
		for _, ter := range alphabet {
			betweens[index] = fmt.Sprintf("%s%s%s", let, ter, let)
			index += 1
		}
	}

	for _, line := range strings.Split(data, "\n") {
		if hasPair(line) && countContains(line, betweens) > 0 {
			total += 1
		}
	}

	return total
}

func countContains(s string, substrings []string) int {
	total := 0

	for _, substring := range substrings {
		total += strings.Count(s, substring)
	}

	return total
}

func hasPair(s string) bool {
	pairs := make([]string, len(s)-1)
	for i := range pairs {
		pairs[i] = s[i : i+2]
	}

	for i := 0; i < len(pairs)-2; i++ {
		for j := i + 2; j < len(pairs); j++ {
			if pairs[i] == pairs[j] {
				return true
			}
		}
	}

	return false
}
