package numerals

import (
	"fmt"
	"testing"
)

var romanAndArabicTests = []struct {
	arabic uint16
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
	for _, tt := range romanAndArabicTests {
		t.Run(fmt.Sprintf("%d gets converted to %s", tt.arabic, tt.roman), func(t *testing.T) {
			got, _ := ConvertToRoman(tt.arabic)
			want := tt.roman

			if got != want {
				t.Errorf("ConvertToRoman(%d) = %q, want %q", tt.arabic, got, want)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, tt := range romanAndArabicTests {
		t.Run(fmt.Sprintf("%s gets converted to %d", tt.roman, tt.arabic), func(t *testing.T) {
			got := ConvertToArabic(tt.roman)
			want := tt.arabic

			if got != want {
				t.Errorf("ConvertToArabic(%q) = %d, want %d", tt.arabic, got, want)
			}
		})
	}
}

func TestConvertInvalidArabicToRoman(t *testing.T) {
	for i := 1; i <= 10; i++ {
		arabic := MaximumRomanNumeralValue + uint16(i*i)
		t.Run(fmt.Sprintf("%d convert roman numeral return error", arabic), func(t *testing.T) {
			_, err := ConvertToRoman(arabic)

			if err == nil {
				t.Errorf("%d can't be convert to roman but this didn't cause an error", arabic)
			}
		})
	}
}

func TestIsValidRomanNumeral(t *testing.T) {
	testcases := []struct {
		roman string
		want  bool
	}{
		{"I", true},
		{"IIII", false},
		{"IIIII", false},
		{"XXXX", false},
		{"VV", false},
		{"IIX", false},
		{"IIIX", false},
		{"VX", false},
		{"IXX", false},
		{"MMMCMXCIX", true},
		{"MMMM", false},
		{"XLX", false},
		{"CMD", false},
		{"CIXI", false},
	}

	for _, tt := range testcases {
		t.Run(tt.roman, func(t *testing.T) {
			got := IsValidRomanNumeral(tt.roman)

			if got != tt.want {
				t.Errorf("IsValidRomanNumeral(%q) = %t, want %t", tt.roman, got, tt.want)
			}
		})
	}
}

func FuzzConvertToRomanAndBackProperty(f *testing.F) {
	for _, tt := range romanAndArabicTests {
		f.Add(tt.arabic)
	}
	f.Add(MaximumRomanNumeralValue)
	f.Add(uint16(10000))

	f.Fuzz(func(t *testing.T, arabic uint16) {
		roman, err := ConvertToRoman(arabic)
		if err != nil {
			t.Skip("Invalid number to convert to roman")
		}

		fromRoman := ConvertToArabic(roman)

		if fromRoman != arabic {
			t.Errorf(
				"ConvertToRoman(%d) = %q, ConvertToArabic(%q) = %d, want %d",
				arabic,
				roman,
				roman,
				fromRoman,
				arabic,
			)
		}
	})
}
