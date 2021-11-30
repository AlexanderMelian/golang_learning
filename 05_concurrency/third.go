package main

import (
	"fmt"
	"time"
)

func goroutune() {
	time.Sleep(1 * time.Second)
	fmt.Println("This is a goroutune", time.Now().Second())
}

func main() {
	go goroutune()
	fmt.Println("Start", time.Now().Second())
	time.Sleep(3 * time.Second)
	fmt.Println("End", time.Now().Second())
}
