package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string, age int) *Person {
	return &Person{name: name, age: age}
}

func main() {
	mati := newPerson("mati", 12)
	fmt.Println(mati)

	dog := struct {
		name   string
		isGood bool
	}{"Rurzyn", true}
	fmt.Println(dog)
}
