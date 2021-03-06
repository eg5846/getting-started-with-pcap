# pcapfile-streamer
Simple tool to stream content of PCAP file as raw ethernet frames (layer 2) on ethernet device

## References
* https://github.com/google/gopacket
* https://www.devdungeon.com/content/packet-capture-injection-and-analysis-gopacket
* https://en.wikipedia.org/wiki/IEEE_802.1Q
* https://de.wikipedia.org/wiki/IEEE_802.1Q
* https://tools.ietf.org/html/rfc1042

## Requirements
* go version go1.15.8
* Linux syste with root access

## Build pcapfile-streamer
```
go build
```

## VLAN tag (IEEE 802.1Q)
See references above ...

### 802.1Q tag format
32-bit field between the source MAC address and the EtherType fileds of the original frame
```
^16 bits^3 bits^1 bit^12 bits^
|TPID 0x8100|PCP|DEI|VID|

TPID: Tag protocol identifier, always 0x8100 (16 bits)
TCI:  Tag control information, contains PCP, DEI, and VID (in total 16 bits)
PCP:  Priority code point, prioritize different classes of traffic (3 bits)
DEI:  Drop eligible indicator, see references (1 bit)
VID:  VLAN identifier (12 bits), range 0x000 to 0xFFF => max 4094 VLANs (0x000 and 0xFFF are reserved)
```

## Run pcapfile-streamer
```
# Usage
./pcapfile-streamer -h

# Create virtual patch cable
sudo ip link add veth0 type veth peer name veth1
sudo ip link set veth0 up
sudo ip link set veth1 up

# Try with higher mtu to avoid `Message to long` error
sudo ip link set veth0 mtu 1522
sudo ip link set veth1 mtu 1522

# Capture on veth1
sudo tshark -i veth1
sudo tshark -i veth1 -F pcap -w /tmp/veth1.pcap

# Stream to veth0
sudo ./pcapfile-streamer -r ../contrib/small_lo.pcap -i veth0
```

## TODOs
* Fix this:
```
2021/04/09 12:16:22.045394 [/tmp/3-17.pcap] Ethernet IPv4 packet size: 1522 Bytes
2021/04/09 12:16:22.045425 [/tmp/3-17.pcap] Streaming ethernet IPv4 packet to veth0 failed: send: Message too long
```
