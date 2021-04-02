# zero-httpd
Dummy HTTP server creating large TCP/IP sessions for pcap capturing

## Requirements
* go version go1.15.8

## Build zero-httpd
```
go build
```

## Run zero-httpd
```
./zero-httpd
```

## Usage
```
# Request small reponse (default: 128 B)
$ curl -s --output /tmp/test.dat http://localhost:8888/
$ ls -lha /tmp/test.dat 
-rw-rw-r-- 1 andi andi 128 Apr  2 19:41 /tmp/test.dat

# Request large response (e.g. 1 GB)
$ curl -s --output /tmp/test.dat http://localhost:8888/?mb=1024
$ ls -lha /tmp/test.dat 
-rw-rw-r-- 1 andi andi 1,0G Apr  2 19:43 /tmp/test.dat

# Capture request and response with tshark (run as root)
# tshark -i lo -f "port 8888" -F pcap -w /tmp/test.pcap
```