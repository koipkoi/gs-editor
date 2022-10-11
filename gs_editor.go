package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

func main() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("GS 에디터")
	mainForm.SetIcon(nil)
	mainForm.EnabledMaximize(false)
	mainForm.EnabledMinimize(true)
	mainForm.SetDoubleBuffered(true)
	mainForm.SetBorderStyle(types.BsSingle)
	mainForm.SetWidth(600)
	mainForm.SetHeight(600)

	vcl.Application.Run()
}
