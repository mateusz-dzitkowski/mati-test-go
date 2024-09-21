package main

import (
	"aoc_2015/utils"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
)

func main() {
	data := utils.LoadExercise("04")
	fmt.Println(solve042(data))
}

func solve041(data string) (int, error) {
	for number := range math.MaxInt {
		if number == 0 {
			continue
		}

		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", data, number)))
		hashHex := hex.EncodeToString(hash[:])

		if hashHex[:5] == "00000" {
			return number, nil
		}
	}
	return 0, errors.New("cant find the number")
}

func solve042(data string) (int, error) {
	for number := range math.MaxInt {
		if number == 0 {
			continue
		}

		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", data, number)))
		hashHex := hex.EncodeToString(hash[:])

		if hashHex[:6] == "000000" {
			return number, nil
		}
	}
	return 0, errors.New("cant find the number")
}
