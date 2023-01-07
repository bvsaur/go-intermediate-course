package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	/*
		Adding a buffer to the channel gives us the behabiour of a traffic light,
		since we are letting just 2 go routines to run at the time.
		Once there's an empty slot on the channel, the program continues with its execution
	*/
	c := make(chan int, 2)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomething(i, &wg, c)
	}

	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("ID %d started\n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("ID %d finished\n", i)
	<-c
}
