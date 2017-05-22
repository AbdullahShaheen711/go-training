package main

import (
	"fmt"
	"excercises"
)

func main() {
	var year int
	fmt.Print("Enter year: ")
	fmt.Scanf("%d", &year)
	if excercises.IsLeapYear(year) {
		fmt.Println(year , " is leap")
	} else {
		fmt.Println(year , " is not leap")
	}
}
