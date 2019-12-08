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

func ConvertRadixStrToPlaceholder(strRadix string) (placeholderStr string) {
	switch strRadix {
	case HexStr:
		placeholderStr = HexFormatPlaceholder
		break
	case DecStr:
		placeholderStr = DecFormatPlaceholder
		break
	case OctStr:
		placeholderStr = OctFormatPlaceholder
		break
	case BinStr:
		placeholderStr = BinFormatPlaceholder
		break
	default:
		placeholderStr = ""
		break
	}
	return placeholderStr
}

func ConvertRadixStrToNum(strRadix string) (num int) {
	switch strRadix {
	case HexStr:
		num = Hex
		break
	case DecStr:
		num = Dec
		break
	case OctStr:
		num = Oct
		break
	case BinStr:
		num = Bin
		break
	default:
		num = 0
		break
	}
	return num
}
