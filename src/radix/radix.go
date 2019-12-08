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

func ConvertRadixNumToStr(intRadix int) (radixStr string) {
	switch intRadix {
	case Hex:
		radixStr = HexStr
		break
	case Dec:
		radixStr = DecStr
		break
	case Oct:
		radixStr = OctStr
		break
	case Bin:
		radixStr = BinStr
		break
	default:
		radixStr = ""
		break
	}
	return radixStr
}
