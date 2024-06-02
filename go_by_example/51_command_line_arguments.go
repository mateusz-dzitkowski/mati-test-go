package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	numberPtr := flag.Int("number", 42, "a number")
	forkPtr := flag.Bool("fork", false, "a fork")
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()
	fmt.Println(*wordPtr, *numberPtr, *forkPtr, svar)
}
