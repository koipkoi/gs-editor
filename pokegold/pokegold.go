package pokegold

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"sync"

	models "gs-editor/pokegold/models"
)

const (
	ItemCount             = 256
	MoveCount             = 251
	PokemonCount          = 251
	PokemonNameCount      = 256
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
	IsOpen bool

	Filename     string
	rawBytes     []byte
	originalHash string

	Items    []models.Item
	Moves    []models.Move
	Pokemons []models.Pokemon
	Pokedex  []models.Pokedex

	Colors  Colors
	Images  Images
	Strings Strings

	onChangedObserversMutex *sync.Mutex
	onChangedObservers      []func(*Pokegold)
}

type PokegoldConverter interface {
	Read(pokegold *Pokegold)
	Write(pokegold *Pokegold)
}

func NewPokegold() *Pokegold {
	return &Pokegold{
		IsOpen:                  false,
		onChangedObserversMutex: &sync.Mutex{},
		onChangedObservers:      nil,
	}
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

	pokegold.IsOpen = true
	pokegold.Filename = filename
	pokegold.NotifyOnChanged()

	return nil
}

func (pokegold *Pokegold) WriteRom(filename string) error {
	for _, converter := range converters {
		converter.Write(pokegold)
	}

	err := os.WriteFile(filename, pokegold.rawBytes, 0644)
	if err != nil {
		return err
	}

	hash := md5.Sum(pokegold.rawBytes)
	pokegold.originalHash = hex.EncodeToString(hash[:])
	pokegold.Filename = filename
	pokegold.NotifyOnChanged()

	return nil
}

func (pokegold *Pokegold) AddOnChanged(obs func(*Pokegold)) {
	pokegold.onChangedObserversMutex.Lock()
	pokegold.onChangedObservers = append(pokegold.onChangedObservers, obs)
	pokegold.onChangedObserversMutex.Unlock()
}

func (pokegold *Pokegold) NotifyOnChanged() {
	pokegold.onChangedObserversMutex.Lock()
	for _, obs := range pokegold.onChangedObservers {
		obs(pokegold)
	}
	pokegold.onChangedObserversMutex.Unlock()
}
