package pokegold

import (
	"crypto/md5"
	"encoding/hex"
	"os"

	models "gs-editor/pokegold/models"
)

const (
	ItemCount             = 256
	MoveCount             = 251
	PokemonCount          = 251
	UnownNo               = 200
	TrainerImageCount     = 66
	UnownCount            = 26
	TrainerClassNameCount = 67
	MoveTypeNameCount     = 28
	ReadImageBufferSize   = 4096
)

var (
	converters = [...]PokegoldConverter{
		&ItemsConverter{},
		&MovesConverter{},
		&PokemonsConverter{},
		&PokedexConverter{},
		&ColorsConverter{},
		&ImagesConverter{},
		&StringsConverter{},
	}
)

type Pokegold struct {
	isOpen bool `default:"false"`

	rawBytes     []byte
	originalHash string

	Items    []models.Item
	Moves    []models.Move
	Pokemons []models.Pokemon
	Pokedex  []models.Pokedex

	Colors  Colors
	Images  Images
	Strings Strings
}

type PokegoldConverter interface {
	Read(pokegold *Pokegold)
	Write(pokegold *Pokegold)
}

func (pokegold *Pokegold) IsOpen() bool {
	return pokegold.isOpen
}

func (pokegold *Pokegold) ReadRom(filename string) error {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	hash := md5.Sum(bytes)
	pokegold.originalHash = hex.EncodeToString(hash[:])
	pokegold.rawBytes = bytes

	for _, converter := range converters {
		converter.Read(pokegold)
	}

	pokegold.isOpen = true
	return nil
}

func (pokegold *Pokegold) WriteRom(filename string) error {
	for _, converter := range converters {
		converter.Write(pokegold)
	}

	hash := md5.Sum(pokegold.rawBytes)
	pokegold.originalHash = hex.EncodeToString(hash[:])

	return os.WriteFile(filename, pokegold.rawBytes, 0644)
}
