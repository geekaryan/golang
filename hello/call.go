package main

import (
	"fmt"

	"examples/../greetings.go"
)

func main() {
	message := greetings.Hello("sudoaptrana")
	fmt.Println(message)
}
