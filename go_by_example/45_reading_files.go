package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	fileName = "mati_test.txt"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dat, err := os.ReadFile(fileName)
	check(err)
	fmt.Println(string(dat))

	f, err := os.Open(fileName)
	defer func(f *os.File) { _ = f.Close() }(f)
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %v\n", n1, string(b1[:n1]))

	o2, err := f.Seek(int64(n1), io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	_, err = f.Seek(0, io.SeekStart)
	check(err)

	r4 := bufio.NewReader(f)
	b3, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b3))
}
