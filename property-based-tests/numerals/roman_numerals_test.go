package numerals

import (
	"fmt"
	"testing"
)

var tests = []struct {
	arabic int
	roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{25, "XXV"},
	{39, "XXXIX"},
	{40, "XL"},
	{45, "XLV"},
	{49, "XLIX"},
	{50, "L"},
	{90, "XC"},
	{100, "C"},
	{148, "CXLVIII"},
	{400, "CD"},
	{494, "CDXCIV"},
	{500, "D"},
	{666, "DCLXVI"},
	{900, "CM"},
	{998, "CMXCVIII"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
}

func TestConvertToRoman(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d gets converted to %s", tt.arabic, tt.roman), func(t *testing.T) {
			t.Parallel()

			if got := ConvertToRoman(tt.arabic); got != tt.roman {
				t.Errorf("ConvertToRoman(%d) = %q, want %q", tt.arabic, got, tt.roman)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s gets converted to %d", tt.roman, tt.arabic), func(t *testing.T) {
			got := ConvertToArabic(tt.roman)
			want := tt.arabic

			if got != want {
				t.Errorf("ConvertToArabic(%q) = %d, want %d", tt.arabic, got, want)
			}
		})
	}
}
