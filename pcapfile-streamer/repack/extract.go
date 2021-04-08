package repack

import (
	"errors"

	"github.com/google/gopacket"
	// gopacket godocs recommends to import layers always
	_ "github.com/google/gopacket/layers"
)

var (
	packetWithoutLinkLayerErr    error = errors.New("Packet without link layer")
	packetWithoutNetworkLayerErr error = errors.New("Packet without network layer")
)

// Extracts layer 2 (link layer) payload from given packet
func ExtractLinkLayerPayload(packet gopacket.Packet) ([]byte, gopacket.LayerType, error) {
	networkLayerType, err := extractNetworkLayerLinkType(packet)
	if err != nil {
		return nil, 0, err
	}

	if linkLayer := packet.LinkLayer(); linkLayer != nil {
		return linkLayer.LayerPayload(), networkLayerType, nil
	}

	return nil, 0, packetWithoutLinkLayerErr
}

func extractNetworkLayerLinkType(packet gopacket.Packet) (gopacket.LayerType, error) {
	if networkLayer := packet.NetworkLayer(); networkLayer != nil {
		return networkLayer.LayerType(), nil
	}

	return 0, packetWithoutNetworkLayerErr
}
