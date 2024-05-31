package main

import (
	"cmp"
	"fmt"
	"slices"
)

func lenCmp(a, b string) int {
	return cmp.Compare(len(a), len(b))
}

type Person34 struct {
	name string
	age  int
}

func main() {
	strs1 := []string{"c", "a", "b"}
	strs2 := []string{"c", "a", "b"}
	slices.Sort(strs1)
	fmt.Println(strs1)
	slices.SortFunc(strs2, lenCmp)
	fmt.Println(strs2)

	ints := []int{4, 2, 3, 1}
	slices.Sort(ints)
	fmt.Println(ints)

	people := []Person34{
		{
			name: "Mati",
			age:  32,
		},
		{
			name: "Andrzej",
			age:  12,
		},
		{
			name: "Łukasz",
			age:  45,
		},
	}
	slices.SortFunc(people, func(a, b Person34) int { return cmp.Compare(a.age, b.age) })
	fmt.Println(people)
}
