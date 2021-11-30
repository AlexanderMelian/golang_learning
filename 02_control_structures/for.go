package main

import "fmt"

func main() {

	for a := 1; a < 10; a++ {
		fmt.Println(a)
	}

	i := 1
	for i <= 10 { //this is an example of a while loop
		fmt.Println(i)
		i++
	}

	j := 1
	for { //this is an example of an infinite loop
		if j > 10 {
			break
		}
		fmt.Println(j)
		j++
	}

}
