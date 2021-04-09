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

## About FRITZ!Box interfaces
Open http://fritz.box/support.lua and navigate to `Paketmitschnitte`.  
Open browser's developer tools and inspect `Start` buttons value.  

|Interface|Description|
|---------|-----------|
|2-0|WAN, but not found with inspect, captures external and internal IPs, but no ethernet layer|
|2-1|Internetverbindung?|
|3-17|Schnittstelle 0 ('internet')?|
|3-18|Schnittstelle 1 ('mstv')?|
|3-0|Routing-Schnittstelle?|
|1-eoam|eoam?|
|1-wifi0|wifi0?|
|1-ing0|ing0?|
|1-ath0|ath0?|
|1-lan|lan?|
|1-eth0|eth0|
|1-eth1|eth1|
|1-eth2|eth2|
|1-eth3|eth3|
|1-miireg|miireg?|
|1-wifi1|wifi1?|
|1-ptm0|ptm0?|
|4-133|AP (2.4 + 5 GHz, ath0) - Schnittstelle 1?|
|4-128|WLAN Management Traffic - Schnittstelle 0?|
|5-161|usb1?|
|5-162|usb2?|
