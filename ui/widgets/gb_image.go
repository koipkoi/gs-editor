package widgets

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TGBImage struct {
	*vcl.TPaintBox
	rows    int
	columns int
	source  []byte
	palette []types.TColor
}

func NewGBImage(owner vcl.IComponent) *TGBImage {
	inst := new(TGBImage)
	inst.TPaintBox = vcl.NewPaintBox(owner)
	inst.TPaintBox.SetOnPaint(inst.onPaint)
	return inst
}

func (inst *TGBImage) onPaint(vcl.IObject) {
	canvas := inst.Canvas()

	for i := 0; i < len(inst.source); i += 2 {
		first := inst.source[i+0]
		second := inst.source[i+1]

		x := (i / (16 * inst.columns)) * 8
		y := (i % (16 * inst.columns)) / 2

		for pixelX := 0; pixelX < 8; pixelX++ {
			hi := byte((second >> (7 - pixelX)) & 1)
			lo := byte((first >> (7 - pixelX)) & 1)
			colorIndex := (hi << 1) | lo

			if int(colorIndex) < len(inst.palette) {
				canvas.SetPixels(int32(x+pixelX), int32(y), inst.palette[colorIndex])
			}
		}
	}
}

func (inst *TGBImage) Rows() int {
	return inst.rows
}

func (inst *TGBImage) SetRows(value int) {
	inst.rows = value
	inst.Invalidate()
}

func (inst *TGBImage) Columns() int {
	return inst.columns
}

func (inst *TGBImage) SetColumns(value int) {
	inst.columns = value
	inst.Invalidate()
}

func (inst *TGBImage) Source() []byte {
	return inst.source
}

func (inst *TGBImage) SetSource(value []byte) {
	inst.source = value
	inst.Invalidate()
}

func (inst *TGBImage) Palette() []types.TColor {
	return inst.palette
}

func (inst *TGBImage) SetPalette(value []types.TColor) {
	inst.palette = value
	inst.Invalidate()
}
