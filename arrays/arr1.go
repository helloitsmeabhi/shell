package main

import "fmt"

func main() {
	//array1()
	//sumavg()
	//usingrange()
	//usingshort()
	checkingvaluesofrange()
}
func array1() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
	fmt.Println(x[4])
}
func sumavg() {
	var x [5]int
	x[0] = 100
	x[1] = 200
	x[2] = 300
	x[3] = 400
	x[4] = 500
	var total float64
	var avg float64
	for i := 0; i < len(x); i++ {
		total += float64(x[i])
	}
	avg = total / float64(len(x))
	fmt.Print(x)
	fmt.Println()
	fmt.Print(total)
	fmt.Println()
	fmt.Print(avg)
}
func usingrange() {
	var total float64 = 0
	var x [5]int
	x[0] = 100
	x[1] = 200
	x[2] = 300
	x[3] = 400
	x[4] = 500
	for _, value := range x {
		total += float64(value)
	}
	fmt.Println(total / float64(len(x)))

}
func usingshort() {
	var total float64
	x := [5]float64{200, 300, 400, 500, 600}
	for _, value := range x {
		total += value
	}
	fmt.Println(total)
}
func checkingvaluesofrange() {
	var total float64
	x := [5]float64{300, 400, 300, 600, 500}
	for i, value := range x {
		fmt.Print(i)
		fmt.Print(" ")
		fmt.Println(value)
		total += value
	}
	fmt.Println(total)
}
