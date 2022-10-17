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
		for i := 0; i < PokemonNameCount; i++ {
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
	utils.FillBytes(pokegold.rawBytes, 0, 0x1b0000, 0x4000)

	textEncodeBuffer := make([]byte, 1024)
	nameWriteAddress := 0x1b0000

	{
		utils.CopyBytes(pokegold.rawBytes, 0x35cc, utils.ConvertToPointerWithBank(nameWriteAddress))
		utils.CopyBytes(pokegold.rawBytes, 0x515cd, utils.ConvertToPointer(nameWriteAddress))
		utils.CopyBytes(pokegold.rawBytes, 0x515d7, utils.ConvertToPointer(nameWriteAddress))

		for i := 0; i < ItemCount; i++ {
			length := utils.TextEncodeBuffered(textEncodeBuffer, pokegold.Strings.ItemNames[i])
			nameWriteAddress = utils.CopyBytesWithLength(pokegold.rawBytes, nameWriteAddress, textEncodeBuffer, length)
			pokegold.rawBytes[nameWriteAddress] = 0x50
			nameWriteAddress++
		}
	}

	{
		utils.CopyBytes(pokegold.rawBytes, 0x35d5, utils.ConvertToPointerWithBank(nameWriteAddress))

		for i := 0; i < TrainerClassNameCount; i++ {
			length := utils.TextEncodeBuffered(textEncodeBuffer, pokegold.Strings.TrainerClassNames[i])
			nameWriteAddress = utils.CopyBytesWithLength(pokegold.rawBytes, nameWriteAddress, textEncodeBuffer, length)
			pokegold.rawBytes[nameWriteAddress] = 0x50
			nameWriteAddress++
		}
	}

	{
		utils.CopyBytes(pokegold.rawBytes, 0x35c3, utils.ConvertToPointerWithBank(nameWriteAddress))
		utils.CopyBytes(pokegold.rawBytes, 0x3667, utils.ConvertToPointer(nameWriteAddress))
		utils.CopyBytes(pokegold.rawBytes, 0x515bf, utils.ConvertToPointer(nameWriteAddress))

		for i := 0; i < PokemonNameCount; i++ {
			length := utils.TextEncodeBuffered(textEncodeBuffer, pokegold.Strings.PokemonNames[i])
			nameWriteAddress = utils.CopyBytesWithLength(pokegold.rawBytes, nameWriteAddress, textEncodeBuffer, length)

			for i := 0; i < 10-length; i++ {
				pokegold.rawBytes[nameWriteAddress] = 0x50
				nameWriteAddress++
			}
		}
	}

	{
		utils.CopyBytes(pokegold.rawBytes, 0x35c6, utils.ConvertToPointerWithBank(nameWriteAddress))

		for i := 0; i < MoveCount; i++ {
			length := utils.TextEncodeBuffered(textEncodeBuffer, pokegold.Strings.MoveNames[i])
			nameWriteAddress = utils.CopyBytesWithLength(pokegold.rawBytes, nameWriteAddress, textEncodeBuffer, length)
			pokegold.rawBytes[nameWriteAddress] = 0x50
			nameWriteAddress++
		}
	}
}
