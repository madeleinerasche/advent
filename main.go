package main

import (
	ad "advent/advent23"
	"fmt"
)

func main() {
	fmt.Println("Advent 2023")

	// Day 1
	num, err := ad.Day1()
	if err != nil {
		fmt.Println(nil)
	} else {
		fmt.Println(num)
	}
}
