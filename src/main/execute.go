package main

import (
	"fmt"
	"exercises"
)

func main() {
	var year int
	fmt.Print("Enter year: ")
	fmt.Scanf("%d", &year)
	if exercises.IsLeapYear(year) {
		fmt.Println(year , " is leap")
	} else {
		fmt.Println(year , " is not leap")
	}
}
