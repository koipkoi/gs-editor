package widgets

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TStackPanel struct {
	*vcl.TPanel
	orientation types.TAlign
	children    []vcl.IControl
}

type TStackPanelBuilder struct {
	Owner *TStackPanel
}

// 수평으로 나열되는 패널
func NewHorizontalStackPanel(owner vcl.IWinControl, builder func(builder *TStackPanelBuilder)) *TStackPanel {
	panel := new(TStackPanel)
	panel.orientation = types.AlLeft
	panel.TPanel = NewBorderlessPanel(owner)
	panel.TPanel.SetAutoSize(true)
	builder(&TStackPanelBuilder{Owner: panel})
	panel.invalidateLayout()
	return panel
}

// 수직으로 나열되는 패널
func NewVerticalStackPanel(owner vcl.IWinControl, builder func(builder *TStackPanelBuilder)) *TStackPanel {
	panel := new(TStackPanel)
	panel.orientation = types.AlTop
	panel.TPanel = NewBorderlessPanel(owner)
	panel.TPanel.SetAutoSize(true)
	builder(&TStackPanelBuilder{Owner: panel})
	panel.invalidateLayout()
	return panel
}

func (panel *TStackPanel) invalidateLayout() {
	for i := 0; i < len(panel.children); i++ {
		childPanel := NewBorderlessPanel(panel)
		childPanel.SetParent(panel)
		childPanel.SetAlign(panel.orientation)
		childPanel.SetAutoSize(true)

		e := panel.children[len(panel.children)-i-1]
		e.SetParent(childPanel)
	}
}

func (builder *TStackPanelBuilder) AddHorizontalStackPanel(b func(builder *TStackPanelBuilder)) *TStackPanel {
	newPanel := NewHorizontalStackPanel(builder.Owner, b)
	builder.Owner.children = append(builder.Owner.children, newPanel)
	return newPanel
}

func (builder *TStackPanelBuilder) AddVerticalStackPanel(b func(builder *TStackPanelBuilder)) *TStackPanel {
	newPanel := NewVerticalStackPanel(builder.Owner, b)
	builder.Owner.children = append(builder.Owner.children, newPanel)
	return newPanel
}

func (builder *TStackPanelBuilder) Add(childBuilder func() vcl.IControl) {
	builder.Owner.children = append(builder.Owner.children, childBuilder())
}

func (builder *TStackPanelBuilder) AddInline(child vcl.IControl) {
	builder.Owner.children = append(builder.Owner.children, child)
}

func (builder *TStackPanelBuilder) Space(w, h int32) {
	space := NewSpace(builder.Owner, w, h)
	builder.Owner.children = append(builder.Owner.children, space)
}
