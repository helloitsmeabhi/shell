package main

import "fmt"

/* data types
uint8/byte
uint16
uint32/rune
uint64
8,16,32,64-bits
***>machine dependent datatypes their size
	depends on the machine
unit
int
uintptr
->floating point
*complex64
*complex128
*float64
*/
func main() {

	addint()
	addfloat()
}
func addint() {
	fmt.Println("IInteger Addition: 1 + 1 = ", 1+1)
}
func addfloat() {
	fmt.Println(" Floating point addition: 1 + 1 = ", 1.0+1.0)
}
