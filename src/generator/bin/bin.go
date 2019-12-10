// project : FpgaMemoryFileGenerator
// author : cw1997 [www.changwei.me]
// datetime : 2019/12/10 22:56
// repo : https://github.com/cw1997/FpgaMemoryFileGenerator

package bin

import (
)

type Format struct {
}

func NewBinGenerator() Format {
	return Format{}
}

func (m Format) Generate(data []byte) []byte {
	return data
}
