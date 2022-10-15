package models

var (
	bits = [...]byte{
		0b00000001,
		0b00000010,
		0b00000100,
		0b00001000,
		0b00010000,
		0b00100000,
		0b01000000,
		0b10000000,
	}
	reverseBits = [...]byte{
		0b10000000,
		0b01000000,
		0b00100000,
		0b00010000,
		0b00001000,
		0b00000100,
		0b00000010,
		0b00000001,
	}
)

type Pokemon struct {
	No            byte
	HP            byte
	Attack        byte
	Defence       byte
	Speed         byte
	SpAttack      byte
	SpDefence     byte
	Type1         byte
	Type2         byte
	CatchRate     byte
	Exp           byte
	Item1         byte
	Item2         byte
	GenderRate    byte
	Unk1          byte
	EggType       byte
	Unk2          byte
	ImageSizeType byte
	Padding1      byte
	Padding2      byte
	Padding3      byte
	Padding4      byte
	GrowthRate    byte
	EggGroup      byte

	TMHMs [64]bool

	Evolutions []Evolution
	LearnMoves []LearnMove
}

func NewPokemonFromBytes(bytes []byte) *Pokemon {
	newPokemon := &Pokemon{
		No:            bytes[0],
		HP:            bytes[1],
		Attack:        bytes[2],
		Defence:       bytes[3],
		Speed:         bytes[4],
		SpAttack:      bytes[5],
		SpDefence:     bytes[6],
		Type1:         bytes[7],
		Type2:         bytes[8],
		CatchRate:     bytes[9],
		Exp:           bytes[10],
		Item1:         bytes[11],
		Item2:         bytes[12],
		GenderRate:    bytes[13],
		Unk1:          bytes[14],
		EggType:       bytes[15],
		Unk2:          bytes[16],
		ImageSizeType: bytes[17],
		Padding1:      bytes[18],
		Padding2:      bytes[19],
		Padding3:      bytes[20],
		Padding4:      bytes[21],
		GrowthRate:    bytes[22],
		EggGroup:      bytes[23],
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			index := (i * 8) + j
			newPokemon.TMHMs[index] = (bytes[24+i] & reverseBits[j]) != 0
		}
	}

	return newPokemon
}

func (pokemon *Pokemon) ToBytes() []byte {
	bytes := []byte{
		pokemon.No,
		pokemon.HP,
		pokemon.Attack,
		pokemon.Defence,
		pokemon.Speed,
		pokemon.SpAttack,
		pokemon.SpDefence,
		pokemon.Type1,
		pokemon.Type2,
		pokemon.CatchRate,
		pokemon.Exp,
		pokemon.Item1,
		pokemon.Item2,
		pokemon.GenderRate,
		pokemon.Unk1,
		pokemon.EggType,
		pokemon.Unk2,
		pokemon.ImageSizeType,
		pokemon.Padding1,
		pokemon.Padding2,
		pokemon.Padding3,
		pokemon.Padding4,
		pokemon.GrowthRate,
		pokemon.EggGroup,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			index := (i * 8) + j
			if pokemon.TMHMs[index] {
				bytes[24+i] = bytes[24+i] | reverseBits[j]
			}
		}
	}

	return bytes
}
