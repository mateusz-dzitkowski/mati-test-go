package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return 42, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println(e)
		} else {
			fmt.Println(r)
		}
	}

	for i := range 5 {
		err := makeTea(i)
		if err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should by more tea")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark")
			} else {
				fmt.Println("Unknown error")
			}
			continue
		}
		fmt.Println("We got the tea")
	}
}
