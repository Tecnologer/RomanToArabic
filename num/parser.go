package num

import (
	"regexp"
	"strings"
)

//IsRoman check if a string is a roman number
func IsRoman(roman string) bool {
	var regex = "(?i)^M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$"
	match, _ := regexp.MatchString(regex, roman)
	return match
}

//ParseRoman Parse arabic number to roman number
func ParseRoman(roman string) int {
	roman = strings.ToUpper(roman)
	var romanValues = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var finalNumber int
	for i := 0; i < len(roman); i++ {
		letter := string(roman[i])

		var nextLetter string
		if (i + 1) < len(roman) {
			nextLetter = string(roman[i+1])
		}

		if nextLetter != "" {
			if romanValues[letter] >= romanValues[nextLetter] {
				finalNumber += romanValues[letter]
			} else {
				finalNumber -= romanValues[letter]
			}
		} else {
			finalNumber += romanValues[letter]
		}
	}

	return finalNumber
}

//ParseArabic Parse Roman number to Arabic number
func ParseArabic(arabic int) string {
	var roman string

	var romanValues = map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}

	var order = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for _, key := range order {
		var value = romanValues[key]
		var qty = arabic / value
		if qty > 3 {
			qty = 3
		}

		for i := 0; i < qty; i++ {
			roman += key
		}
		arabic = arabic - (qty * value)
	}
	return roman
}
