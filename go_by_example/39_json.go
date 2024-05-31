package main

import (
	"encoding/json"
	"fmt"
)

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(5)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(3.14)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("mati-test")
	fmt.Println(string(strB))

	slcB, _ := json.Marshal([]string{"apple", "peach", "my balls"})
	fmt.Println(string(slcB))

	mapB, _ := json.Marshal(map[string]int{"apple": 5, "peach": 69})
	fmt.Println(string(mapB))

	res1B, _ := json.Marshal(Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "my balls"},
	})
	fmt.Println(string(res1B))

	res2B, _ := json.Marshal(Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "my balls"},
	})
	fmt.Println(string(res2B))

	byt := []byte(`{"num":3.14,"strs":["a","b"]}`)
	var dat map[string]any
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	fmt.Println(dat["num"].(float64))
	fmt.Println(dat["strs"].([]any)[0])
}
