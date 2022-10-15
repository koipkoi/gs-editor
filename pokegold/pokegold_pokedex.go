package pokegold

import (
	models "gs-editor/pokegold/models"
	"gs-editor/pokegold/utils"
)

type PokedexConverter struct{}

func (*PokedexConverter) Read(pokegold *Pokegold) {
	var buffer []byte
	var address int

	pokegold.Pokedex = nil

	for i := 0; i < PokemonCount; i++ {
		pokegold.Pokedex = append(pokegold.Pokedex, models.Pokedex{})

		if i < 128 {
			address = utils.ConvertToAddress(0x68, utils.SliceBytes(pokegold.rawBytes, 0x442ff+(i*2), 2))
		} else {
			address = utils.ConvertToAddress(0x69, utils.SliceBytes(pokegold.rawBytes, 0x443ff+((i-128)*2), 2))
		}

		buffer = nil
		for pokegold.rawBytes[address] != 0x50 {
			buffer = append(buffer, pokegold.rawBytes[address])
			address++
		}
		address++
		pokegold.Pokedex[i].SpecificName = utils.TextDecode(buffer)

		pokegold.Pokedex[i].Height = pokegold.rawBytes[address]
		address++
		pokegold.Pokedex[i].Weight = int(pokegold.rawBytes[address]) | (int(pokegold.rawBytes[address+1]) << 8)
		address += 2

		buffer = nil
		for pokegold.rawBytes[address] != 0x50 {
			buffer = append(buffer, pokegold.rawBytes[address])
			address++
		}
		pokegold.Pokedex[i].Description = utils.TextDecode(buffer)
	}
}

func (*PokedexConverter) Write(pokegold *Pokegold) {
	// todo 추가
}
