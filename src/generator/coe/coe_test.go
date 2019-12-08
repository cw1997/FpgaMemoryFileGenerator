// project : FpgaMemoryFileGenerator
// author : cw1997 [www.changwei.me]
// datetime : 2019-12-08 16:02
// repo : https://github.com/cw1997/FpgaMemoryFileGenerator

package coe

import "testing"

type params struct {
	width int
	dataRadix int
	data []byte
}

type paramsNew struct {
	width int
	dataRadix int
}

func TestNewMifGenerator(t *testing.T) {
	var testCases = []struct {
		in       paramsNew // input
		expected Format    // expected result
	} {
		//{{depth: 8, width: 8, addressRadix: "hex", dataRadix: "hex", data: data1}, "16"},
		{paramsNew{8, 8},
			Format{8, 8}},
	}


	for _, tt := range testCases {
		actual := NewCoeGenerator( tt.in.width, tt.in.dataRadix)
		if actual.width != tt.expected.width || actual.dataRadix != tt.expected.dataRadix  {
			t.Errorf("Generate(%v) = %v; expected %v", tt.in, actual, tt.expected)
		}
	}
}

func TestGenerate(t *testing.T) {
	data1 := []byte{byte(1), byte(2), byte(3), byte(4)}
	except1 := `memory_initialization_radix=8;
memory_initialization_vector=
01,
02,
03,
04;
`

	var testCases = []struct {
		in       params // input
		expected string // expected result
	} {
		//{{depth: 8, width: 8, addressRadix: "hex", dataRadix: "hex", data: data1}, "16"},
		{params{8, 8, data1}, except1},
	}


	for _, tt := range testCases {
		f := NewCoeGenerator(tt.in.width, tt.in.dataRadix)
		actual := f.Generate(tt.in.data)
		if actual != tt.expected {
			t.Errorf("Generate(%v) \n ========== actual start ========== \n%v\n  ========== actual end ========== \n" +
				"expected : \n ========== expected start ========== \n%v\n  ========== expected end ========== \n", tt.in, actual, tt.expected)
		}
	}
}

