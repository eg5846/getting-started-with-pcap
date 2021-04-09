package repack

import (
	"errors"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

var (
	packetContentContainsUnsupportedLayerErr error = errors.New("Packet content contains unsupported layer")
)

// TODO: Try faster, but does not work
// See: https://pkg.go.dev/github.com/google/gopacket#DecodingLayerParser
func ParseIpContent(packetContent []byte, first gopacket.LayerType) ([]byte, gopacket.LayerType, error) {
	var ipv4 layers.IPv4
	var ipv6 layers.IPv6
	// TODO: Add other protocols?

	parser := gopacket.NewDecodingLayerParser(first, &ipv4, &ipv6)
	decoded := []gopacket.LayerType{}
	if err := parser.DecodeLayers(packetContent, &decoded); err != nil {
		return nil, 0, err
	}

	log.Printf("%#v", decoded)

	for _, layerType := range decoded {
		switch layerType {
		case layers.LayerTypeIPv4:
			return ipv4.Contents, layers.LayerTypeIPv4, nil
		case layers.LayerTypeIPv6:
			return ipv6.Contents, layers.LayerTypeIPv6, nil
		}
	}

	return nil, 0, packetContentContainsUnsupportedLayerErr
}
