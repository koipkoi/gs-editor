package res

import (
	"crypto/md5"
	"embed"
	_ "embed"
	"encoding/hex"
	"errors"
	"os"
	"path/filepath"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

var (
	//go:embed icons\*.bmp
	bmpFiles embed.FS

	//go:embed icons\*.ico
	iconFiles embed.FS
)

func (it *Resources) GetBitmap(name string) *vcl.TBitmap {
	read, readErr := bmpFiles.ReadFile("icons/" + name)
	if readErr != nil {
		return nil
	}

	newTempFile := getTemporaryFileName(read)
	if _, err := os.Stat(newTempFile); errors.Is(err, os.ErrNotExist) {
		writeErr := os.WriteFile(newTempFile, read, 0644)
		if writeErr != nil {
			return nil
		}
	}

	bitmap := vcl.NewBitmap()
	bitmap.LoadFromFile(newTempFile)
	bitmap.SetTransparent(true)
	bitmap.SetTransparentMode(types.TmFixed)

	it.bitmaps = append(it.bitmaps, *bitmap)

	return bitmap
}

func (it *Resources) GetIcon(name string) *vcl.TIcon {
	read, readErr := iconFiles.ReadFile("icons/" + name)
	if readErr != nil {
		return nil
	}

	newTempFile := getTemporaryFileName(read)
	if _, err := os.Stat(newTempFile); errors.Is(err, os.ErrNotExist) {
		writeErr := os.WriteFile(newTempFile, read, 0644)
		if writeErr != nil {
			return nil
		}
	}

	icon := vcl.NewIcon()
	icon.LoadFromFile(newTempFile)
	icon.SetTransparent(true)

	it.icons = append(it.icons, *icon)

	return icon
}

func getTemporaryFileName(read []byte) string {
	hash := md5.Sum(read)
	hashName := hex.EncodeToString(hash[:])
	return filepath.Join(GetAppTempPath(), hashName)
}
