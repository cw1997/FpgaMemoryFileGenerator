// project : FpgaMemoryFileGenerator
// author : cw1997 [www.changwei.me]
// datetime : 2019-12-08 15:50
// repo : https://github.com/cw1997/FpgaMemoryFileGenerator

package generator

type Format int

const (
	FORMAT_MIF format = iota
	FORMAT_COE
)

type Generator interface {
	Generate(data []byte) string
}



