package main

import (
	"gs-editor/ui/forms"

	"github.com/ying32/govcl/vcl"
)

func main() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := forms.NewMainForm()
	mainForm.ShowModal()
}
