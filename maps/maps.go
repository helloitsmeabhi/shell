package main

import "fmt"

func main() {
	//map1()
	//map2()
	//map3()
	map4()
}
func map1() {
	x := make(map[string]int)
	x["car"] = 1000
	fmt.Print(x["car"])
}
func map2() {
	y := make(map[int]int)
	y[1] = 1
	fmt.Print(y[1])
}
func map3() {
	x := make(map[string]string)
	x["H"] = "Hydrogen"
	x["He"] = "Helium"
	x["Li"] = "Lithium"
	x["Be"] = "Berillyium"
	for i, name := range x {
		fmt.Println(i, name)
	}
	if x["H"] == "Hydrogen" {
		fmt.Print("Hydrogen")
	}
}
func map4() {
	x := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"State": "Gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"State": "Gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"State": "Solid",
		},
	}
	for x, y := range x {
		for z, i := range y {
			fmt.Println(x, y, z, i)
		}
	}
}
