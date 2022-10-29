package main

import (
	"fmt"
)

func greet(c chan string) {
	name := <-c //receiving value from channel
	fmt.Println("hello", name)
}

func main() {
	c := make(chan string)

	go greet(c)
	c <- "world"
	close(c)
}
