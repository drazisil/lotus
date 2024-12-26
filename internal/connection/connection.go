package connection

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/rustymotors/lotus/internal/authlogin"
	"github.com/rustymotors/lotus/internal/shard"
)

// HandleTCPConnection handles an incoming TCP connection.
// It reads data from the connection, unmarshals it into a RawNPSPacket,
// and logs the received packet along with the local port number.
//
// Parameters:
//   - conn: The net.Conn representing the TCP connection.
//
// The function performs the following steps:
//  1. Logs the remote address of the incoming connection.
//  2. Reads up to 1024 bytes of data from the connection.
//  3. Unmarshals the read data into a RawNPSPacket.
//  4. Logs any errors encountered during reading or unmarshalling.
//  5. Logs the received packet and the local port number.
func HandleTCPConnection(conn net.Conn) {
	log.Println("TCP connection from", conn.RemoteAddr())

	// read incoming data
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("Error reading:", err)
		return
	}

	var packet Packet

	// determine the packet type based on the local port
	switch conn.LocalAddr().(*net.TCPAddr).Port {
	case 8226:
		log.Println("Handling packet for port 8226")
		packet = LoginPacket{}
	case 8227:
		log.Println("Handling packet for port 8227")
		packet = RawNPSPacket{}
	case 8228:
		log.Println("Handling packet for port 8228")
		packet = RawNPSPacket{}
	case 7003:
		log.Println("Handling packet for port 7003")
		packet = RawNPSPacket{}
	case 43300:
		log.Println("Handling packet for port 43300")
		packet = RawNPSPacket{}
	default:
		log.Panicln("Unknown port: %d", conn.LocalAddr().(*net.TCPAddr).Port)
		return
	}

	// unmarshal the raw packet
	err = packet.UnmarshalBinary(buf[:n])
	if err != nil {
		log.Println("Error unmarshalling:", err)
		return
	}

	localPort := conn.LocalAddr().(*net.TCPAddr).Port

	// print the raw packet
	log.Println("Received packet from", localPort, ":", packet)
}

// HandleHTTPRequest processes incoming HTTP requests and routes them to the appropriate handler
// based on the request URL path. It also logs the request details and body.
//
// Parameters:
//   - w: http.ResponseWriter to write the response
//   - r: *http.Request containing the request details
//
// Supported URL paths:
//   - "/AuthLogin": Routes the request to authlogin.HandleAuthLogin
//   - "/ShardList/": Routes the request to shard.HandleShardList
//   - Default: Logs "Other" for all other paths
func HandleHTTPRequest(w http.ResponseWriter, r *http.Request) {
	// print the request
	fmt.Println("Request received: ", r.RemoteAddr, r.Method, r.URL.Path, r.URL.Query())

	// print the request body
	fmt.Println("Request body: ", r.Body)

	switch r.URL.Path {
	case "/AuthLogin":
		// handle AuthLogin
		authlogin.HandleAuthLogin(r, w)

	case "/ShardList/":
		// handle ShardList
		shard.HandleShardList(r, w)

	default:
		// handle all other requests
		fmt.Println("Other")
	}
}
