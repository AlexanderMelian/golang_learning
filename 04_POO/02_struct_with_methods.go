package main

import "fmt"

type Person struct {
	firstName string
	age       int
}

func (p *Person) speak(text string) {
	fmt.Println(text, p.firstName)
}

func main() {
	p1 := Person{
		firstName: "Ale",
		age:       32,
	}

	p1.speak("Hello, my name is")
}
