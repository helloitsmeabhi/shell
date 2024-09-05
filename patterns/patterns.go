package main

import "fmt"

func main() {
	sine()
}
func test() {
	for i := 1; i < 10; i++ {
		fmt.Println("*")
	}
}
func sine() {
	var n int
	fmt.Println("Enter a number")
	fmt.Scanf("%d", &n)
	fmt.Print("  ")
	for i := 1; i <= n*3; i++ {
		fmt.Print(i)
		fmt.Print(" ")
	}
	fmt.Println()
	for i := 1; i <= (n+1)/2; i++ {
		fmt.Print(i)
		fmt.Print(" ")
		for j := 1; j <= i; j++ {
			fmt.Print(j)
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print(j)
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print(j)
			fmt.Print(" ")
		}
		fmt.Println()
	}
	for i := 1; i <= (n+1)/2; i++ {
		fmt.Print(i)
		fmt.Print(" ")
		for j := 1; j <= i; j++ {
			fmt.Print(j)
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print(j)
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print(j)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
