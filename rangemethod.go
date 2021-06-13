package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

	for i, v := range pow {
		fmt.Printf("2 ** %d = %d\n", i, v)
	}

	// array

	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item ", i, " is ", a[i])
	}

	// map
	capitals := map[string]string{"Turkey": "Ankara", "Greece": "Athens", "France": "Paris", "Japan": "Tokyo"}
	for key, val := range capitals {
		fmt.Println("Map item : Capital of ", key, " is ", val)
	}
}
