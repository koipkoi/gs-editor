package pokegold

import (
	"gs-editor/pokegold/utils"
)

var (
	imageAddrs = [...][2]int{
		{0x485e2, 0x4bfff},
		{0x54000, 0x57fff},
		{0x58000, 0x5bfce},
		{0x5c000, 0x5ffff},
		{0x60000, 0x63fc8},
		{0x64000, 0x67fbd},
		{0x68000, 0x6bffc},
		{0x6c000, 0x6ffff},
		{0x70000, 0x73fff},
		{0x74000, 0x77fde},
		{0x78000, 0x7bfc5},
		{0x7c09c, 0x7ffff},
		{0x800c6, 0x83fd9},
	}
)

type Images struct {
	PokemonImages         [][]byte
	PokemonBacksideImages [][]byte
	TrainerImages         [][]byte
	UnownImages           [][]byte
	UnownBacksideImages   [][]byte
}

type ImagesConverter struct{}

func (*ImagesConverter) Read(pokegold *Pokegold) {
	// pokemon
	{
		pokegold.Images.PokemonImages = nil
		pokegold.Images.PokemonBacksideImages = nil

		for i := 0; i < PokemonCount; i++ {
			if i != UnownNo {
				var bank byte
				var bytes []byte
				var address int

				bank = utils.FixBank(pokegold.rawBytes[0x48000+(i*6)])
				address = utils.ConvertToAddress(bank, utils.SliceBytes(pokegold.rawBytes, 0x48001+(i*6), 2))
				bytes = utils.LZDecompress(pokegold.rawBytes[address : address+ReadImageBufferSize])
				pokegold.Images.PokemonImages = append(pokegold.Images.PokemonImages, bytes)

				bank = utils.FixBank(pokegold.rawBytes[0x48000+(i*6)+3])
				address = utils.ConvertToAddress(bank, utils.SliceBytes(pokegold.rawBytes, 0x48001+(i*6)+3, 2))
				bytes = utils.LZDecompress(pokegold.rawBytes[address : address+ReadImageBufferSize])
				pokegold.Images.PokemonBacksideImages = append(pokegold.Images.PokemonBacksideImages, bytes)
			}
		}
	}

	// trainer
	{
		pokegold.Images.TrainerImages = nil

		for i := 0; i < TrainerImageCount; i++ {
			bank := utils.FixBank(pokegold.rawBytes[0x80000+(i*3)])
			address := utils.ConvertToAddress(bank, utils.SliceBytes(pokegold.rawBytes, 0x80001+(i*3), 2))
			bytes := utils.LZDecompress(utils.SliceBytes(pokegold.rawBytes, address, ReadImageBufferSize))
			pokegold.Images.TrainerImages = append(pokegold.Images.TrainerImages, bytes)
		}
	}

	// unowns
	{
		pokegold.Images.UnownImages = nil
		pokegold.Images.UnownBacksideImages = nil

		for i := 0; i < UnownCount; i++ {
			address := utils.ConvertToAddress(utils.FixBank(pokegold.rawBytes[0x7c000+(i*6)]), utils.SliceBytes(pokegold.rawBytes, 0x7c001+(i*6), 2))
			bytes := utils.LZDecompress(utils.SliceBytes(pokegold.rawBytes, address, ReadImageBufferSize))
			pokegold.Images.UnownImages = append(pokegold.Images.UnownImages, bytes)

			backsideAddress := utils.ConvertToAddress(utils.FixBank(pokegold.rawBytes[0x7c003+(i*6)]), utils.SliceBytes(pokegold.rawBytes, 0x7c004+(i*6), 2))
			backsideBytes := utils.LZDecompress(utils.SliceBytes(pokegold.rawBytes, backsideAddress, ReadImageBufferSize))
			pokegold.Images.UnownBacksideImages = append(pokegold.Images.UnownBacksideImages, backsideBytes)
		}
	}
}

func (*ImagesConverter) Write(pokegold *Pokegold) {
	// todo 추가
}
