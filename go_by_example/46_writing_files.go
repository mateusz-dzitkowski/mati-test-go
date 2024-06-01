package main

import (
	"bufio"
	"fmt"
	"os"
)

func check46(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("data_1", d1, 0644)
	check46(err)

	f, err := os.Create("data_2")
	check46(err)
	defer func(*os.File) {
		check46(f.Close())
	}(f)

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check46(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check46(err)
	fmt.Printf("wrote %d bytes\n", n3)

	check46(f.Sync())

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check46(err)
	fmt.Printf("wrote %d bytes\n", n4)

	check46(w.Flush())
}
