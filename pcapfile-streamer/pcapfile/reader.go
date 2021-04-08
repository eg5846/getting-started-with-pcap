package pcapfile

import (
	"github.com/google/gopacket"
	// gopacket godocs recommends to import layers always
	_ "github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type reader struct {
	handle *pcap.Handle
}

func NewReader(file string) (*reader, error) {
	handle, err := pcap.OpenOffline(file)
	if err != nil {
		return nil, err
	}

	return &reader{handle}, nil
}

func (r *reader) Close() {
	r.handle.Close()
}

func (r *reader) Packets() chan gopacket.Packet {
	return gopacket.NewPacketSource(r.handle, r.handle.LinkType()).Packets()
}
