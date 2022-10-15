package models

type Item struct {
	Price     int
	Effect    byte
	Parameter byte
	Property  byte
	Pocket    byte
	Menu      byte
}

func NewItemFromBytes(bytes []byte) *Item {
	return &Item{
		Price:     int(bytes[0]) + (int(bytes[1]) << 8),
		Effect:    bytes[2],
		Parameter: bytes[3],
		Property:  bytes[4],
		Pocket:    bytes[5],
		Menu:      bytes[6],
	}
}

func (item *Item) ToBytes() []byte {
	return []byte{
		byte(item.Price & 0x00ff),
		byte((item.Price & 0xff00) >> 8),
		item.Effect,
		item.Parameter,
		item.Property,
		item.Pocket,
		item.Menu,
	}
}
