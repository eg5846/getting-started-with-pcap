package utils

import (
	"encoding/binary"
	"unicode/utf16"
)

func ToUtf16(s string) []uint16 {
	return utf16.Encode([]rune(s))
}

func ToBytes(uints []uint16) []byte {
	bytes := []byte{}
	for _, ui := range uints {
		b := make([]byte, 2)
		binary.LittleEndian.PutUint16(b, ui)
		bytes = append(bytes, b[0], b[1])
	}
	return bytes
}
