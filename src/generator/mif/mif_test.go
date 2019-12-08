// project : FpgaMemoryFileGenerator
// author : cw1997 [www.changwei.me]
// datetime : 2019-12-08 14:22
// repo : https://github.com/cw1997/FpgaMemoryFileGenerator

package mif

import "testing"

type params struct {
	depth int
	width int
	addressRadix string
	dataRadix string
	data []byte
}

type paramsNew struct {
	depth int
	width int
	addressRadix string
	dataRadix string
}

func TestNewMifGenerator(t *testing.T) {
	var testCases = []struct {
		in       paramsNew // input
		expected Format    // expected result
	} {
		//{{depth: 8, width: 8, addressRadix: "hex", dataRadix: "hex", data: data1}, "16"},
		{
			paramsNew{8, 8, "hex", "hex"},
			Format{8, 8, "HEX", "HEX"}},
	}


	for _, tt := range testCases {
		actual := NewMifGenerator(tt.in.depth, tt.in.width, tt.in.addressRadix, tt.in.dataRadix)
		if actual.depth != tt.expected.depth || actual.width != tt.expected.width || actual.addressRadix != tt.expected.addressRadix || actual.dataRadix != tt.expected.dataRadix  {
			t.Errorf("Generate(%v) = %v; expected %v", tt.in, actual, tt.expected)
		}
	}
}

func TestGenerate(t *testing.T) {
	data1 := []byte{byte(1), byte(2), byte(3), byte(4)}
	except1 := `DEPTH = 8;
WIDTH = 8;
ADDRESS_RADIX = HEX;
DATA_RADIX = HEX;
CONTENT
BEGIN
0:1
1:2
2:3
3:4
4:0
5:0
6:0
7:0
END;
`


	var testCases = []struct {
		in       params // input
		expected string // expected result
	} {
		//{{depth: 8, width: 8, addressRadix: "hex", dataRadix: "hex", data: data1}, "16"},
		{params{8, 8, "hex", "hex", data1}, except1},
	}


	for _, tt := range testCases {
		f := NewMifGenerator(tt.in.depth, tt.in.width, tt.in.addressRadix, tt.in.dataRadix)
		actual := f.Generate(tt.in.data)
		if actual != tt.expected {
			t.Errorf("Generate(%v) \n ========== actual start ========== \n%v\n  ========== actual end ========== \n" +
				"expected : \n ========== expected start ========== \n%v\n  ========== expected end ========== \n", tt.in, actual, tt.expected)
		}
	}
}

