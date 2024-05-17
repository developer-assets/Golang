package main

import (
	"fmt"
	"strings"
)

func main() {
	var temp, convertedTemp float64
	var fromUnit, toUnit string

	fmt.Print("Enter temprature: ")
	fmt.Scanln(&temp)

	fmt.Print("Enter current unit (C, F, K):")
	fmt.Scanln(&fromUnit)
	fromUnit = strings.ToUpper(fromUnit)

	fmt.Print("Enter unit to convert to (C, F, K):")
	fmt.Scanln(&toUnit)
	toUnit = strings.ToUpper(toUnit)

	switch fromUnit {
	case "C":
		if toUnit == "F" {
			convertedTemp = temp*9/5 + 32
		} else if toUnit == "K" {
			convertedTemp = temp + 273.15
		}

	case "F":
		if toUnit == "C" {
			convertedTemp = (temp - 32) * 5 / 9
		} else if toUnit == "K" {
			convertedTemp = (temp-32)*5/9 + 273.15
		}

	case "K":
		if toUnit == "C" {
			convertedTemp = temp - 273.15
		} else if toUnit == "F" {
			convertedTemp = (temp-273.15)*9/5 + 32
		}
	}

	fmt.Printf("Converted temprature: %.2f %s\n", convertedTemp, toUnit)

}
