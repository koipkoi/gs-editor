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
	textEncodeBuffer := make([]byte, 1024)
	firstAddress := 0x1a0000
	secondAddress := 0x1a4000

	utils.FillBytes(pokegold.rawBytes, 0, 0x1a0000, 0x8000)

	var length, address, pointerAddress int
	for i := 0; i < PokemonCount; i++ {
		if i < 128 {
			address = firstAddress
			pointerAddress = 0x442ff + (i * 2)
		} else {
			address = secondAddress
			pointerAddress = 0x443ff + ((i - 128) * 2)
		}

		utils.CopyBytes(pokegold.rawBytes, pointerAddress, utils.ConvertToPointer(address))

		length = utils.TextEncodeBuffered(textEncodeBuffer, pokegold.Pokedex[i].SpecificName)
		address = utils.CopyBytesWithLength(pokegold.rawBytes, address, textEncodeBuffer, length)
		pokegold.rawBytes[address] = 0x50
		address++

		pokegold.rawBytes[address] = pokegold.Pokedex[i].Height
		address++

		pokegold.rawBytes[address] = byte(pokegold.Pokedex[i].Weight & 0x00ff)
		address++
		pokegold.rawBytes[address] = byte((pokegold.Pokedex[i].Weight & 0xff00) >> 8)
		address++

		length = utils.TextEncodeBuffered(textEncodeBuffer, pokegold.Pokedex[i].Description)
		address = utils.CopyBytesWithLength(pokegold.rawBytes, address, textEncodeBuffer, length)
		pokegold.rawBytes[address] = 0x50
		address++

		if i < 128 {
			firstAddress = address
		} else {
			secondAddress = address
		}
	}
}
