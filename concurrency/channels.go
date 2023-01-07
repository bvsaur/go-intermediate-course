package main

import "fmt"

func main() {
	// Unbufered channel
	// ubChannel := make(chan int)

	// Buffered channel
	bChannel := make(chan int, 1)

	// ubChannel <- 1		this will lead to an error

	bChannel <- 150

	fmt.Println(<-bChannel)
}
