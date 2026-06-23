package main

import (
	"strings"
)

func romanToArabic(roman string) int {
	roman = strings.ToUpper(roman)
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	for i := 0; i < len(roman); i++ {
		if i+1 < len(roman) && values[roman[i]] < values[roman[i+1]] {
			result -= values[roman[i]]
		} else {
			result += values[roman[i]]
		}
	}
	return result
}
