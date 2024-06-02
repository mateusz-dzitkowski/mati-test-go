package main

import _ "embed"

//go:embed mati_test.txt
var fileString string

//go:embed mati_test.txt
var fileByte []byte

func main() {
	print(fileString)
	print(string(fileByte))
}
