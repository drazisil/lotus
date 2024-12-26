package connection

import (
	"fmt"
)

type Packet interface {
	String() string
	MarshalBinary() ([]byte, error)
	UnmarshalBinary([]byte) error
}

type RawNPSPacket struct {
	Packet
	MsgId uint16
	Length uint16
	Data   []byte
}

func (p RawNPSPacket) String() string {
	return fmt.Sprintf("Header: %X, Length: %v, Data: %X", p.MsgId, p.Length, p.Data)
}

func (p RawNPSPacket) MarshalBinary() ([]byte, error) {
	return []byte{byte(p.MsgId >> 8), byte(p.MsgId), byte(p.Length >> 8), byte(p.Length)}, nil
}

func (p RawNPSPacket) UnmarshalBinary(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("data too short")
	}
	
	p.MsgId = uint16(data[0])<<8 | uint16(data[1])
	p.Length = uint16(data[2])<<8 | uint16(data[3])
	p.Data = data[4:]

	if len(p.Data) + 4 != int(p.Length) {
		return fmt.Errorf("data length mismatch: expected %d, got %d", p.Length, len(p.Data))
	}

	fmt.Println("Unmarshaled RawNPSPacket:", p)

	return nil
}

