package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("tick at:", t.UTC())
			}
		}
	}()

	time.Sleep(time.Second * 2)
	ticker.Stop()
	done <- true
	fmt.Println("ticker stopped")
}
