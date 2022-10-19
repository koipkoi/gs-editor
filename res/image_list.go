package res

import (
	"fmt"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

func (res *Resources) GetImageListItem(
	componet vcl.IComponent,
	name string,
	transparentColor types.TColor,
) int32 {
	key := fmt.Sprintf("%s_%d", name, transparentColor)
	if index, ok := res.imageListItemCache[key]; ok {
		return index
	}

	res.GetImageList().AddMasked(res.GetBitmap(name), transparentColor)
	res.imageListItemCache[key] = res.imageListIndex
	res.imageListIndex++

	return res.imageListItemCache[key]
}
