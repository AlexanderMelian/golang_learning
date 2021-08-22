package main

import "fmt"

func main() {
	var string_test string = "MyString"
	fmt.Println(string_test)

	var float_test float64 = 3.14 // float64 is a 64-bit floating-point number
	fmt.Println(float_test)

	var bool_test bool = true
	fmt.Println(bool_test)

	var int_test int = 9223372036854775807 // int is a signed integer of 32 or 64 bits
	fmt.Println(int_test)

	var uint_test uint = 18446744073709551615 // uint is unsigned int
	fmt.Println(uint_test)

	var byte_test byte = 255 // byte is an alias for uint8
	fmt.Println(byte_test)

}
