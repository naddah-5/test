package udp2

import (
	"fmt"
	"log"
	"net"
	"time"
)

// Should be able to replace read response now since it returns the marshalled response?
// Still need to fix the listener side? Maybe just rework how the handler responds to data?

/**
 * Establish udp4 connection with given address created from ip and port.
 * Send message over connection.
 */
func UDPSender(remote *net.UDPAddr, message []byte) {

	//addr := ip.String() + ":" + strconv.Itoa(port)

	conn, err := net.Dial("udp4", remote.String())
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	conn.Write(message)
	buffer := make([]byte, 4096)
	conn.SetReadDeadline(time.Now().Add(4 * time.Second))
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("ERROR!", err)
	}
	fmt.Println(n)
	fmt.Println(string(buffer[:n]))

	// addr := ip.String() + ":" + strconv.Itoa(port)
	// conn, err := net.Dial("udp4", addr)
	// if err != nil {
	// 	println("v%", err)
	// }
	// defer conn.Close()

	// res := make([]byte, 4096)
	// conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	// _, err = conn.Write(message)
	// if err != nil {
	// 	println("Something went wrong in the sender...")
	// }

	// _, err = conn.Read(res)
}
