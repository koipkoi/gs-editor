package forms

import (
	"gs-editor/pokegold"
	"gs-editor/res"
	"gs-editor/ui/helpers"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
)

type TMainForm struct {
	*vcl.TForm

	res      *res.Resources
	pokegold *pokegold.Pokegold
	menu     *vcl.TMainMenu
	toolBar  *vcl.TToolBar
}

func NewMainForm() *TMainForm {
	f := new(TMainForm)
	f.TForm = vcl.Application.CreateForm()
	f.res = res.NewResources(f)
	f.pokegold = pokegold.NewPokegold()

	f.TForm.SetIcon(f.res.GetIcon("app_icon.ico"))
	f.TForm.SetShowInTaskBar(types.StAlways)

	f.TForm.SetCaption("GS 에디터")
	f.pokegold.AddOnChanged(func(p *pokegold.Pokegold) {
		f.TForm.SetCaption("GS 에디터 - " + p.Filename)
	})

	f.TForm.SetOnCloseQuery(func(sender vcl.IObject, canClose *bool) {
		f.res.Free()
	})

	f.buildMainMenu()
	f.buildToolBar()

	return f
}

func (f *TMainForm) buildMainMenu() {
	f.menu = vcl.NewMainMenu(f)
	f.menu.SetImages(f.res.GetImageList())

	{
		menu := vcl.NewMenuItem(f)
		menu.SetCaption("파일(&F)")
		f.menu.Items().Add(menu)

		menu.Add(helpers.NewMenuItem(f.menu, func(ti *vcl.TMenuItem) {
			ti.SetCaption("열기(&O)...")
			ti.SetImageIndex(f.res.GetImageListItem(f, "if_folder.bmp", types.TColor(colors.RGB(255, 0, 255))))
			ti.SetShortCutFromString("Ctrl+O")
			ti.SetOnClick(f.onOpenClick)
		}))

		menu.Add(helpers.NewMenuItem(f.menu, func(ti *vcl.TMenuItem) {
			ti.SetCaption("저장(&S)")
			ti.SetEnabled(false)
			ti.SetImageIndex(f.res.GetImageListItem(f, "if_save.bmp", types.TColor(colors.RGB(255, 0, 255))))
			ti.SetShortCutFromString("Ctrl+S")
			ti.SetOnClick(f.onSaveClick)

			f.pokegold.AddOnChanged(func(p *pokegold.Pokegold) {
				ti.SetEnabled(p.IsOpen)
			})
		}))

		menu.Add(helpers.NewMenuItem(f.menu, func(ti *vcl.TMenuItem) {
			ti.SetCaption("-")
		}))

		menu.Add(helpers.NewMenuItem(f.menu, func(ti *vcl.TMenuItem) {
			ti.SetCaption("종료(&X)")
			ti.SetImageIndex(f.res.GetImageListItem(f, "if_exit.bmp", types.TColor(colors.RGB(255, 0, 255))))
			ti.SetShortCutFromString("Alt+F4")
			ti.SetOnClick(f.onExitClick)
		}))
	}

	{
		menu := vcl.NewMenuItem(f)
		menu.SetCaption("게임(&G)")
		f.menu.Items().Add(menu)

		menu.Add(helpers.NewMenuItem(f.menu, func(ti *vcl.TMenuItem) {
			ti.SetCaption("테스트 플레이(&P)...")
			ti.SetEnabled(false)
			ti.SetImageIndex(f.res.GetImageListItem(f, "if_play.bmp", types.TColor(colors.RGB(255, 0, 255))))
			ti.SetShortCutFromString("F5")
			ti.SetOnClick(f.onTestPlayClick)

			f.pokegold.AddOnChanged(func(p *pokegold.Pokegold) {
				ti.SetEnabled(p.IsOpen)
			})
		}))

		menu.Add(helpers.NewMenuItem(f.menu, func(ti *vcl.TMenuItem) {
			ti.SetCaption("-")
		}))

		menu.Add(helpers.NewMenuItem(f.menu, func(ti *vcl.TMenuItem) {
			ti.SetCaption("에뮬레이터 설정(&O)...")
			ti.SetImageIndex(f.res.GetImageListItem(f, "if_settings.bmp", types.TColor(colors.RGB(255, 0, 255))))
			ti.SetOnClick(f.onEmulatorSettingsClick)
		}))
	}

	{
		menu := vcl.NewMenuItem(f)
		menu.SetCaption("도움말(&H)")
		f.menu.Items().Add(menu)

		menu.Add(helpers.NewMenuItem(f.menu, func(ti *vcl.TMenuItem) {
			ti.SetCaption("GS 에디터 정보(&A)...")
			ti.SetOnClick(f.onAppInformationClick)
		}))
	}
}

func (f *TMainForm) buildToolBar() {
	f.toolBar = vcl.NewToolBar(f)
	f.toolBar.SetParent(f)
	f.toolBar.SetImages(f.res.GetImageList())

	{
		helpers.NewToolBarButton(f.toolBar, func(tb *vcl.TToolButton) {
			tb.SetCaption("열기")
			tb.SetOnClick(f.onOpenClick)
			tb.SetImageIndex(f.res.GetImageListItem(f, "if_folder.bmp", types.TColor(colors.RGB(255, 0, 255))))
		})

		helpers.NewToolBarButton(f.toolBar, func(tb *vcl.TToolButton) {
			tb.SetCaption("저장")
			tb.SetEnabled(false)
			tb.SetOnClick(f.onSaveClick)
			tb.SetImageIndex(f.res.GetImageListItem(f, "if_save.bmp", types.TColor(colors.RGB(255, 0, 255))))

			f.pokegold.AddOnChanged(func(p *pokegold.Pokegold) {
				tb.SetEnabled(p.IsOpen)
			})
		})
	}

	helpers.NewToolBarButton(f.toolBar, func(tb *vcl.TToolButton) {
		tb.SetStyle(types.TbsSeparator)
	})

	{
		helpers.NewToolBarButton(f.toolBar, func(tb *vcl.TToolButton) {
			tb.SetCaption("테스트 플레이")
			tb.SetEnabled(false)
			tb.SetOnClick(f.onTestPlayClick)
			tb.SetImageIndex(f.res.GetImageListItem(f, "if_play.bmp", types.TColor(colors.RGB(255, 0, 255))))

			f.pokegold.AddOnChanged(func(p *pokegold.Pokegold) {
				tb.SetEnabled(p.IsOpen)
			})
		})
	}

	helpers.NewToolBarButton(f.toolBar, func(tb *vcl.TToolButton) {
		tb.SetStyle(types.TbsSeparator)
	})

	{
		helpers.NewToolBarButton(f.toolBar, func(tb *vcl.TToolButton) {
			tb.SetCaption("종료")
			tb.SetOnClick(f.onExitClick)
			tb.SetImageIndex(f.res.GetImageListItem(f, "if_exit.bmp", types.TColor(colors.RGB(255, 0, 255))))
		})
	}
}

func (f *TMainForm) onOpenClick(vcl.IObject) {
	dialog := vcl.NewOpenDialog(f)
	dialog.SetTitle("열기")
	dialog.SetFilter("지원하는 파일|*.gb; *.gbc; *.bin|모든 파일|*.*")
	if dialog.Execute() {
		f.pokegold.ReadRom(dialog.FileName())
	}
}

func (f *TMainForm) onSaveClick(vcl.IObject) {
	f.pokegold.WriteRom(f.pokegold.Filename)
}

func (f *TMainForm) onExitClick(vcl.IObject) {
	f.Close()
}

func (f *TMainForm) onTestPlayClick(vcl.IObject) {
	// todo 추가
}

func (f *TMainForm) onEmulatorSettingsClick(vcl.IObject) {
	// todo 추가
}

func (f *TMainForm) onAppInformationClick(vcl.IObject) {
	// todo 추가
}
