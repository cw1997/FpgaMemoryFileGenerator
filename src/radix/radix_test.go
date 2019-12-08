// project : FpgaMemoryFileGenerator
// author : cw1997 [www.changwei.me]
// datetime : 2019-12-08 6:13
// repo : https://github.com/cw1997/FpgaMemoryFileGenerator

package radix

import "testing"

func TestConvertRadixStrToNum(t *testing.T) {
	var testCases = []struct {
		in       string // input
		expected int // expected result
	}{
		{"HEX", 16},
		{"DEC", 10},
		{"OCT", 8},
		{"BIN", 2},
	}

	for _, tt := range testCases {
		actual := ConvertRadixStrToNum(tt.in)
		if actual != tt.expected {
			t.Errorf("ConvertRadixStrToNum(%s) = %d; expected %d", tt.in, actual, tt.expected)
		}
	}
}

func TestConvertRadixStrToPlaceholder(t *testing.T) {
	var testCases = []struct {
		in       string // input
		expected string // expected result
	}{
		{"HEX", "x"},
		{"DEC", "d"},
		{"OCT", "o"},
		{"BIN", "b"},
	}

	for _, tt := range testCases {
		actual := ConvertRadixStrToPlaceholder(tt.in)
		if actual != tt.expected {
			t.Errorf("ConvertRadixStrToNum(%s) = %s; expected %s", tt.in, actual, tt.expected)
		}
	}
}
