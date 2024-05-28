package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println(a, len(a))

	b := [4]int{1, 2, 3, 4}
	fmt.Println(b)

	b = [...]int{9, 8, 7, 6}
	fmt.Println(b)

	var matrix [4][4]int
	for i := range 4 {
		for j := range 4 {
			matrix[i][j] = i * j
		}
	}
	fmt.Println(matrix)
}
