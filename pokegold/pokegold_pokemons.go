package pokegold

import (
	models "gs-editor/pokegold/models"
	"gs-editor/pokegold/utils"
)

type PokemonsConverter struct{}

func (*PokemonsConverter) Read(pokegold *Pokegold) {
	var buffer []byte
	var bank byte
	var address int

	pokegold.Pokemons = nil
	for i := 0; i < PokemonCount; i++ {
		// 기본 정보
		pokegold.Pokemons = append(pokegold.Pokemons, *models.NewPokemonFromBytes(utils.SliceBytes(pokegold.rawBytes, 0x51bdf+(i*32), 32)))

		// 진화 정보 / 배우는 기술
		{
			bank = utils.ConvertToBank(0x423ed)
			address = utils.ConvertToAddress(bank, utils.SliceBytes(pokegold.rawBytes, 0x423ed+(i*2), 2))

			buffer = nil
			for pokegold.rawBytes[address] != 0 {
				buffer = append(buffer, pokegold.rawBytes[address])
				address++
			}
			address++
			pokegold.Pokemons[i].Evolutions = models.NewEvolutionsFromBytes(buffer)

			buffer = nil
			for pokegold.rawBytes[address] != 0 {
				buffer = append(buffer, pokegold.rawBytes[address])
				address++
			}
			pokegold.Pokemons[i].LearnMoves = models.NewLearnMovesFromBytes(buffer)
		}
	}
}

func (*PokemonsConverter) Write(pokegold *Pokegold) {
	// todo 추가
}
