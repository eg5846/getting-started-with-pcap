package main

import (
	"flag"
	"log"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

var opts struct {
	pcapInFile  string
	pcapOutFile string
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.SetOutput(os.Stdout)

	flag.StringVar(&opts.pcapInFile, "i", "", "Path of PCAP input file")
	flag.StringVar(&opts.pcapOutFile, "o", "", "Path of PCAP output file (raw)")
	flag.Parse()
}

func main() {
	log.Printf("Opening PCAP input file '%s' ...", opts.pcapInFile)
	handle, err := pcap.OpenOffline(opts.pcapInFile)
	if err != nil {
		log.Fatalf("Opening PCAP input file '%s' failed: %s", opts.pcapInFile, err)
	}
	defer handle.Close()

	log.Printf("LinkType: %s (%d)", handle.LinkType(), handle.LinkType())

	log.Printf("Opening PCAP output file '%s' ...", opts.pcapOutFile)
	f, err := os.Create(opts.pcapOutFile)
	if err != nil {
		log.Fatalf("Opening PCAP output file '%s' failed: %s", opts.pcapOutFile, err)
	}
	defer f.Close()

	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(1024, layers.LinkTypeRaw)

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		layerType := packet.NetworkLayer().LayerType()
		log.Printf("packet's network layer type: %s (%d)", layerType, layerType)
		captureInfo := packet.Metadata().CaptureInfo
		payload := packet.LinkLayer().LayerPayload()
		captureInfo.CaptureLength = len(payload)
		if err := w.WritePacket(packet.Metadata().CaptureInfo, payload); err != nil {
			log.Fatal(err)
		}
	}
}
