package main

import (
	"fmt"
	"time"
)

func goroutune() {
	fmt.Println("This is a goroutune", time.Now().Second())
}

func main() {
	fmt.Println("Start", time.Now().Second())
	go goroutune()
	time.Sleep(2 * time.Second)
	fmt.Println("End", time.Now().Second())
}
