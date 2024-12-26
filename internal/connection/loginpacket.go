package connection

import (
	"fmt"
)

type LoginPacket struct {
	Packet
	RawNPSPacket
	CustomerId uint
	SessionKey string
}

func (p LoginPacket) String() string {
	return fmt.Sprintf("LoginPacket{RawNPSPacket: %v, CustomerId: %d, SessionKey: %X}", p.RawNPSPacket, p.CustomerId, p.SessionKey)
}

func decodeSessionKey(data []byte) string {
	for i, b := range data {
		if b == 0 {
			return string(data[:i])
		}
	}
	return string(data)
}

func (p LoginPacket) MarshalBinary() ([]byte, error) {
	header, err := p.RawNPSPacket.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return append(header, []byte{byte(p.CustomerId >> 24), byte(p.CustomerId >> 16), byte(p.CustomerId >> 8), byte(p.CustomerId)}...), nil
}

func (p LoginPacket) UnmarshalBinary(data []byte) error {
	if len(data) < 8 {
		return fmt.Errorf("data too short")
	}

	if err := p.RawNPSPacket.UnmarshalBinary(data); err != nil {
		return err
	}

	p.CustomerId = uint(data[4])<<24 | uint(data[5])<<16 | uint(data[6])<<8 | uint(data[7])
	p.SessionKey = decodeSessionKey(data[8:])

	fmt.Println("Unmarshaled LoginPacket:", p)

	return nil
}
