package main

import "fmt"

func main() {
	var i = 1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}
	forloop()
}
func forloop() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
