package main

import "fmt"

func main() {

	a := 1

	if a == 1 {
		fmt.Println("a is 1")
	} else if a == 2 {
		fmt.Println("a is 2")
	} else {
		fmt.Println("a is not 1 or 2")
	}

	a = 2

	if a == 1 {
		fmt.Println("a is 1")
	} else if a == 2 {
		fmt.Println("a is 2")
	} else {
		fmt.Println("a is not 1 or 2")
	}

	a = 3

	if a == 1 {
		fmt.Println("a is 1")
	} else if a == 2 {
		fmt.Println("a is 2")
	} else {
		fmt.Println("a is not 1 or 2")
	}

}
