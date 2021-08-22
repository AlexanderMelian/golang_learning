package main

import "fmt"

type persona struct {
	name string
	edad int16
}

type worker struct {
	persona
	money float32
}

func main() {
	var persona_1 = persona{"Alex", 21}
	fmt.Println(persona_1)
	persona_1.name = "Pepe"
	persona_1.edad = 20
	fmt.Println(persona_1)
	var worker_1 = worker{persona{"Pepe", 20}, 100.0}
	fmt.Println(worker_1)
}
