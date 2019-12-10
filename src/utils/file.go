// project : FpgaMemoryFileGenerator
// author : cw1997 [www.changwei.me]
// datetime : 2019/12/8 21:07
// repo : https://github.com/cw1997/FpgaMemoryFileGenerator

package utils

import "io/ioutil"

func WriteBytesToFile(data []byte, path string) error {
	err := ioutil.WriteFile(path, data, 0666)
	return err
}

func ReadBytesFromFile(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(path)
	return b, err
}

func ReadStringFromFile(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	return string(b), err
}
