package main

import "fmt"

func main() {
	var x string = "hello world"
	fmt.Println(x)
	same()
	sumassign()
	stringcomp()
	autoassign()
}

// same variable assignment
func same() {
	var x string
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)
}

// sum assignment
func sumassign() {
	var x string
	x = "first"
	fmt.Println(x)
	x += " second"
	fmt.Println(x)
}

// string comparision
func stringcomp() {
	var x string = "hello"
	var y string = "hello"
	fmt.Println(x == y)
}

// auto type assign
func autoassign() {
	x := "hello world"
	fmt.Println(x)
}
