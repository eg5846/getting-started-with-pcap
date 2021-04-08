# fritzdump
Capture PCAPs from FRITZ!Box

## References
* http://fritz.box/support.lua
* http://fritz.box/html/capture.html
* https://www.heise.de/ratgeber/Paketmitschnitte-der-Fritzbox-automatisch-an-Wireshark-weitergeben-4155867.html
* https://www.ntop.org/ntopng/how-to-use-ntopng-for-realtime-traffic-analysis-on-fritzbox-routers/
* https://github.com/ntop/ntopng/tree/dev/tools
* https://github.com/cardigliano/wireshark-fritzbox

## Requirements
* FRITZ!Box connected to internet
* FRITZ!Box login credentials
* go version go1.15.8

## Build fritzdump
```
go build
```

## Run fritzdump
```
# Usage
$ ./fritzdump -h
Usage of ./fritzdump:
  -i string
    	FRITZ!Box interface (2-0: WAN, 1-lan: LAN, ...) (default "2-0")
  -o string
    	Output file path for PCAP dump (default "/tmp/test.pcap")
  -p string
    	FRITZ!Box login password
  -s string
    	FRITZ!Box url (default "http://fritz.box")
  -u string
    	FRITZ!Box login username (default "dslf-config")

$ ./fritzdump -p 123fritz -i 2-1
```