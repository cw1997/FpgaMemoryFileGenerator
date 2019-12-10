// project : FpgaMemoryFileGenerator
// author : cw1997 [www.changwei.me]
// datetime : 2019-12-10 23:04:50
// repo : https://github.com/cw1997/FpgaMemoryFileGenerator

package bin

import (
	"bytes"
	"reflect"
	"testing"
)

type params struct {
}

type paramsNew struct {
}

func TestNewMifGenerator(t *testing.T) {
	var testCases = []struct {
		in       paramsNew // input
		expected Format    // expected result
	}{
		//{{depth: 8, width: 8, addressRadix: "hex", dataRadix: "hex", data: data1}, "16"},
		{paramsNew{},
			Format{}},
	}

	for _, tt := range testCases {
		actual := NewBinGenerator()
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("Generate(%v) = %v; expected %v", tt.in, actual, tt.expected)
		}
	}
}

func TestGenerate(t *testing.T) {
	data1 := []byte{byte(1), byte(2), byte(3), byte(4)}
	except1 := []byte{byte(1), byte(2), byte(3), byte(4)}

	var testCases = []struct {
		in       []byte // input
		expected []byte // expected result
	}{
		//{{depth: 8, width: 8, addressRadix: "hex", dataRadix: "hex", data: data1}, "16"},
		{data1, except1},
	}

	for _, tt := range testCases {
		f := NewBinGenerator()
		actual := f.Generate(tt.in)
		if !bytes.Equal(actual, tt.expected) {
			t.Errorf("Generate(%v) \n ========== actual start ========== \n%v\n  ========== actual end ========== \n"+
				"expected : \n ========== expected start ========== \n%v\n  ========== expected end ========== \n", tt.in, actual, tt.expected)
		}
	}
}
