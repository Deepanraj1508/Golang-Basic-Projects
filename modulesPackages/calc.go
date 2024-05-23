package main

import "fmt"

func PrintCalc(){
	fmt.Println("Enter two numbers to calculate:")
	var num1, num2 int
	fmt.Scanf("%d %d", &num1, &num2)
	fmt.Println("The sum of", num1, "and", num2, "is:", num1+num2)
}