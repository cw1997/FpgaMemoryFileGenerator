// project : FpgaMemoryFileGenerator
// author : cw1997 [www.changwei.me]
// datetime : 2019-12-08 19:02
// repo : https://github.com/cw1997/FpgaMemoryFileGenerator

package coe

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cw1997/FpgaMemoryFileGenerator/radix"
)

type Format struct {
	width     int
	dataRadix int
}

func NewCoeGenerator(width int, dataRadix int) Format {
	if !radix.CheckRadix(dataRadix) {
		log.Fatalf("[Error] Input radix is not in [2 8 10 16]. \n")
	}
	return Format{
		width,
		dataRadix}
}

func (m Format) Generate(data []byte) string {
	content := ""

	dataRadixNum := m.dataRadix
	//dataRadixStr := radix.ConvertRadixNumToStr(m.dataRadix)
	dataRadixFormatPlaceholder := radix.ConvertRadixNumToPlaceholder(m.dataRadix)
	dataWidth := strconv.Itoa(m.width/dataRadixNum + 1)

	// example: %s%08x:%08x\n
	contentFormat := "%s%0" + dataWidth + dataRadixFormatPlaceholder + ",\n"

	dataLength := len(data)

	for i := 0; i < dataLength; i++ {
		content = fmt.Sprintf(contentFormat, content, data[i])
	}

	// remove suffix : ",\n"
	content = strings.TrimSuffix(content, ",\n")

	fileContent := fmt.Sprintf(`memory_initialization_radix=%d;
memory_initialization_vector=
%s;
`, dataRadixNum, content)
	return fileContent
}
