package pcapfile

import (
	"fmt"

	"github.com/google/gopacket"
	// gopacket godocs recommends to import layers always
	"github.com/google/gopacket/layers"
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

func (r *reader) LinkType() layers.LinkType {
	return r.handle.LinkType()
}

func (r *reader) Packets() (chan gopacket.Packet, error) {
	if r.LinkType() == layers.LinkTypeEthernet {
		return gopacket.NewPacketSource(r.handle, r.handle.LinkType()).Packets(), nil
	}

	// TODO: Improve this
	if r.LinkType() == 12 {
		return gopacket.NewPacketSource(r.handle, layers.LinkTypeIPv4).Packets(), nil
	}

	return nil, fmt.Errorf("Unsupported LinkType: %s (%d)", r.LinkType(), r.LinkType())
}
