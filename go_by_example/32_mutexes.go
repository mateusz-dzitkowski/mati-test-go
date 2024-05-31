package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println("Incrementing", name)
	c.counters[name]++
}

func doIncrement(wg *sync.WaitGroup, c *Container, name string, n int) {
	for range n {
		c.inc(name)
	}
	wg.Done()
}

func main() {
	c := Container{counters: map[string]int{"a": 0, "b": 0}}

	var wg sync.WaitGroup
	wg.Add(5)
	go doIncrement(&wg, &c, "a", 10)
	go doIncrement(&wg, &c, "b", 10)
	go doIncrement(&wg, &c, "a", 10)
	go doIncrement(&wg, &c, "b", 10)
	go doIncrement(&wg, &c, "a", 10)
	wg.Wait()
	fmt.Println(c.counters)
}
