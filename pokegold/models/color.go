package models

import (
	"image/color"

	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
)

const (
	gbRedBits   int = 0b000000000011111
	gbBlueBits  int = 0b000001111100000
	gbGreenBits int = 0b111110000000000
)

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

func NewColorFromColor(color color.Color) Color {
	r, g, b, _ := color.RGBA()
	r = r / 8 << 0
	g = g / 8 << 5
	b = b / 8 << 10
	colorInt := r | g | b
	return Color{
		highByte: byte((colorInt & 0xff00) >> 8),
		lowByte:  byte((colorInt & 0x00ff) >> 0),
	}
}

func NewColorFromVCLColor(color types.TColor) Color {
	r := int(color.R()/8) << 0
	g := int(color.G()/8) << 5
	b := int(color.B()/8) << 10
	colorInt := r | g | b
	return Color{
		highByte: byte((colorInt & 0xff00) >> 8),
		lowByte:  byte((colorInt & 0x00ff) >> 0),
	}
}

func (color *Color) ToBytes() []byte {
	return []byte{
		color.lowByte,
		color.highByte,
	}
}

func (c *Color) ToColor() color.Color {
	colorInt := (int(c.highByte) << 8) | int(c.lowByte)
	r := byte(((colorInt & gbRedBits) >> 0) * 8)
	g := byte(((colorInt & gbBlueBits) >> 5) * 8)
	b := byte(((colorInt & gbGreenBits) >> 10) * 8)
	return color.RGBA{r, g, b, 0xff}
}

func (color *Color) ToVCLColor() types.TColor {
	colorInt := (int(color.highByte) << 8) | int(color.lowByte)
	r := byte(((colorInt & gbRedBits) >> 0) * 8)
	g := byte(((colorInt & gbBlueBits) >> 5) * 8)
	b := byte(((colorInt & gbGreenBits) >> 10) * 8)
	return types.TColor(colors.RGB(r, g, b))
}
