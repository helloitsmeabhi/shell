package main

import "fmt"

func main() {
	//matrix()
	construct()
}
func test() {
	for i := 1; i < 10; i++ {
		fmt.Println("*")
	}
}
func matrix() {
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
func construct() {
	fmt.Print("Enter: ")
	var n int
	fmt.Scanf("%d", &n)
	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
