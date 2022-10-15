package models

type Move struct {
	Animation    byte
	Effect       byte
	Power        byte
	MoveType     byte
	Accuracy     byte
	PP           byte
	EffectChance byte
}

func NewMoveFromBytes(bytes []byte) *Move {
	return &Move{
		Animation:    bytes[0],
		Effect:       bytes[1],
		Power:        bytes[2],
		MoveType:     bytes[3],
		Accuracy:     bytes[4],
		PP:           bytes[5],
		EffectChance: bytes[6],
	}
}

func (move *Move) ToBytes() []byte {
	return []byte{
		move.Animation,
		move.Effect,
		move.Power,
		move.MoveType,
		move.Accuracy,
		move.PP,
		move.EffectChance,
	}
}
