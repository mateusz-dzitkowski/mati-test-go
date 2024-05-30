package main

import "fmt"

func addToChannel(ch chan string, value string) {
	ch <- value
}

func main() {
	messages := make(chan string)

	go addToChannel(messages, "aaa")
	go addToChannel(messages, "bbb")
	go addToChannel(messages, "ccc")
	go addToChannel(messages, "ddd")
	go addToChannel(messages, "eee")
	go addToChannel(messages, "fff")

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
