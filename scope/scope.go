package main

import "fmt"

var x = "hello"

func main() {
	fmt.Println(x)
	scopevarchange()
}
func scopevarchange() {
	x = "world"
	fmt.Println(x)
}
