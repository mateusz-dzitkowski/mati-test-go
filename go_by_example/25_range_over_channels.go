package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "two"
	queue <- "one"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
