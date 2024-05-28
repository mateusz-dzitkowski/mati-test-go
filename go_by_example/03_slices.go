package main

import (
	"fmt"
)

func main() {
	var s []string
	fmt.Println(s, s == nil, len(s))

	s = make([]string, 3)
	fmt.Println(s, s == nil, len(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "cde"
	fmt.Println(s)

	s = append(s, "xd")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println(s[2:])

}
