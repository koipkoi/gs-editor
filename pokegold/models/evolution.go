package models

type Evolution struct {
	EvolutionType byte
	PokemonNo     byte
	Level         byte
	ItemNo        byte
	Affection     byte
	BaseStats     byte
}

func NewEvolutionFromBytes(bytes []byte) *Evolution {
	evolution := &Evolution{}
	if len(bytes) > 0 {
		evolution.EvolutionType = bytes[0]
		switch evolution.EvolutionType {
		case 1:
			if len(bytes) == 3 {
				evolution.Level = bytes[1]
				evolution.PokemonNo = bytes[2]
			}
		case 2, 3:
			if len(bytes) == 3 {
				evolution.ItemNo = bytes[1]
				evolution.PokemonNo = bytes[2]
			}
		case 4:
			if len(bytes) == 3 {
				evolution.Affection = bytes[1]
				evolution.PokemonNo = bytes[2]
			}
		case 5:
			if len(bytes) == 4 {
				evolution.Level = bytes[1]
				evolution.BaseStats = bytes[2]
				evolution.PokemonNo = bytes[3]
			}
		}
	}
	return evolution
}

func NewEvolutionsFromBytes(bytes []byte) []Evolution {
	// 바이트 토큰 처리
	var arrays [][]byte
	for i := 0; i < len(bytes); i++ {
		var buffer []byte
		switch bytes[i] {
		case 1, 2, 3, 4:
			buffer = append(buffer, bytes[i], bytes[i+1], bytes[i+2])
			arrays = append(arrays, buffer)
			i += 2
		case 5:
			buffer = append(buffer, bytes[i], bytes[i+1], bytes[i+2], bytes[i+3])
			arrays = append(arrays, buffer)
			i += 3
		}
	}

	// 진화 데이터 파싱
	var evolutions []Evolution
	for _, e := range arrays {
		evolution := NewEvolutionFromBytes(e)
		evolutions = append(evolutions, *evolution)
	}

	return evolutions
}

func (evolution *Evolution) ToBytes() []byte {
	switch evolution.EvolutionType {
	case 1:
		return []byte{
			evolution.EvolutionType,
			evolution.Level,
			evolution.PokemonNo,
		}
	case 2, 3:
		return []byte{
			evolution.EvolutionType,
			evolution.ItemNo,
			evolution.PokemonNo,
		}
	case 4:
		return []byte{
			evolution.EvolutionType,
			evolution.Affection,
			evolution.PokemonNo,
		}
	case 5:
		return []byte{
			evolution.EvolutionType,
			evolution.Level,
			evolution.BaseStats,
			evolution.PokemonNo,
		}
	}
	return []byte{0}
}
