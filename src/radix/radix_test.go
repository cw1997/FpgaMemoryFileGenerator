// project : FpgaMemoryFileGenerator
// author : cw1997 [www.changwei.me]
// datetime : 2019-12-08 6:13
// repo : https://github.com/cw1997/FpgaMemoryFileGenerator

package radix

import "testing"

func TestCheckRadix(t *testing.T) {
	var testCases = []struct {
		in       int  // input
		expected bool // expected result
	}{
		{16, true},
		{10, true},
		{8, true},
		{2, true},
		{1, false},
		{3, false},
		{5, false},
		{-10, true},
	}

	for _, tt := range testCases {
		actual := CheckRadix(tt.in)
		if actual != tt.expected {
			t.Errorf("CheckRadix(%v) = %v; expected %v", tt.in, actual, tt.expected)
		}
	}
}

func TestConvertRadixStrToPlaceholder(t *testing.T) {
	var testCases = []struct {
		in       int    // input
		expected string // expected result
	}{
		{16, "x"},
		{-10, "d"},
		{10, "d"},
		{8, "o"},
		{2, "b"},
	}

	for _, tt := range testCases {
		actual := ConvertRadixNumToPlaceholder(tt.in)
		if actual != tt.expected {
			t.Errorf("ConvertRadixStrToNum(%v [%T]) = %v [%T]; expected %v [%T]\n", tt.in, tt.in, actual, actual, tt.expected, tt.expected)
		}
	}
}

func TestConvertDataRadixNumToStr(t *testing.T) {
	var testCases = []struct {
		in       int    // input
		expected string // expected result
	}{
		{16, "HEX"},
		{-10, "DEC"},
		{10, "UNS"},
		{8, "OCT"},
		{2, "BIN"},
	}

	for _, tt := range testCases {
		actual := ConvertDataRadixNumToStr(tt.in)
		if actual != tt.expected {
			t.Errorf("ConvertDataRadixNumToStr(%v [%T]) = %v [%T]; expected %v [%T]\n", tt.in, tt.in, actual, actual, tt.expected, tt.expected)
		}
	}
}

func TestConvertAddressRadixNumToStr(t *testing.T) {
	var testCases = []struct {
		in       int    // input
		expected string // expected result
	}{
		{16, "HEX"},
		{-10, "DEC"},
		{10, "DEC"},
		{8, "OCT"},
		{2, "BIN"},
	}

	for _, tt := range testCases {
		actual := ConvertAddressRadixNumToStr(tt.in)
		if actual != tt.expected {
			t.Errorf("ConvertDataRadixNumToStr(%v [%T]) = %v [%T]; expected %v [%T]\n", tt.in, tt.in, actual, actual, tt.expected, tt.expected)
		}
	}
}
