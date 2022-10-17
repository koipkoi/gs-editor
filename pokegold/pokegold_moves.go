package pokegold

import (
	models "gs-editor/pokegold/models"
	"gs-editor/pokegold/utils"
)

type MovesConverter struct{}

func (*MovesConverter) Read(pokegold *Pokegold) {
	pokegold.Moves = nil
	for i := 0; i < MoveCount; i++ {
		address := 0x4172e + (i * 7)
		pokegold.Moves = append(pokegold.Moves, *models.NewMoveFromBytes(utils.SliceBytes(pokegold.rawBytes, address, 7)))
	}
}

func (*MovesConverter) Write(pokegold *Pokegold) {
	for i := 0; i < MoveCount; i++ {
		address := 0x4172e + (i * 7)
		bytes := pokegold.Moves[i].ToBytes()
		for j := 0; j < len(bytes); j++ {
			pokegold.rawBytes[address+j] = bytes[j]
		}
	}
}
