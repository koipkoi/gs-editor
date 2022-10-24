package widgets

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

// 외선을 완전히 제거한 패널
func NewBorderlessPanel(owner vcl.IComponent) *vcl.TPanel {
	panel := vcl.NewPanel(owner)
	panel.SetBorderStyle(types.BsNone)
	panel.SetBevelInner(types.BvNone)
	panel.SetBevelOuter(types.BvNone)
	return panel
}

// 스크롤 상자
func NewScrollBox(owner vcl.IComponent) *vcl.TScrollBox {
	scroll := vcl.NewScrollBox(owner)
	scroll.HorzScrollBar().SetTracking(true)
	scroll.VertScrollBar().SetTracking(true)
	return scroll
}

// 메뉴 항목 (익명 생성용)
func NewMenuItem(owner vcl.IComponent, creatorFunc func(*vcl.TMenuItem)) *vcl.TMenuItem {
	inst := vcl.NewMenuItem(owner)
	creatorFunc(inst)
	return inst
}

// 툴바 항목 (익명 생성용)
func NewToolBarButton(owner vcl.IWinControl, creatorFunc func(*vcl.TToolButton)) *vcl.TToolButton {
	inst := vcl.NewToolButton(owner)
	inst.SetParent(owner)
	creatorFunc(inst)
	return inst
}

// 영역 분리용 라벨
func NewSectionLabel(owner vcl.IWinControl, caption string) *vcl.TPanel {
	panel := NewBorderlessPanel(owner)
	panel.SetParent(owner)
	panel.SetAlign(types.AlClient)

	label := vcl.NewLabel(panel)
	label.SetParent(panel)
	label.SetAlign(types.AlLeft)
	label.SetCaption(caption)
	label.SetLayout(types.TlCenter)
	label.BorderSpacing().SetRight(4)

	bevel := vcl.NewBevel(panel)
	bevel.SetParent(panel)
	bevel.SetAlign(types.AlClient)
	bevel.Constraints().SetMinHeight(2)
	bevel.Constraints().SetMaxHeight(2)

	panel.SetOnResize(func(sender vcl.IObject) {
		bevel.BorderSpacing().SetTop((panel.Height() - 2) / 2)
	})

	return panel
}

// 여백 패널
func NewSpace(owner vcl.IWinControl, w, h int32) *vcl.TPanel {
	panel := NewBorderlessPanel(owner)
	panel.SetParent(owner)
	panel.SetWidth(w)
	panel.SetHeight(h)
	return panel
}
