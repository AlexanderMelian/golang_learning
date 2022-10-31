package main

import "fmt"

func main() {
	first()
	fmt.Println("--------")
	lifo()
}

func first() {
	defer fmt.Println("2")
	fmt.Println("1")
}

func lifo() {
	defer fmt.Println("3")
	defer fmt.Println("2")
	defer fmt.Println("1")
}
