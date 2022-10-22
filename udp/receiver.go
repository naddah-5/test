package udp

import (
	"encoding/json"
	"fmt"
	"net"
)

// Listener of the udp messages.
// Establish udp4 listner by address reccived from ip and port.
// Read from udp connection into buffer to reccive whole message.
func UDPListener() {
	dummyAddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:4001")
	if err != nil {
		fmt.Println("v%", err)
	}
	connection, err := net.ListenUDP("udp4", dummyAddr)
	if err != nil {
		fmt.Println("There was an error in the ListenUDP:", err)
	} else {
		fmt.Println("Setup for listening to udp over", dummyAddr.IP, ":", dummyAddr.Port)
	}
	defer connection.Close()
	buffer := make([]byte, 4096)
	// connection.SetReadDeadline(time.Now().Add(4 * time.Second))
	n, addr, _ := connection.ReadFromUDP(buffer)
	fmt.Println("RETURNED MARSHALED", string(buffer[:n]))

	retMessage, _ := json.Marshal("THIS IS A RETURN TEST")

	connection.WriteToUDP(retMessage, addr)
}
