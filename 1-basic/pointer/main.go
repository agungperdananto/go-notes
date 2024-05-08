package main

import (
	"fmt"
	"strings"
)

var badWords = map[string]string{
	"dang":  "****",
	"shoot": "*****",
	"heck":  "****",
}

func removeProfanity(message *string) {

	// value <- pointer
	messageVal := *message
	for k, v := range badWords {
		messageVal = strings.ReplaceAll(messageVal, k, v)
	}
	// pointer <- value
	*message = messageVal
}

// don't touch below this line

func test(messages []string) {
	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}

func main() {
	messages1 := []string{
		"well shoot, this is awful",
		"dang robots",
		"dang them to heck",
	}

	messages2 := []string{
		"well shoot",
		"Allan is going straight to heck",
		"dang... that's a tough break",
	}

	test(messages1)
	test(messages2)
}
