package main

import (
	"flag"
	"log"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var opts struct {
	pcapInFile string
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.SetOutput(os.Stdout)

	flag.StringVar(&opts.pcapInFile, "i", "", "Path of PCAP input file")
	flag.Parse()
}

func logHandleInfos(handle *pcap.Handle) {
	log.Printf("Link type:  %s", handle.LinkType())
	log.Printf("Snap len:   %d (max length of captured packets, in octets?)", handle.SnapLen())
	log.Printf("Resolution: %s", handle.Resolution().ToDuration())
}

// Statistics aren't available from savefiles
// func logHandleStats(handle *pcap.Handle) error {
// 	stats, err := handle.Stats()
// 	if err != nil {
// 		return err
// 	}

// 	log.Printf("Packets received:   %d", stats.PacketsReceived)
// 	log.Printf("Packets dropped:    %d", stats.PacketsDropped)
// 	log.Printf("Packets if dropped: %d", stats.PacketsIfDropped)

// 	return nil
// }

// Packet can read only once from chan PacketSource
func readPackets(handle *pcap.Handle) []gopacket.Packet {
	packets := []gopacket.Packet{}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		packets = append(packets, packet)
	}
	return packets
}

func logPackets(packets []gopacket.Packet) {
	for i, packet := range packets {
		log.Printf("[%d] %s", i, packet)
	}
}

// TODO: Check if logic is ok
func checkOrderOfPcapTimestamps(packets []gopacket.Packet) {
	warnings := 0
	for i := 1; i < len(packets); i++ {
		currentPacketTimestamp := packets[i-1].Metadata().Timestamp
		nextPacketTimestamp := packets[i].Metadata().Timestamp
		if nextPacketTimestamp.Before(currentPacketTimestamp) {
			log.Printf("[WARNING] Timestamp of packet %d is not in order", i)
			warnings += 1
		}
	}
	log.Printf("Checking order of PCAP timestamps, result: %d warnings", warnings)
}

func main() {
	log.Printf("Reading from '%s' ...", opts.pcapInFile)

	handle, err := pcap.OpenOffline(opts.pcapInFile)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	logHandleInfos(handle)

	// Statistics aren't available from savefiles
	// if err := logHandleStats(handle); err != nil {
	// 	log.Fatal(err)
	// }

	packets := readPackets(handle)

	log.Printf("Packets: %d", len(packets))

	logPackets(packets)

	checkOrderOfPcapTimestamps(packets)
}
