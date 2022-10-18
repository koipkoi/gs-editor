package widgets

import (
	"strings"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TCustomListBox struct {
	*vcl.TListBox
	oddColor          types.TColor
	evenColor         types.TColor
	selectedColor     types.TColor
	textColor         types.TColor
	selectedTextColor types.TColor
}

func NewCustomListBox(owner vcl.IComponent) *TCustomListBox {
	inst := new(TCustomListBox)
	inst.TListBox = vcl.NewListBox(owner)

	inst.oddColor = 0xffffff
	inst.evenColor = 0xfff8f0
	inst.selectedColor = 0xd77800
	inst.textColor = 0x000000
	inst.selectedTextColor = 0xffffff

	inst.TListBox.SetDoubleBuffered(true)
	inst.TListBox.SetStyle(types.LbOwnerDrawVariable)
	inst.TListBox.SetOnMeasureItem(inst.onMeasureItem)
	inst.TListBox.SetOnDrawItem(inst.onDrawItem)

	return inst
}

func (inst *TCustomListBox) onMeasureItem(control *vcl.TWinControl, index int32, height *int32) {
	if index < inst.Items().Count() {
		canvas := inst.Canvas()
		text := inst.Items().Strings(index)
		lines := len(strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n"))
		*height = canvas.TextHeight(text) * int32(lines)
	}
}

func (inst *TCustomListBox) onDrawItem(control vcl.IWinControl, index int32, aRect types.TRect, state types.TOwnerDrawState) {
	canvas := inst.Canvas()
	isSelected := state.In(types.OdSelected)

	if isSelected {
		canvas.Font().SetColor(inst.selectedTextColor)
		canvas.Brush().SetColor(inst.selectedColor)
	} else {
		canvas.Font().SetColor(inst.textColor)
		if index%2 == 0 {
			canvas.Brush().SetColor(inst.evenColor)
		} else {
			canvas.Brush().SetColor(inst.oddColor)
		}
	}

	canvas.FillRect(aRect)

	text := inst.Items().Strings(index)
	boxHeight := aRect.Bottom - aRect.Top

	lines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	lineOffset := int32((boxHeight) / int32(len(lines)))
	textOffset := (lineOffset - canvas.TextHeight(text)) / 2
	for i, s := range lines {
		canvas.TextOut(aRect.Left, aRect.Top+(lineOffset*int32(i))+textOffset, s)
	}
}

func (inst *TCustomListBox) OddColor() types.TColor {
	return inst.oddColor
}

func (inst *TCustomListBox) SetOddColor(color types.TColor) {
	inst.oddColor = color
	inst.Invalidate()
}

func (inst *TCustomListBox) EvenColor() types.TColor {
	return inst.evenColor
}

func (inst *TCustomListBox) SetEvenColor(color types.TColor) {
	inst.evenColor = color
	inst.Invalidate()
}

func (inst *TCustomListBox) SelectedColor() types.TColor {
	return inst.selectedColor
}

func (inst *TCustomListBox) SetSelectedColor(color types.TColor) {
	inst.selectedColor = color
	inst.Invalidate()
}

func (inst *TCustomListBox) TextColor() types.TColor {
	return inst.textColor
}

func (inst *TCustomListBox) SetTextColor(color types.TColor) {
	inst.textColor = color
	inst.Invalidate()
}

func (inst *TCustomListBox) SelectedTextColor() types.TColor {
	return inst.selectedTextColor
}

func (inst *TCustomListBox) SetSelectedTextColor(color types.TColor) {
	inst.selectedTextColor = color
	inst.Invalidate()
}
