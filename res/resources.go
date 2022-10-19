package res

import (
	"os"
	"path/filepath"

	"github.com/ying32/govcl/vcl"
)

const (
	appName     = "gs-editor"
	appTempName = "temp"
)

type Resources struct {
	bitmaps            []vcl.TBitmap
	icons              []vcl.TIcon
	imageList          *vcl.TImageList
	imageListItemCache map[string]int32
	imageListIndex     int32
}

func NewResources(owner vcl.IComponent) *Resources {
	return &Resources{
		imageList:          vcl.NewImageList(owner),
		imageListItemCache: map[string]int32{},
		imageListIndex:     0,
	}
}

func GetAppDataPath() string {
	path, _ := os.UserConfigDir()
	ret := filepath.Join(path, appName)
	os.MkdirAll(ret, 0777)
	return ret
}

func GetAppTempPath() string {
	path, _ := os.UserConfigDir()
	ret := filepath.Join(path, appName, appTempName)
	os.MkdirAll(ret, 0777)
	return ret
}

func (it *Resources) GetImageList() *vcl.TImageList {
	return it.imageList
}

func (it *Resources) Free() {
	for _, bitmap := range it.bitmaps {
		bitmap.Free()
	}
	for _, icon := range it.icons {
		icon.Free()
	}
	it.bitmaps = nil
	it.icons = nil
	it.imageList = nil
}
