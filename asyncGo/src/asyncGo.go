package main

import (
	"fmt"
)

func example(s string) {
	fmt.Println(s)
}

func channelExample(s string) {
	//Print passed in message
	fmt.Println(s)

	//Create an output channel
	outputChannel := make(chan string)
	//Create a goroutine to pass strings into the output channel
	go func(s string) { outputChannel <- s }("I'm a string being sent to the channel!")
	//When the message is received capture it
	result := <- outputChannel
	//Then print it out
	fmt.Println(result, "...", "got printed from the channel!")
}

func channelTestExample(s string) {
	//Print passed in message
	fmt.Println(s)

	//Create an output channel
	outputChannel := make(chan string)

	//Anonymous go routine functions must be self-called
	go func() {
		outputChannel <- "Test Message One"
		outputChannel <- "Test Message Two"
	}()

	//Create a goroutine to pass strings into the output channel
	resultOne := <- outputChannel

	//The following go rutine will not send its message
	go func() { outputChannel <- "Test Message Three" }()
	resultTwo := <- outputChannel

	//Then print it out
	fmt.Println(resultOne, "...", "got printed from the channel!")

	//resultOne's assignment prevents resultTwo from ever being assigned to by outputChannel
	fmt.Println(resultTwo, "...", "got printed from the channel!")
}

func main() {
	example("I'm called directly!")

	go example("I'm a go routine!")

	channelExample("I'm a concurrent goroutine!")

	channelTestExample("I'm a channel test!")
}