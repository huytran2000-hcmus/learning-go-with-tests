package numerals

import (
	"fmt"
	"strings"
)

type romanNumeral struct {
	value  uint16
	symbol string
}

const MaximumRomanNumeralVal uint16 = 3999

type romanNumerals []romanNumeral

func (r *romanNumerals) ValueOf(symbols ...byte) uint16 {
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

func ConvertToRoman(arabic uint16) (string, error) {
	if arabic > 3999 {
		return "", fmt.Errorf("%d is above the roman numeral maximum value", arabic)
	}
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.value {
			result.WriteString(numeral.symbol)
			arabic -= numeral.value
		}
	}

	return result.String(), nil
}

func ConvertToArabic(roman string) uint16 {
	var total uint16
	var i int
	for i = 0; i < len(roman)-1; i++ {
		symbol := roman[i]
		nextSymbol := roman[i+1]

		// Look ahead to see if it is possible to combine symbols
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
