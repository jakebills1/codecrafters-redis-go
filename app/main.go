package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	for {
		conn, acceptErr := l.Accept()
		if acceptErr != nil {
			fmt.Println("Error accepting connection: ", acceptErr.Error())
			os.Exit(1)
		}
		go handleConn(conn)
	}
	os.Exit(0)
}

func handleConn(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	endReached := false
	for !endReached {
		commandArr, err := extractCommandParts(scanner)
		if err != nil {
			endReached = true
		} else {
			command := parseCommand(*commandArr)
			conn.Write(command.encodedResponse())
		}
	}

}
