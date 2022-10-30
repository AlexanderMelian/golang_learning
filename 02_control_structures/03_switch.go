package main

import "fmt"

func main() {
	x := 1

	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("Nop")
	}

	var age int16 = 30

	switch {
	case age >= 18:
		fmt.Println("Adult")
	case age >= 0:
		fmt.Println("Not Adult")
	default:
		fmt.Println("Error")
	}
}
