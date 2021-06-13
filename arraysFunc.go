package main

import "fmt"

func main() {
	/*
		myArray1 := [3]int{}
		myArray1[0] = 32
		myArray1[1] = 23
		myArray1[2] = 54

		fmt.Println(myArray1)
	*/

	// color array
	/*
		var colors [3]string
		colors[0] = "Red"
		colors[1] = "Blue"
		colors[2] = "Green"
		fmt.Println(colors)
		fmt.Println(colors[1])
		colors[1] = "Purple"
		fmt.Println(colors[1])
	*/

	var numbers = [5]int{4, 6, 7, 2, 29}
	fmt.Println(numbers)
	fmt.Println("Number of numbers: ", len(numbers))

	myArray2 := [...]int{4, 5, 6, 1, 2, 3}
	fmt.Println(myArray2)

	var cars [3]string
	cars[0] = "Tesla"
	cars[1] = "BMW"
	fmt.Println(len(cars))
	fmt.Println(cap(cars))

	/*
		i := 0
		for i <= len(cars)-1 {
			fmt.Println(cars[i])
			i++
		}
	*/

	for i := 0; i < len(cars); i++ {
		fmt.Println(cars[i])
	}

	for _, value := range cars {
		//fmt.Println("i = ", i, "& value = ", cars[i])
		fmt.Println(value)
	}
}
