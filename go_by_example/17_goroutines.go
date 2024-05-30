package main

import (
	"fmt"
	"time"
)

func fo(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {
	fo("direct")

	go fo("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
