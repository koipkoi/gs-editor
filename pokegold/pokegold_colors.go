package pokegold

import (
	models "gs-editor/pokegold/models"
	"gs-editor/pokegold/utils"
)

type Colors struct {
	PokemonColors      [][2]models.Color
	PokemonShinyColors [][2]models.Color
	TrainerColors      [][2]models.Color
}

type ColorsConverter struct{}

func (*ColorsConverter) Read(pokegold *Pokegold) {
	// pokemon
	{
		pokegold.Colors.PokemonColors = nil
		pokegold.Colors.PokemonShinyColors = nil

		for i := 0; i < PokemonCount; i++ {
			address := 0xad15 + (i * 8)
			pokegold.Colors.PokemonColors = append(pokegold.Colors.PokemonColors, [...]models.Color{
				models.NewColorFromBytes(utils.SliceBytes(pokegold.rawBytes, address, 2)),
				models.NewColorFromBytes(utils.SliceBytes(pokegold.rawBytes, address+2, 2)),
			})
			pokegold.Colors.PokemonShinyColors = append(pokegold.Colors.PokemonShinyColors, [...]models.Color{
				models.NewColorFromBytes(utils.SliceBytes(pokegold.rawBytes, address+4, 2)),
				models.NewColorFromBytes(utils.SliceBytes(pokegold.rawBytes, address+6, 2)),
			})
		}
	}

	// trainer
	{
		pokegold.Colors.TrainerColors = nil

		for i := 0; i < TrainerImageCount; i++ {
			pokegold.Colors.TrainerColors = append(pokegold.Colors.TrainerColors, [...]models.Color{
				models.NewColorFromBytes(utils.SliceBytes(pokegold.rawBytes, 0xb511+(i*4), 2)),
				models.NewColorFromBytes(utils.SliceBytes(pokegold.rawBytes, 0xb513+(i*4), 2)),
			})
		}
	}
}

func (*ColorsConverter) Write(pokegold *Pokegold) {
	// todo 추가
}
