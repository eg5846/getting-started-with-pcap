package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/eg5846/getting-started-with-pcap/fritzdump/fritzbox"
)

var opts struct {
	url      string
	username string
	password string
	iface    string
	outFile  string
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.SetOutput(os.Stdout)

	flag.StringVar(&opts.url, "s", "http://fritz.box", "FRITZ!Box url")
	flag.StringVar(&opts.username, "u", "dslf-config", "FRITZ!Box login username")
	flag.StringVar(&opts.password, "p", "", "FRITZ!Box login password")
	flag.StringVar(&opts.iface, "i", "2-0", "FRITZ!Box interface (2-0: WAN, 1-lan: LAN, ...)")
	flag.StringVar(&opts.outFile, "o", "/tmp/test.pcap", "Output file path for PCAP dump")

	flag.Parse()
}

func main() {
	log.Printf("Creating device for %s ...", opts.url)
	device, err := fritzbox.NewDevice(opts.url, opts.username, opts.password, opts.iface)
	if err != nil {
		log.Fatalf("Creating device for %s failed: %s", opts.url, err)
	}

	log.Printf("Authenticating at device with username %s ...", device.Username())
	if err := device.Authenticate(); err != nil {
		log.Fatalf("Authenticating at device with username %s failed: %s", device.Username(), err)
	}

	log.Printf("Capturing from interface %s ...", device.Interface())
	body, err := device.StartCapturing()
	if err != nil {
		log.Fatalf("Capturing from interface %s failed: %s", device.Interface(), err)
	}
	defer body.Close()

	log.Printf("Dumping to %s ...", opts.outFile)
	f, err := os.Create(opts.outFile)
	if err != nil {
		log.Fatalf("Dumping to %s failed: %s", opts.outFile, err)
	}
	defer f.Close()

	n, err := io.Copy(f, body)
	if err != nil {
		log.Fatalf("Dumping to %s failed: %s", opts.outFile, err)
	}
	log.Printf("%d Bytes dumped to %s", n, opts.outFile)
}
