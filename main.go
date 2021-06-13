package main

import "fmt"

func main() {

	/*
		var message string
		message = "Hello Go"

		var message string = "Hello!"

		var message = "Hello Go!"

		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)

		var message = "Hello World!"
		var a, b, c = 3, true, 5

		fmt.Println(message, a, b, c)

		u := 55
		v, n := "abc", true

		Eğer tipinden eminsen tipini mutlaka yaz

		a := "String"
		b := 'M' char
		c := `String`

		var myFloat32 float32 = 44.
		myComplex := complex(3,4)

		println(myFloat32)
		println(myComplex)

		const explanation = "Go Playground"
		const pi = 3.14

		a := 10
		b := 20

		total := a + b
		println("total: ", total)

		total = total - 5
		println("total: ", total)

		total *= 20
		println("total: ", total)

		total = a / b
		println("total: ", total)

		total = a % b
		println("total: ", total)

		total++
		println("total: ", total)

		total--
		println("total: ", total)

	*/

	// var myString = "1"
	// var x = 10
	// var f float32 = 2.0

	// fmt.Println(myString, x, f)

	// string'den int'e dönüştürme

	// number, _ := strconv.Atoi(myString)
	// fmt.Println(number)
	// result := number + 2
	// fmt.Println(result)

	var i int = 55
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Println(i, f, u)
}
