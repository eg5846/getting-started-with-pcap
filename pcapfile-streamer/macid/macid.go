package macid

import (
	"crypto/rand"
	"fmt"
)

type macId struct {
	destination []byte
	source      []byte
}

func New() (*macId, error) {
	destination, err := generateRandomMAC()
	if err != nil {
		return nil, err
	}

	source, err := generateRandomMAC()
	if err != nil {
		return nil, err
	}

	return &macId{destination, source}, nil
}

func (m *macId) Destination() []byte {
	return m.destination
}

func (m *macId) Source() []byte {
	return m.source
}

func (m *macId) Id() []byte {
	b := []byte{}
	b = append(b, m.destination...)
	return append(b, m.source...)
}

func (m *macId) IdString() string {
	return fmt.Sprintf("%x", m.Id())
}

func generateRandomMAC() ([]byte, error) {
	for {
		mac := make([]byte, 6)
		if _, err := rand.Read(mac); err != nil {
			return nil, err
		}
		if isValidMAC(mac) {
			return mac, nil
		}
	}
}

func isValidMAC(mac []byte) bool {
	s := fmt.Sprintf("%x", mac)
	switch s {
	case "000000000000":
		return false
	case "ffffffffffff":
		return false
	default:
		return true
	}
}
