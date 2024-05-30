package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func foo(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with this"}
	}

	return arg + 3, nil
}

func main() {
	_, err := foo(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae)
	} else {
		fmt.Println("Error doesn't match argError")
	}
}
