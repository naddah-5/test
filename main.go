package main

import (
	"D7024E/udp"
	"encoding/json"
	"net"
	"time"
)

func main() {
	var dummyRPC []byte
	var text string = "this is a test"
	dummyRPC, _ = json.Marshal(text)
	dummyAddr, _ := net.ResolveUDPAddr("udp4", "127.0.0.1:4001")
	go udp.UDPListener()
	time.Sleep(10 * time.Millisecond)
	udp.UDPSender(dummyAddr.IP, dummyAddr.Port, dummyRPC)
	for {
		time.Sleep(time.Second)
	}
}
