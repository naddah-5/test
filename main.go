package main

import (
	"D7024E/udp"
	"encoding/json"
	"net"
)

func main() {
	var dummyRPC []byte
	dummyRPC, _ = json.Marshal("This is a test")
	dummyAddr, _ := net.ResolveUDPAddr("udp4", "127.0.0.1:4001")
	go udp.UDPListener()
	udp.UDPSender(dummyAddr.IP, dummyAddr.Port, dummyRPC)
	for {

	}
}
