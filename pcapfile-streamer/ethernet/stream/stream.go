package stream

import (
	"time"

	"github.com/google/gopacket/pcap"
)

type stream struct {
	handle *pcap.Handle
}

func New(device string) (*stream, error) {
	handle, err := pcap.OpenLive(device, 1024, false, 30*time.Second)
	if err != nil {
		return nil, err
	}

	return &stream{handle}, nil
}

func (s *stream) Close() {
	s.handle.Close()
}

func (s *stream) WritePacketData(data []byte) error {
	return s.handle.WritePacketData(data)
}
