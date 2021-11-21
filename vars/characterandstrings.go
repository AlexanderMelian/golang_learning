package main

import (
	"fmt"
)

func main() {
	var b byte = 'a'
	fmt.Printf("%c \n", b) //this work because the byte can be converted to a character

	var c string = "abcdefghijklmnopqr"
	fmt.Printf("%s \n", c)

	fmt.Printf("%d \n", len(c))
	fmt.Printf("%s", c[0:5]) //this is a slice of the string

}
