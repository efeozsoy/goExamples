package main

import (
	"fmt"
	"time"
)

func main() {
	// Console Time & Date casting

	// Date() made a date data of our input
	t := time.Date(2021, time.April, 23, 20, 0, 0, 0, time.UTC)

	// t adıyla tari tipinde oluşturulan veriyi string tipinde ekrana yazar
	fmt.Printf("Efe launch at %s\n", t)

	fmt.Println("----------------")

	// time.Now() ile şu anın zaman bilgisini alıyoruz
	now := time.Now()

	// Elde edilen zaman bilgisini ekrana yazar
	fmt.Printf("The time now is %s\n", now)

	fmt.Println("-----------------")

	// İlk başta oluşturduğumuz t adındaki zaman bilgisinden ay, gün ve haftanın günü bilgilerini ekrana yazar
	fmt.Println("The month is ", t.Month())
	fmt.Println("The day is ", t.Day())
	fmt.Println("The weekday is ", t.Weekday())

	fmt.Println("------------------")

	// Tarihe 1 gün ekle
	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v, %v, %v\n",
		tomorrow.Weekday(),
		tomorrow.Month(),
		tomorrow.Day(),
		tomorrow.Year())

	fmt.Println("-------------------")

	longFormat := "Monday, Junuary 2, 2006"
	fmt.Println("Tomorrow date is ", tomorrow.Format(longFormat))

	fmt.Println("--------------------")

	shortFormat := "1/2/06"
	fmt.Println("Tomorrow is ", tomorrow.Format(shortFormat))

}
