// project : FpgaMemoryFileGenerator
// author : cw1997 [www.changwei.me]
// datetime : 2019/12/8 20:27
// repo : https://github.com/cw1997/FpgaMemoryFileGenerator

package converter

type Converter interface {
	Convert(input interface{}) []byte
}
