package main

import (
	"fmt"
)

func main() {

	name := "sudo"
	fmt.Println("Before if")
	if name == "sudo" {
		fmt.Println("inside if")
	}
	fmt.Println("After if")

}
