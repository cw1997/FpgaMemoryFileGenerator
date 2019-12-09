// project : FpgaMemoryFileGenerator
// author : cw1997 [www.changwei.me]
// datetime : 2019-12-08 6:11
// repo : https://github.com/cw1997/FpgaMemoryFileGenerator

package radix

const (
	Hex = 16
	Dec = -10
	Uns = 10
	Oct = 8
	Bin = 2

	HexStr = "HEX"
	DecStr = "DEC"
	UnsStr = "UNS"
	OctStr = "OCT"
	BinStr = "BIN"

	HexFormatPlaceholder = "x"
	DecFormatPlaceholder = "d"
	OctFormatPlaceholder = "o"
	BinFormatPlaceholder = "b"
)

func CheckRadix(intRadix int) bool {
	if intRadix == Hex || intRadix == Dec || intRadix == Uns || intRadix == Oct || intRadix == Bin {
		return true
	}
	return false
}

func ConvertRadixNumToPlaceholder(intRadix int) (placeholderStr string) {
	switch intRadix {
	case Hex:
		placeholderStr = HexFormatPlaceholder
		break
	// golang's fmt.printf() doesn't have unsigned decimal placeholder
	// so it also uses %d instead of %u(c -> stdio.h -> printf())
	case Dec, Uns:
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

func ConvertDataRadixNumToStr(intRadix int) (radixStr string) {
	switch intRadix {
	case Hex:
		radixStr = HexStr
		break
	case Dec:
		radixStr = DecStr
		break
	case Uns:
		radixStr = UnsStr
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

func ConvertAddressRadixNumToStr(intRadix int) (radixStr string) {
	switch intRadix {
	case Hex:
		radixStr = HexStr
		break
	// address doesn't distinguish sign
	case Dec, Uns:
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
