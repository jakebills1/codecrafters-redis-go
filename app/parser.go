package main

import (
	"bufio"
	"errors"
	"log"
	"strconv"
	"strings"
)

// commands are arrays of bulk strings
// *n means an array where n is the number of elements
// bulk strings are prefixed with $n where n is the length of the string
// at this stage, a 1 element array has to be PING
// and a 2 has to be ECHO

type Command interface {
	encodedResponse() []byte
}
type Ping struct {
	Message string
}
type Echo struct {
	Message string
}

func (p Ping) encodedResponse() []byte {
	return encodeAsSimpleString(p.Message)
}

func (e Echo) encodedResponse() []byte {
	return encodeAsBulkString(e.Message)
}

func (u Unknown) encodedResponse() []byte {
	return []byte("Unknown command")
}

type Unknown struct{}

func parseCommand(commandParts []string) Command {
	// fmt.Println("commandParts =", commandParts)
	// fmt.Println("commandParts[0] =", commandParts[0])

	switch commandParts[0] {
	case "ECHO":
		return &Echo{Message: commandParts[1]}
	case "PING":
		return &Ping{Message: "PONG"}
	default:
		return &Unknown{}
	}
}

func extractCommandParts(scanner *bufio.Scanner) (*[]string, error) {
	scan := scanner.Scan()
	if !scan {
		return nil, errors.New("end of input")
	}
	// first will be *n, indicating the size of the RESP array
	arraySize, err := strconv.Atoi(strings.TrimPrefix(scanner.Text(), "*"))
	if err != nil {
		log.Fatalln("could not parse", scanner.Text(), "as int:", err)
	}
	commandArr := make([]string, arraySize)
	// fmt.Println("len(commandArr) = ", len(commandArr))
	for i := 0; i < arraySize; i++ {
		scanner.Scan() // gets the string size, don't need right now
		scanner.Scan() // gets the string
		line := scanner.Text()
		commandArr[i] = line
	}
	return &commandArr, nil
}
