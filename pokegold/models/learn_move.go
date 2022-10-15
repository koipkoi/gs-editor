package models

type LearnMove struct {
	Level  byte
	MoveNo byte
}

func NewLearnMove(level byte, moveNo byte) *LearnMove {
	return &LearnMove{
		Level:  level,
		MoveNo: moveNo,
	}
}

func NewLearnMovesFromBytes(bytes []byte) []LearnMove {
	var learnMoves []LearnMove
	for i := 0; i < len(bytes); i += 2 {
		newLearnMove := NewLearnMove(bytes[i], bytes[i+1])
		learnMoves = append(learnMoves, *newLearnMove)
	}
	return learnMoves
}

func (learnmove *LearnMove) ToBytes() []byte {
	return []byte{
		learnmove.Level,
		learnmove.MoveNo,
	}
}
