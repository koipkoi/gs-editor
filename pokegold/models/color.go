package models

type Color struct {
	highByte byte
	lowByte  byte
}

func NewColorFromBytes(bytes []byte) Color {
	return Color{
		highByte: bytes[1],
		lowByte:  bytes[0],
	}
}

func (color *Color) ToBytes() []byte {
	return []byte{
		color.lowByte,
		color.highByte,
	}
}

// todo byte 데이터 > 컬러 정보 변환 추가
