// project : FpgaMemoryFileGenerator
// author : cw1997 [www.changwei.me]
// datetime : 2019-12-08 14:22
// repo : https://github.com/cw1997/FpgaMemoryFileGenerator

package mif

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cw1997/FpgaMemoryFileGenerator/radix"
)

type Format struct {
	depth int
	width int
	addressRadix string
	dataRadix string
}

func NewMifGenerator(depth int, width int, addressRadix string, dataRadix string) Format {
	return Format{depth,
		width,
		strings.ToUpper(addressRadix),
		strings.ToUpper(dataRadix)}
}

func (m Format) Generate(data []byte) string {
	content := ""
	addressRadixNum := radix.ConvertRadixStrToNum(m.addressRadix)
	addressRadixFormatPlaceholder := radix.ConvertRadixStrToPlaceholder(m.addressRadix)
	addressLength := strconv.Itoa(m.depth / addressRadixNum + 1)

	dataRadixNum := radix.ConvertRadixStrToNum(m.dataRadix)
	dataRadixFormatPlaceholder := radix.ConvertRadixStrToPlaceholder(m.dataRadix)
	dataWidth := strconv.Itoa(m.width / dataRadixNum + 1)

	// example: %s%08x:%08x\n
	contentFormat := "%s%0" + addressLength + addressRadixFormatPlaceholder + ":%0" + dataWidth + dataRadixFormatPlaceholder + "\n"



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
`, m.depth, m.width, m.addressRadix, m.dataRadix, content)
	return fileContent
}
