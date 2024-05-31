package main

import (
	"fmt"
	"sync"
	"time"
)

func worker29(id int) {
	fmt.Println("worker", id, "starting")
	time.Sleep(time.Second)
	fmt.Println("worker", id, "done")
}

func main() {
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker29(i)
		}()
	}

	wg.Wait()
}
