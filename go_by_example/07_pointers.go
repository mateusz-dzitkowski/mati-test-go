package main

import "fmt"

func zeroVal(n int) {
	n = 0
}

func zeroPtr(n *int) {
	*n = 0
}

func main() {
	n := 1
	fmt.Println(n)
	zeroVal(n)
	fmt.Println(n)
	zeroPtr(&n)
	fmt.Println(n)
}
