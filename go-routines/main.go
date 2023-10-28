package main

import (
	"fmt"
	"time"
)

func getMessage(channel chan string, message string) {
	for i := 0; i < 5; i += 1 {
		fmt.Println(fmt.Sprintf("Hello from go routine! %d: %s", i, message))

		channel <- fmt.Sprintf("channel response (%s)", message)
		time.Sleep(time.Second)
	}
}

func main() {
	var messageCount = 5
	messageChan := make(chan string, messageCount)
	go getMessage(messageChan, "hey friends")

	fmt.Println("Channel got a value: ", <-messageChan)
	fmt.Println("Channel got a value: ", <-messageChan)
	fmt.Println("Channel got a value: ", <-messageChan)
	fmt.Println("Channel got a value: ", <-messageChan)
	fmt.Println("Channel got a value: ", <-messageChan)

	close(messageChan)
}
