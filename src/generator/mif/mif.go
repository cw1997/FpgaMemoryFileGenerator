// project : FpgaMemoryFileGenerator
// author : cw1997 [www.changwei.me]
// datetime : 2019-12-08 14:22
// repo : https://github.com/cw1997/FpgaMemoryFileGenerator

package mif

import (
	"fmt"
	"github.com/cw1997/FpgaMemoryFileGenerator/radix"
	"log"
	"strconv"
)

type Format struct {
	depth        int
	width        int
	addressRadix int
	dataRadix    int
}

func NewMifGenerator(depth int, width int, addressRadix int, dataRadix int) Format {
	if !radix.CheckRadix(addressRadix) || !radix.CheckRadix(dataRadix) {
		log.Fatalf("[Error] Input radix is not in [2 8 10 16]. \n")
	}
	return Format{depth,
		width,
		addressRadix,
		dataRadix}
}

func (m Format) Generate(data []byte) []byte {
	content := ""
	addressRadixNum := m.addressRadix
	addressRadixStr := radix.ConvertAddressRadixNumToStr(m.addressRadix)
	addressRadixFormatPlaceholder := radix.ConvertRadixNumToPlaceholder(m.addressRadix)
	addressLength := strconv.Itoa(m.depth/addressRadixNum + 1)

	dataRadixNum := m.dataRadix
	dataRadixStr := radix.ConvertDataRadixNumToStr(m.dataRadix)
	dataRadixFormatPlaceholder := radix.ConvertRadixNumToPlaceholder(m.dataRadix)
	dataWidth := strconv.Itoa(m.width/dataRadixNum + 1)

	// example: %s%08x:%08x\n
	contentFormat := "%s%0" + addressLength + addressRadixFormatPlaceholder + ":%0" + dataWidth + dataRadixFormatPlaceholder + ";\n"

	for i := 0; i < m.depth; i++ {
		if i >= len(data) {
			content = fmt.Sprintf(contentFormat, content, i, 0)
		} else {
			content = fmt.Sprintf(contentFormat, content, i, data[i])
		}

	}

	fileContent := fmt.Sprintf(`DEPTH = %d;
WIDTH = %d;
ADDRESS_RADIX = %s;
DATA_RADIX = %s;
CONTENT
BEGIN
%sEND;
`, m.depth, m.width, addressRadixStr, dataRadixStr, content)
	return []byte(fileContent)
}
