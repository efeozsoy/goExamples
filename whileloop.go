package main

import (
	"fmt"
	"time"
)

func main() {
	// for ile while döngüsü kurmak. Normalde while diye bir döngümüz yok

	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
		time.Sleep(300 * time.Millisecond)
	}
}
