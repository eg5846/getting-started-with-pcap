package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/gopacket/layers"
	"github.com/mdlayher/ethernet"

	"github.com/eg5846/getting-started-with-pcap/pcapfile-streamer/ethernet/stream"
	"github.com/eg5846/getting-started-with-pcap/pcapfile-streamer/pcapfile"
	"github.com/eg5846/getting-started-with-pcap/pcapfile-streamer/repack"
)

var opts struct {
	pcapInFile      string
	streamingDevice string
	vlanId          uint
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.SetOutput(os.Stdout)

	flag.StringVar(&opts.pcapInFile, "r", "", "Path of PCAP input file")
	flag.StringVar(&opts.streamingDevice, "i", "eth0", "Name of ethernet device to stream to")
	flag.UintVar(&opts.vlanId, "t", 0x00a, "VLAN identifier for raw ethernet stream")

	flag.Parse()
}

func main() {
	destination := []byte{0xbb, 0xbb, 0xbb, 0xbb, 0xbb, 0xbb}
	source := []byte{0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa}

	log.Printf("[%s] Starting with VLAN identifier %d ...", opts.pcapInFile, opts.vlanId)
	vlan := &ethernet.VLAN{
		Priority:     ethernet.PriorityBestEffort,
		DropEligible: false,
		ID:           uint16(opts.vlanId),
	}

	log.Printf("[%s] Opening stream on %s ...", opts.pcapInFile, opts.streamingDevice)
	stream, err := stream.New(opts.streamingDevice)
	if err != nil {
		log.Fatalf("[%s] Opening stream on %s failed: %s", opts.pcapInFile, opts.streamingDevice, err)
	}
	defer stream.Close()

	log.Printf("[%s] Opening PCAP input file ...", opts.pcapInFile)
	reader, err := pcapfile.NewReader(opts.pcapInFile)
	if err != nil {
		log.Fatalf("[%s] Opening PCAP input file failed: %s", opts.pcapInFile, err)
	}
	defer reader.Close()

	for packet := range reader.Packets() {
		ipContent, ipLayerType, err := repack.ExtractLinkLayerPayload(packet)
		if err != nil {
			log.Printf("[%s] Parsing IP content from packet failed: %s", opts.pcapInFile, err)
			continue
		}

		md5sum := fmt.Sprintf("%x", md5.Sum(ipContent))
		log.Printf("[%s] IP content md5sum: %s, IP layer type: %s (%d), size: %d Bytes", opts.pcapInFile, md5sum, ipLayerType, ipLayerType, len(ipContent))
		//log.Printf("[%s] %#v", opts.pcapInFile, ipContent)

		switch ipLayerType {
		case layers.LayerTypeIPv4:
			ethernetPacket, err := repack.EncodeEthernetIPv4Packet(destination, source, vlan, ipContent)
			if err != nil {
				log.Printf("[%s] Encoding ethernet IPv4 packet failed: %s", opts.pcapInFile, err)
				continue
			}
			log.Printf("[%s] Ethernet IPv4 packet size: %d Bytes", opts.pcapInFile, len(ethernetPacket))
			//log.Printf("[%s] %#v", opts.pcapInFile, ethernetPacket)

			if err := stream.WritePacketData(ethernetPacket); err != nil {
				log.Printf("[%s] Streaming ethernet IPv4 packet to %s failed: %s", opts.pcapInFile, opts.streamingDevice, err)
				continue
			}

		case layers.LayerTypeIPv6:
			ethernetPacket, err := repack.EncodeEthernetIPv6Packet(destination, source, vlan, ipContent)
			if err != nil {
				log.Printf("[%s] Encoding ethernet IPv6 packet failed: %s", opts.pcapInFile, err)
				continue
			}
			log.Printf("[%s] Ethernet IPv6 packet size: %d Bytes", opts.pcapInFile, len(ethernetPacket))
			//log.Printf("[%s] %#v", opts.pcapInFile, ethernetPacket)

			if err := stream.WritePacketData(ethernetPacket); err != nil {
				log.Printf("[%s] Streaming ethernet IPv6 packet to %s failed: %s", opts.pcapInFile, opts.streamingDevice, err)
				continue
			}
		}
	}
}
