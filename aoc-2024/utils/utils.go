package utils

import (
	"fmt"
	"os"
)

func LoadExercise(exercise string) string {
	data, err := os.ReadFile(fmt.Sprintf("ex_%s/input", exercise))
	if err != nil {
		panic(err)
	}
	return string(data)
}
