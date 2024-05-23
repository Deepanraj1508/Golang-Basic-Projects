package main

import (
	"fmt"
	"strings"
)

type Menu struct {
	itemNo    uint
	itemName  string
	itemPrice float64
}

var menu = []Menu{
	//itemNo  itemName   itemPrice
	{1, "Dhosa", 20},
	{2, "Italy", 25},
	{3, "Poori", 30},
	{4, "Briyani", 100},
	{5, "Chicken Rice", 100},
	{6, "Pongal", 30},
	{7, "Mutton Briyani", 300},
	{8, "Parota", 20},
	{9, "Veg Rice", 70},
	{10, "Chicken 65(100gram)", 50},
	{11, "Chicken Grill", 250.00},
	{12, "Samosa", 10},
	{13, "Ice Cream", 25},
}
func printMenu() {
	fmt.Printf("%15s\n", "Menu")
	fmt.Printf("%s\n", strings.Repeat("-", 35))
	fmt.Printf("%-7s %6s    %12s\n", "S.No.", "Price", "Item Name")
	fmt.Printf("%s\n", strings.Repeat("-", 35))

	for _, element := range menu {
		fmt.Printf(" %-7d %.2f    %-4s\n", element.itemNo, element.itemPrice, element.itemName)
	}
	fmt.Printf("%s", strings.Repeat("-", 35))
	fmt.Println()
}
