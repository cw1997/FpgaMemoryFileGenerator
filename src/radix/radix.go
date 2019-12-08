// project : FpgaMemoryFileGenerator
// author : cw1997 [www.changwei.me]
// datetime : 2019-12-08 6:11
// repo : https://github.com/cw1997/FpgaMemoryFileGenerator

package radix

const (
	Hex = 16
	Dec = 10
	Oct = 8
	Bin = 2

	HexStr = "HEX"
	DecStr = "DEC"
	OctStr = "OCT"
	BinStr = "BIN"

	HexFormatPlaceholder = "x"
	DecFormatPlaceholder = "d"
	OctFormatPlaceholder = "o"
	BinFormatPlaceholder = "b"
)

func ConvertRadixStrToPlaceholder(intRadix int) (placeholderStr string) {
	switch intRadix {
	case Hex:
		placeholderStr = HexFormatPlaceholder
		break
	case Dec:
		placeholderStr = DecFormatPlaceholder
		break
	case Oct:
		placeholderStr = OctFormatPlaceholder
		break
	case Bin:
		placeholderStr = BinFormatPlaceholder
		break
	default:
		placeholderStr = ""
		break
	}
	return placeholderStr
}

func CheckRadix(intRadix int) bool {
	if intRadix == Hex || intRadix == Dec || intRadix == Oct || intRadix == Bin {
		return true
	}
	return false
}
