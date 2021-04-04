# pcap-reader

## References
* https://github.com/google/gopacket
* https://www.devdungeon.com/content/packet-capture-injection-and-analysis-gopacket
* https://wiki.wireshark.org/Development/LibpcapFileFormat

## Requirements
* go version go1.15.8

```
$ sudo apt-get install libpcap-dev
```

## Build pcap-reader
```
go build
```

## Run pcap-reader
```
# Usage
$ ./pcap-reader -h
Usage of ./pcap-reader:
  -i string
    	Path of PCAP input file
```

## TODOs
* Is it possible to check order of ethernet frame sequence?

## Vendor golang dependencies (with go mod)
See: https://golang.org/ref/mod#vendoring

```
$ go help mod
Go mod provides access to operations on modules.

Note that support for modules is built into all the go commands,
not just 'go mod'. For example, day-to-day adding, removing, upgrading,
and downgrading of dependencies should be done using 'go get'.
See 'go help modules' for an overview of module functionality.

Usage:

	go mod <command> [arguments]

The commands are:

	download    download modules to local cache
	edit        edit go.mod from tools or scripts
	graph       print module requirement graph
	init        initialize new module in current directory
	tidy        add missing and remove unused modules
	vendor      make vendored copy of dependencies
	verify      verify dependencies have expected content
	why         explain why packages or modules are needed

Use "go help mod <command>" for more information about a command.

$ go mod init
go: creating new go.mod: module github.com/eg5846/getting-started-with-pcap/pcap-reader

# Add some *.go files with includes

$ go get
go: finding module for package github.com/google/gopacket/pcap
...

$ go mod vendor

$ go mod tidy
```
