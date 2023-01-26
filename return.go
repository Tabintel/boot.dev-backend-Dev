package main

import (
	"fmt"
)

func main() {
	age := 12
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(age)
	fmt.Println("you are", age, "years old")
	fmt.Println("you are an adult in", yearsUntilAdult, "years")
	fmt.Println("you can drink in", yearsUntilDrinking, "years")
	fmt.Println("you can rent a car in", yearsUntilCarRental, "years")
}

func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	return
}