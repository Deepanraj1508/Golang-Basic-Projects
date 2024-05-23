package main

import (
	"fmt"
)

func Greet(customerName string) {
	fmt.Printf("%52s %s\n", "வ ண க்கம்", customerName)
	fmt.Printf("%72s\n", "_/\\_ Welcome to Mini Meals! _/\\_ ")
	fmt.Println()
}

func SayTata(customerName string) {
	fmt.Println()
	fmt.Printf("%17s", " ")
	fmt.Printf("_/\\_ Thank you %v for visiting Mini Meals! _/\\_\n", customerName)
	fmt.Printf("%55s\n", "We hope to see you again!")
	fmt.Printf("%51s\n", "Have a nice day! ")
}
