package main

import (
	"fmt"
	"time"
)

const (
	x string = "bruh"
	y int16  = 69
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println("1 + 1 =", 1+1)
	fmt.Println(1 + 1)
	fmt.Println(true)

	var a = "initial"
	fmt.Println(a)

	f := "apple"
	fmt.Println(f)

	fmt.Println(x, y)

	for n := range 20 {
		if n%3 == 0 {
			fmt.Println("foo")
		}
		if n%5 == 0 {
			fmt.Println("bar")
		} else {
			fmt.Println(n)
		}
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("good morning")
	default:
		fmt.Println("fuck you")
	}
}
