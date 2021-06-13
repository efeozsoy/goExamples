package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("dosyam.txt")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("ERROR : ", err.Error())
		// loglandı
		os.Exit(1)
		// log package
		log.Fatal("ERROR : ", err)
	}
}