package main

import "fmt"

func add_two(inp int16) int16 {
	return inp + 2
}

func print_string(inp string) {
	fmt.Println(inp)
}

func recursive_sum(sum int16, n int16) int16 {
	if n == 0 {
		return sum
	}
	return recursive_sum(sum+sum, n-1)
}

func sum_and_substraction(a int16, b int16) (sum int16, substraction int16) {
	sum = a + b
	substraction = a - b
	return
}

func main() {
	var x int16 = 2
	fmt.Println(x)
	fmt.Println(add_two(x))
	fmt.Println(recursive_sum(1, 10))
	fmt.Println(sum_and_substraction(5, 3))
}
