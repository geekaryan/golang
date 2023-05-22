package main

import "fmt"

func main() {
	// var arr [5]int
	// arr[0] = 100
	// fmt.Println(arr)

	arr := [3]int{4, 5, 6}
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	fmt.Println(sum)
}
