package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{num: 1},
		str:  "some name",
	}
	fmt.Println(co.num, co.str, co.base.num, co.describe())

	type describer interface {
		describe() string
	}

	func(d describer) { fmt.Println(d.describe()) }(co)
}
