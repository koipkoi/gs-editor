package pokegold

import (
	models "gs-editor/pokegold/models"
)

type ItemsConverter struct{}

func (*ItemsConverter) Read(pokegold *Pokegold) {
	pokegold.Items = nil

	for i := 0; i < ItemCount; i++ {
		address := 0x697b + (i * 7)
		read := pokegold.rawBytes[address : address+7]
		pokegold.Items = append(pokegold.Items, *models.NewItemFromBytes(read))
	}
}

func (*ItemsConverter) Write(pokegold *Pokegold) {
	// todo 추가
}
