package helpers

import "github.com/ying32/govcl/vcl"

func NewMenuItem(owner vcl.IComponent, creatorFunc func(*vcl.TMenuItem)) *vcl.TMenuItem {
	inst := vcl.NewMenuItem(owner)
	creatorFunc(inst)
	return inst
}

func NewToolBarButton(owner vcl.IWinControl, creatorFunc func(*vcl.TToolButton)) *vcl.TToolButton {
	inst := vcl.NewToolButton(owner)
	inst.SetParent(owner)
	creatorFunc(inst)
	return inst
}
