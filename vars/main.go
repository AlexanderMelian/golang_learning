package main

import "fmt"

func main() {

	var string_test string = "MyString"
	fmt.Println(string_test)

	var float_test float64 = 3.14 // float64 is a 64-bit floating number
	fmt.Println(float_test)
	//float32, float64

	var bool_test bool = true // simple boolean
	fmt.Println(bool_test)

	var int_test int = 9223372036854775807 // int is a signed integer of 32 or 64 bits, depends of the cpu architecture
	fmt.Println(int_test)
	//int, int8, int16, int32, int64

	var uint_test uint = 18446744073709551615 // uint is unsigned int
	fmt.Println(uint_test)
	//uint, uint8, uint16, uint32, uint64

	var byte_test byte = 255 // byte is an alias for uint8, or integer of 8 bits
	fmt.Println(byte_test)

	var rune_test rune = 12849823 // rune is an alias for int32, or integer of 32 bits
	fmt.Println(rune_test)

	var uintptr_test uintptr = 182941239 // uintptr is an alias for uint64, or integer of 64 bits
	fmt.Println(uintptr_test)



}
