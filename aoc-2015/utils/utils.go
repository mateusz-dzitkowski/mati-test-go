package utils

import (
	"fmt"
	"os"
)

func LoadExercise(exercise string) string {
	data, err := os.ReadFile(fmt.Sprintf("inputs/%s", exercise))
	if err != nil {
		panic(err)
	}
	return string(data)
}
