package pokegold

import (
	"gs-editor/pokegold/utils"
)

type Strings struct {
	ItemNames         []string
	TrainerClassNames []string
	PokemonNames      []string
	MoveNames         []string

	ItemDescriptions []string
	MoveDescriptions []string

	MoveTypeNames []string
}

type StringsConverter struct{}

func (*StringsConverter) Read(pokegold *Pokegold) {
	// item_name
	{
		pokegold.Strings.ItemNames = nil
		address := utils.ConvertToAddressBy3BytePointer(utils.SliceBytes(pokegold.rawBytes, 0x35cc, 3))
		for i := 0; i < ItemCount; i++ {
			var buffer []byte
			for pokegold.rawBytes[address] != 0x50 {
				buffer = append(buffer, pokegold.rawBytes[address])
				address++
			}
			address++
			pokegold.Strings.ItemNames = append(pokegold.Strings.ItemNames, utils.TextDecode(buffer))
		}
	}

	// trainer_class
	{
		pokegold.Strings.TrainerClassNames = nil
		address := utils.ConvertToAddressBy3BytePointer(utils.SliceBytes(pokegold.rawBytes, 0x35d5, 3))
		for i := 0; i <= TrainerClassNameCount; i++ {
			var buffer []byte
			for pokegold.rawBytes[address] != 0x50 {
				buffer = append(buffer, pokegold.rawBytes[address])
				address++
			}
			address++
			pokegold.Strings.TrainerClassNames = append(pokegold.Strings.TrainerClassNames, utils.TextDecode(buffer))
		}
	}

	// pokemon_name
	{
		pokegold.Strings.PokemonNames = nil
		for i := 0; i < PokemonCount; i++ {
			var buffer []byte
			address := utils.ConvertToAddressBy3BytePointer(utils.SliceBytes(pokegold.rawBytes, 0x35c3, 3)) + (i * 10)
			for j := 0; j < 10; j++ {
				if pokegold.rawBytes[address] == 0x50 {
					break
				}
				buffer = append(buffer, pokegold.rawBytes[address])
				address++
			}
			pokegold.Strings.PokemonNames = append(pokegold.Strings.PokemonNames, utils.TextDecode(buffer))
		}
	}

	// move_name
	{
		pokegold.Strings.MoveNames = nil
		address := utils.ConvertToAddressBy3BytePointer(utils.SliceBytes(pokegold.rawBytes, 0x35c6, 3))
		for i := 0; i < MoveCount; i++ {
			var buffer []byte
			for pokegold.rawBytes[address] != 0x50 {
				buffer = append(buffer, pokegold.rawBytes[address])
				address++
			}
			address++
			pokegold.Strings.MoveNames = append(pokegold.Strings.MoveNames, utils.TextDecode(buffer))
		}
	}

	// item_description
	{
		pokegold.Strings.ItemDescriptions = nil

		for i := 0; i < ItemCount; i++ {
			bank := utils.ConvertToBank(0x1b8000)
			address := utils.ConvertToAddress(bank, utils.SliceBytes(pokegold.rawBytes, 0x1b8000+(i*2), 2))

			var buffer []byte
			for pokegold.rawBytes[address] != 0x50 {
				buffer = append(buffer, pokegold.rawBytes[address])
				address++
			}

			pokegold.Strings.ItemDescriptions = append(pokegold.Strings.ItemDescriptions, utils.TextDecode(buffer))
		}
	}

	// move_description
	{
		pokegold.Strings.MoveDescriptions = nil

		for i := 0; i < MoveCount; i++ {
			bank := utils.ConvertToBank(0x1b4000)
			address := utils.ConvertToAddress(bank, utils.SliceBytes(pokegold.rawBytes, 0x1b4000+(i*2), 2))

			var buffer []byte
			for pokegold.rawBytes[address] != 0x50 {
				buffer = append(buffer, pokegold.rawBytes[address])
				address++
			}

			pokegold.Strings.MoveDescriptions = append(pokegold.Strings.MoveDescriptions, utils.TextDecode(buffer))
		}
	}

	// move_type_name
	{
		pokegold.Strings.MoveTypeNames = nil
		for i := 0; i < MoveTypeNameCount; i++ {
			bank := utils.ConvertToBank(0x50a57)
			address := utils.ConvertToAddress(bank, utils.SliceBytes(pokegold.rawBytes, 0x50a57+(i*2), 2))
			var buffer []byte
			for pokegold.rawBytes[address] != 0x50 {
				buffer = append(buffer, pokegold.rawBytes[address])
				address++
			}
			pokegold.Strings.MoveTypeNames = append(pokegold.Strings.MoveTypeNames, utils.TextDecode(buffer))
		}
	}
}

func (*StringsConverter) Write(pokegold *Pokegold) {
	// todo 추가
}

