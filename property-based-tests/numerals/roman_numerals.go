package numerals

import "strings"

type romanNumeral struct {
	value  int
	symbol string
}

type romanNumerals []romanNumeral

func (r *romanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)

	for _, s := range allRomanNumerals {
		if symbol == s.symbol {
			return s.value
		}
	}

	return 0
}

var allRomanNumerals = romanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.value {
			result.WriteString(numeral.symbol)
			arabic -= numeral.value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) int {
	var total int
	var i int
	for i = 0; i < len(roman)-1; i++ {
		symbol := roman[i]
		nextSymbol := roman[i+1]

		// Look ahead to see if can combine symbols
		potentialValue := allRomanNumerals.ValueOf(symbol, nextSymbol)
		if potentialValue != 0 {
			total += potentialValue
			i += 1
		} else {
			total += allRomanNumerals.ValueOf(symbol)
		}
	}

	// Handle the last symbol if there is one left
	if i < len(roman) {
		total += allRomanNumerals.ValueOf(roman[i])
	}

	return total
}
