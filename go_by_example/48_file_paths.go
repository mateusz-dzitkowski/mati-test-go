package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println(p)

	fmt.Println(filepath.Join("dir///", "filename"))
	fmt.Println(filepath.Join("dir/../dir", "filename"))

	fmt.Println(filepath.Dir(p), filepath.Base(p))

	filename := "config.json"
	fmt.Println(filepath.Ext(filename))

	rel, err := filepath.Rel("a/b/", "a/b/t/filename")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
