// project : FpgaMemoryFileGenerator
// author : cw1997 [www.changwei.me]
// datetime : 2019/12/8 20:27
// repo : https://github.com/cw1997/FpgaMemoryFileGenerator

package ascii

import (
	"bytes"
	"testing"
)

func TestConvert(t *testing.T) {
	var testCases = []struct {
		in       string // input
		expected []byte // expected result
	}{
		{"abcd", []byte{97, 98, 99, 100}},
		{"ABCD", []byte{65, 66, 67, 68}},
		{"0123", []byte{48, 49, 50, 51}},
		{"+-*/", []byte{43, 45, 42, 47}},
	}

	for _, tt := range testCases {
		actual := Converter(tt.in)
		if !bytes.Equal(actual, tt.expected) {
			t.Errorf("Ascii.Converter(%v [%T]) = %v [%T]; expected %v [%T]\n", tt.in, tt.in, actual, actual, tt.expected, tt.expected)
		}
	}
}
