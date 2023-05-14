package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Sudoapt")
	fmt.Println(message)
}
