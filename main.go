package main

import (
	"D7024E/udp"
	"D7024E/udp2"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func main() {
	var dummyRPC []byte
	var text string = "this is a test"
	dummyRPC, _ = json.Marshal(text)
	dummyAddr, _ := net.ResolveUDPAddr("udp4", "127.0.0.1:4001")
	fmt.Println(dummyAddr.String())
	fmt.Println(dummyAddr.IP.String())
	go udp.UDPListener()
	time.Sleep(10 * time.Millisecond)
	udp2.UDPSender(dummyAddr, dummyRPC)
	//for {
	//	time.Sleep(time.Second)
	//}
}
