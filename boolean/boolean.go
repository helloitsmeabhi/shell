package main

import (
	"fmt"
)

func main() {
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(!true)
	problem()
}
func problem() {
	fmt.Println((true && false) || (false && true) || !(false && false))
}
