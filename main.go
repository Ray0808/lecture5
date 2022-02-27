package main

import (
	"fmt"
	"log"
	"votes/calendar"
)

func main() {

	data := calendar.Date{}
	err := data.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}

	err = data.SetMonth(2)
	if err != nil {
		log.Fatal(err)
	}

	err = data.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data.Year())
	fmt.Println(data.Month())
	fmt.Println(data.Day())

}
