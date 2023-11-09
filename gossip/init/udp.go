package init

import (
	"net"
	"os"

	"github.com/libp2p/go-libp2p/core/peer"
)

func SendACK(hostID peer.ID) {
	udpServer, err := net.ResolveUDPAddr("udp", "15.165.22.47:1053")

	if err != nil {
		println("ResolveUDPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, udpServer)
	if err != nil {
		println("Listen failed:", err.Error())
		os.Exit(1)
	}

	//close the connection
	defer conn.Close()

	_, err = conn.Write([]byte(hostID + "  - ACK"))
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}

	// buffer to get data
	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		println("Read data failed:", err.Error())
		os.Exit(1)
	}

	println(string(received))
}
