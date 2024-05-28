package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["69"] = 69
	m["42"] = 42
	fmt.Println(m)

	fmt.Println(m["69"])

	delete(m, "69")
	fmt.Println(m)

	clear(m)
	fmt.Println(m)

	n := map[string]int{"foo": 1, "bar": 2}
	nn := map[string]int{"bar": 2, "foo": 1}

	fmt.Println(maps.Equal(n, nn))
}
