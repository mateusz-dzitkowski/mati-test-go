package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 1)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("Got message:", msg)
	default:
		fmt.Println("No message received")
	}

	message := "hi"
	select {
	case messages <- message:
		fmt.Println("Sent messages:", message)
	default:
		fmt.Println("No messages sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("Got message:", msg)
	case signal := <-signals:
		fmt.Println("Got signal:", signal)
	default:
		fmt.Println("No activity")
	}
}
