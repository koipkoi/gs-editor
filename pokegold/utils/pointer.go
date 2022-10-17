package utils

func ConvertToBank(address int) byte {
	return byte(address / 0x4000)
}

func ConvertToPointer(address int) []byte {
	var pointer int
	if address >= 0x4000 {
		pointer = (address % 0x4000) + 0x4000
	} else {
		pointer = (address % 0x4000)
	}
	return []byte{
		byte(pointer & 0x00ff),
		byte((pointer & 0xff00) >> 8),
	}
}

func ConvertToPointerWithBank(address int) []byte {
	var pointer int
	if address >= 0x4000 {
		pointer = (address % 0x4000) + 0x4000
	} else {
		pointer = (address % 0x4000)
	}
	return []byte{
		byte(address / 0x4000),
		byte(pointer & 0x00ff),
		byte((pointer & 0xff00) >> 8),
	}
}

func ConvertToAddress(bank byte, pointer []byte) int {
	pointerAddress := int(pointer[0]) | (int(pointer[1]) << 8)
	return (int(bank) * 0x4000) | (pointerAddress - 0x4000)
}

func ConvertToAddressBy3BytePointer(pointer3bytes []byte) int {
	pointerAddress := int(pointer3bytes[1]) | (int(pointer3bytes[2]) << 8)
	return (int(pointer3bytes[0]) * 0x4000) | (pointerAddress - 0x4000)
}

func DecodeBank(bank byte) byte {
	switch bank {
	case 0x13:
		return 0x1f
	case 0x14:
		return 0x20
	case 0x1f:
		return 0x2e
	}
	return bank
}

func EncodeBank(bank byte) byte {
	switch bank {
	case 0x1f:
		return 0x13
	case 0x20:
		return 0x14
	case 0x2e:
		return 0x1f
	}
	return bank
}

func SliceBytes(s []byte, start, length int) []byte {
	return s[start : start+length]
}

func CopyBytes(bytes []byte, start int, src []byte) {
	for i := 0; i < len(src); i++ {
		bytes[start+i] = src[i]
	}
}

func CopyBytesWithLength(bytes []byte, start int, src []byte, length int) int {
	for i := 0; i < length; i++ {
		bytes[start+i] = src[i]
	}
	return start + length
}

func FillBytes(bytes []byte, v byte, start int, length int) {
	for i := 0; i < length; i++ {
		bytes[start+i] = v
	}
}
