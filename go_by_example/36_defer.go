package main

import (
	"fmt"
	"os"
)

func createFile(filename string) *os.File {
	fmt.Println("creating:", filename)
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing to", f.Name())
	_, err := fmt.Fprintf(f, "my balls\n")
	if err != nil {
		panic(err)
	}
}

func closeFile(f *os.File) {
	fmt.Println("closing:", f.Name())
	err := f.Close()
	if err != nil {
		panic(err)
	}
}

func main() {
	f := createFile("mati_test.txt")
	defer closeFile(f)
	writeFile(f)
}
