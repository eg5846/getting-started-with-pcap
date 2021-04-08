package repack

import (
	"net"

	"github.com/mdlayher/ethernet"
)

func EncodeEthernetIPv4Packet(destination net.HardwareAddr, source net.HardwareAddr, vlan *ethernet.VLAN, payload []byte) ([]byte, error) {
	frame := &ethernet.Frame{
		Destination: destination,
		Source:      source,
		ServiceVLAN: nil,
		VLAN:        vlan,
		EtherType:   ethernet.EtherTypeIPv4,
		Payload:     payload,
	}
	return frame.MarshalBinary()
}

func EncodeEthernetIPv6Packet(destination net.HardwareAddr, source net.HardwareAddr, vlan *ethernet.VLAN, payload []byte) ([]byte, error) {
	frame := &ethernet.Frame{
		Destination: destination,
		Source:      source,
		ServiceVLAN: nil,
		VLAN:        vlan,
		EtherType:   ethernet.EtherTypeIPv6,
		Payload:     payload,
	}
	return frame.MarshalBinary()
}
