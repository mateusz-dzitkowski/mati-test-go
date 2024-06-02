package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	_ = os.Setenv("MY_BALLS", "my balls")
	fmt.Println(os.Getenv("MY_BALLS"))
	fmt.Println(os.Getenv("FOO"))

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
