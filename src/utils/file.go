// project : FpgaMemoryFileGenerator
// author : cw1997 [www.changwei.me]
// datetime : 2019/12/8 21:07
// repo : https://github.com/cw1997/FpgaMemoryFileGenerator

package utils

import "io/ioutil"

func WriteStrToFile(data string, path string) error {
	err := ioutil.WriteFile(path, []byte(data), 0777)
	return err
}

func ReadBytesFromFile(data string, path string) ([]byte, error) {
	b, err := ioutil.ReadFile(path)
	return b, err
}
