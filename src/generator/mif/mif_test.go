// project : FpgaMemoryFileGenerator
// author : cw1997 [www.changwei.me]
// datetime : 2019-12-08 14:22
// repo : https://github.com/cw1997/FpgaMemoryFileGenerator

package mif

import "testing"

type params struct {
	depth        int
	width        int
	addressRadix int
	dataRadix    int
	data         []byte
}

type paramsNew struct {
	depth        int
	width        int
	addressRadix int
	dataRadix    int
}

func TestNewMifGenerator(t *testing.T) {
	var testCases = []struct {
		in       paramsNew // input
		expected Format    // expected result
	}{
		//{{depth: 8, width: 8, addressRadix: "hex", dataRadix: "hex", data: data1}, "16"},
		{paramsNew{8, 8, 16, 16},
			Format{8, 8, 16, 16}},
		{paramsNew{8, 8, 5, 16},
			Format{8, 8, 16, 16}},
	}

	for _, tt := range testCases {
		actual := NewMifGenerator(tt.in.depth, tt.in.width, tt.in.addressRadix, tt.in.dataRadix)
		if actual.depth != tt.expected.depth || actual.width != tt.expected.width || actual.addressRadix != tt.expected.addressRadix || actual.dataRadix != tt.expected.dataRadix {
			t.Errorf("NewMifGenerator(%v [%T]) = %v [%T]; expected %v [%T]\n", tt.in, tt.in, actual, actual, tt.expected, tt.expected)
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
	}{
		//{{depth: 8, width: 8, addressRadix: "hex", dataRadix: "hex", data: data1}, "16"},
		{params{8, 8, 16, 16, data1}, except1},
	}

	for _, tt := range testCases {
		f := NewMifGenerator(tt.in.depth, tt.in.width, tt.in.addressRadix, tt.in.dataRadix)
		actual := f.Generate(tt.in.data)
		if actual != tt.expected {
			t.Errorf("Generate(%v) \n ========== actual start ========== \n%v\n  ========== actual end ========== \n"+
				"expected : \n ========== expected start ========== \n%v\n  ========== expected end ========== \n", tt.in, actual, tt.expected)
		}
	}
}
