package main

import "fmt"

func main() {
	// break & continue
	
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Println("Value of i : ", i)
		i++
	}
	fmt.Println("----------------")

	for j:=0; j < 7; j++ {
		if j == 3 {
			continue
		} else if j == 5 {
			break
		}
		fmt.Println("Exit : ", j)
	}
	fmt.Println("Working...")
}