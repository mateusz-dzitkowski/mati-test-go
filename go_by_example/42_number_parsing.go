package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("21.37", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	k, _ := strconv.Atoi("321")
	fmt.Println(k)
}
