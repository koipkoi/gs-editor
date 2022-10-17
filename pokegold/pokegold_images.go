package pokegold

import (
	"gs-editor/pokegold/utils"
)

var (
	imageAddrs = [...][2]int{
		{0x0485e2, 0x04bfff},
		{0x054000, 0x057fff},
		{0x058000, 0x05bfff},
		{0x05c000, 0x05ffff},
		{0x060000, 0x063fff},
		{0x064000, 0x067fff},
		{0x068000, 0x06bfff},
		{0x06c000, 0x06ffff},
		{0x070000, 0x073fff},
		{0x074000, 0x077fff},
		{0x078000, 0x07bfff},
		{0x07c09c, 0x07ffff},
		{0x0800c6, 0x083fff},

		// 추가 공간
		{0x088000, 0x08bfff},
		{0x09c000, 0x09ffff},
		{0x0a0000, 0x0a3fff},
		{0x0a4000, 0x0a7fff},
		{0x0b0000, 0x0b3fff},
		{0x0b4000, 0x0b7fff},
		{0x0bc000, 0x0bffff},
		{0x0d0000, 0x0d3fff},
		{0x0d4000, 0x0d7fff},
		{0x160000, 0x163fff},
		{0x18c000, 0x18ffff},
		{0x1a8000, 0x1abfff},
		{0x1ac000, 0x1affff},
		{0x1bc000, 0x1bffff},
		{0x1cc000, 0x1cffff},
		{0x1d0000, 0x1d3fff},
		{0x1d4000, 0x1d7fff},
		{0x1f0000, 0x1f3fff},
		{0x1f4000, 0x1f7fff},
		{0x1f8000, 0x1fbfff},
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

				bank = utils.DecodeBank(pokegold.rawBytes[0x48000+(i*6)])
				address = utils.ConvertToAddress(bank, utils.SliceBytes(pokegold.rawBytes, 0x48001+(i*6), 2))
				bytes = utils.LZDecompress(pokegold.rawBytes[address : address+ReadImageBufferSize])
				pokegold.Images.PokemonImages = append(pokegold.Images.PokemonImages, bytes)

				bank = utils.DecodeBank(pokegold.rawBytes[0x48000+(i*6)+3])
				address = utils.ConvertToAddress(bank, utils.SliceBytes(pokegold.rawBytes, 0x48001+(i*6)+3, 2))
				bytes = utils.LZDecompress(pokegold.rawBytes[address : address+ReadImageBufferSize])
				pokegold.Images.PokemonBacksideImages = append(pokegold.Images.PokemonBacksideImages, bytes)
			} else {
				pokegold.Images.PokemonImages = append(pokegold.Images.PokemonImages, nil)
				pokegold.Images.PokemonBacksideImages = append(pokegold.Images.PokemonBacksideImages, nil)
			}
		}
	}

	// trainer
	{
		pokegold.Images.TrainerImages = nil

		for i := 0; i < TrainerImageCount; i++ {
			bank := utils.DecodeBank(pokegold.rawBytes[0x80000+(i*3)])
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
			address := utils.ConvertToAddress(utils.DecodeBank(pokegold.rawBytes[0x7c000+(i*6)]), utils.SliceBytes(pokegold.rawBytes, 0x7c001+(i*6), 2))
			bytes := utils.LZDecompress(utils.SliceBytes(pokegold.rawBytes, address, ReadImageBufferSize))
			pokegold.Images.UnownImages = append(pokegold.Images.UnownImages, bytes)

			backsideAddress := utils.ConvertToAddress(utils.DecodeBank(pokegold.rawBytes[0x7c003+(i*6)]), utils.SliceBytes(pokegold.rawBytes, 0x7c004+(i*6), 2))
			backsideBytes := utils.LZDecompress(utils.SliceBytes(pokegold.rawBytes, backsideAddress, ReadImageBufferSize))
			pokegold.Images.UnownBacksideImages = append(pokegold.Images.UnownBacksideImages, backsideBytes)
		}
	}
}

func (*ImagesConverter) Write(pokegold *Pokegold) {
	var indexes []int

	for _, e := range imageAddrs {
		indexes = append(indexes, 0)
		size := e[1] - e[0] + 1
		utils.FillBytes(pokegold.rawBytes, 0, e[0], size)
	}

	for i := 0; i < PokemonCount; i++ {
		if i != UnownNo {
			e := pokegold.Images.PokemonImages[i]
			pointerAddr := 0x48000 + (i * 6)
			writeImage(pokegold.rawBytes, pointerAddr, e, indexes)
		}
	}

	for i := 0; i < PokemonCount; i++ {
		if i != UnownNo {
			e := pokegold.Images.PokemonBacksideImages[i]
			pointerAddr := 0x48003 + (i * 6)
			writeImage(pokegold.rawBytes, pointerAddr, e, indexes)
		}
	}

	for i, e := range pokegold.Images.TrainerImages {
		pointerAddr := 0x80000 + (i * 3)
		writeImage(pokegold.rawBytes, pointerAddr, e, indexes)
	}

	for i, e := range pokegold.Images.UnownImages {
		pointerAddr := 0x7c000 + (i * 6)
		writeImage(pokegold.rawBytes, pointerAddr, e, indexes)
	}

	for i, e := range pokegold.Images.UnownBacksideImages {
		pointerAddr := 0x7c003 + (i * 6)
		writeImage(pokegold.rawBytes, pointerAddr, e, indexes)
	}
}

func writeImage(rom []byte, pointerAddr int, bytes []byte, indexes []int) {
	compressed := utils.LZCompress(bytes)
	size := len(compressed)

	for indexAddr, e := range imageAddrs {
		newAddr := e[0] + indexes[indexAddr]
		if newAddr+size < e[1] {
			pointer := utils.ConvertToPointer(newAddr)
			utils.CopyBytes(rom, pointerAddr, []byte{
				utils.EncodeBank(utils.ConvertToBank(newAddr)),
				pointer[0],
				pointer[1],
			})
			utils.CopyBytes(rom, newAddr, compressed)
			indexes[indexAddr] += size
			break
		}
	}
}
