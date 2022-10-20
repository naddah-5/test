package udp

import (
	"net"
)

// Listener of the udp messages.
// Establish udp4 listner by address reccived from ip and port.
// Read from udp connection into buffer to reccive whole message.
func UDPListener() {
	dummyAddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:4001")
	if err != nil {
		println("v%", err)
	}
	connection, err := net.ListenUDP("udp4", dummyAddr)
	if err != nil {
		println("There was an error in the ListenUDP:", err)
	} else {
		println("Setup for listening to udp over v%:v%", dummyAddr.IP, dummyAddr.Port)
	}
	defer connection.Close()
	for {
		buffer := make([]byte, 4096)
		n, _, _ := connection.ReadFromUDP(buffer)
		println((buffer[0:n]))
	}
}
