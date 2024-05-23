package hello

import (
	"fmt"
	"time"
)

func Datefunc() {
	// Get the current time
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)
	fmt.Println("")
	fmt.Println("----------------------------")
	// Formatting the current time
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted Time:", formattedTime)
	fmt.Println("")
	fmt.Println("----------------------------")
	// Parsing a time string
	timeString := "2024-05-20 12:34:56"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeString)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	} else {
		fmt.Println("Parsed Time:", parsedTime)
	}
	fmt.Println("")
	fmt.Println("----------------------------")
	// Getting components of the date
	year, month, day := currentTime.Date()
	fmt.Printf("Year: %d, Month: %d, Day: %d\n", year, month, day)
	fmt.Println("")
	fmt.Println("----------------------------")
	// Getting the weekday
	weekday := currentTime.Weekday()
	fmt.Println("Today is:", weekday)
	fmt.Println("")
	fmt.Println("----------------------------")
}
