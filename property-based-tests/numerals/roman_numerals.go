package numerals

import (
	"fmt"
	"strings"
)

const MaximumRomanNumeralValue uint16 = 3999

var (
	repeatableSymbols   = []string{"I", "X", "C", "M"}
	UnrepeatableSymbols = []string{
		"IV",
		"V",
		"IX",
		"XL",
		"L",
		"XC",
		"CD",
		"D",
		"CM",
	}
)

const maxRomanSymbolValue uint16 = 1000

type romanNumeral struct {
	value  uint16
	symbol string
}

type romanNumerals []romanNumeral

var descendingAllRomanNumerals = romanNumerals{
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

	for _, numeral := range descendingAllRomanNumerals {
		for arabic >= numeral.value {
			result.WriteString(numeral.symbol)
			arabic -= numeral.value
		}
	}

	return result.String(), nil
}

func ConvertToArabic(roman string) uint16 {
	var total uint16
	symbols := parseRomanSymbols(roman)
	for _, symbol := range symbols {
		total += valueOfRomanSymbol(symbol)
	}

	return total
}

func IsValidRomanNumeral(roman string) bool {
	if !hasValidRepeatableSymbols(roman) {
		return false
	}

	symbols := parseRomanSymbols(roman)
	if !isRomanSymbolValuesAscending(symbols) {
		return false
	}

	if !isSymbolsCombineCorrectly(symbols) {
		return false
	}

	return true
}

func hasValidRepeatableSymbols(roman string) bool {
	if hasTooManyRepeatableSymbols(roman) {
		return false
	}

	if hasNonRepeatableSymbols(roman) {
		return false
	}

	return true
}

func isRomanSymbolValuesAscending(symbols []string) bool {
	for i := 0; i < len(symbols)-1; i++ {
		symbol := symbols[i]
		nextSymbol := symbols[i+1]
		val := valueOfRomanSymbol(symbol)
		nextVal := valueOfRomanSymbol(nextSymbol)

		if val < nextVal {
			return false
		}
	}
	return true
}

func isSymbolsCombineCorrectly(symbols []string) bool {
	for i, symbol := range symbols {
		if i+1 == len(symbols) {
			break
		}

		nextSymbol := symbols[i+1]
		if !isValidCombination(symbol, nextSymbol) {
			return false
		}
	}
	return true
}

func valueOfRomanSymbol(symbol string) uint16 {
	for _, s := range descendingAllRomanNumerals {
		if symbol == s.symbol {
			return s.value
		}
	}

	return 0
}

func parseRomanSymbols(roman string) []string {
	var symbols []string
	for i := 0; i < len(roman); i++ {
		character := string(roman[i])
		nexCharacter := ""
		atTheEndOfLoop := i+1 == len(roman)
		if !atTheEndOfLoop {
			nexCharacter = string(roman[i+1])
		}

		// Look ahead to see if it is possible to combine symbols
		potentialSymbol := character + nexCharacter
		if isValidRomanSymbol(potentialSymbol) {
			symbols = append(symbols, potentialSymbol)
			i++
		} else {
			symbols = append(symbols, character)
		}
	}

	return symbols
}

func hasTooManyRepeatableSymbols(roman string) bool {
	for _, symbol := range repeatableSymbols {
		invalidRepeatedSymbols := strings.Repeat(symbol, 4)
		if strings.Contains(roman, invalidRepeatedSymbols) {
			return true
		}
	}

	return false
}

func hasNonRepeatableSymbols(roman string) bool {
	for _, symbol := range UnrepeatableSymbols {
		invalidSymbols := symbol + symbol
		if strings.Contains(roman, invalidSymbols) {
			return true
		}
	}

	return false
}

func isValidRomanSymbol(symbol string) bool {
	for _, s := range descendingAllRomanNumerals {
		if symbol == s.symbol {
			return true
		}
	}

	return false
}

func isValidCombination(symbol string, nextSymbol string) bool {
	if symbol == nextSymbol {
		if isRepeatableSymbol(symbol) {
			return true
		} else {
			return false
		}
	}

	if isSymbolsHasDifferentRepresentation(symbol, nextSymbol) {
		return false
	}

	return true
}

func isRepeatableSymbol(symbol string) bool {
	for _, s := range repeatableSymbols {
		if s == symbol {
			return true
		}
	}

	return false
}

func isSymbolsHasDifferentRepresentation(symbol string, nextSymbol string) bool {
	symValue := valueOfRomanSymbol(symbol)
	nexSymValue := valueOfRomanSymbol(nextSymbol)
	nextBiggerValue := nextBiggerSymbolValue(symbol)

	if nextBiggerValue <= symValue+nexSymValue {
		return true
	}
	return false
}

func nextBiggerSymbolValue(symbol string) uint16 {
	for i := 0; i < len(descendingAllRomanNumerals)-1; i++ {
		nextSymbol := descendingAllRomanNumerals[i+1].symbol
		if nextSymbol == symbol {
			currentSymbol := descendingAllRomanNumerals[i].value
			return currentSymbol
		}
	}

	return descendingAllRomanNumerals[0].value * 10
}
