package main

import "fmt"

func main() {
	//array1()
	//sumavg()
	//usingrange()
	//usingshort()
	//checkingvaluesofrange()
	//usingmake()
	slicing()
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
func usingmake() {
	x := [5]float64{200, 300, 400, 500, 600}
	y := make([]float64, 5, 10)
	z := x[1:4]
	var n int
	fmt.Println("Enter the array size: ")
	fmt.Scanf("%d", &n)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	k := make([]float64, n)
	fmt.Println(k)

}
func slicing() {
	slice1 := []int{100, 200, 300, 400}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
	sl1 := []int{100, 300, 400, 500}
	sl2 := []int{300, 700, 600, 300, 200}
	sl3 := append(sl1, sl2...)
	fmt.Println(sl3)
}
