package main

import "strconv"

// commands are arrays of bulk strings
// *n means an array where n is the number of elements
// bulk strings are prefixed with $n where n is the length of the string
// at this stage, a 1 element array has to be PING
// and a 2 has to be ECHO

type Command interface {
	message() []byte
}
type Ping struct {
	Message string
}
type Echo struct {
	Message string
}

func (p Ping) message() []byte {
	return encodeAsSimpleString(p.Message)
}

func (e Echo) message() []byte {
	return encodeAsBulkString(e.Message)
}

func (u Unknown) message() []byte {
	return []byte("Unknown command")
}

type Unknown struct{}

func parseCommand(command []byte) Command {
	elements, err := strconv.Atoi(string(command[1]))
	if err != nil {
		return &Unknown{}
	}
	if elements == 1 {
		// ping
		return &Ping{Message: "PONG"}
	} else if elements == 2 {
		return &Echo{Message: parseBulkString(string(command[15:])).Contents}
	} else {
		// presently unknown command
		return &Unknown{}
	}
}

type BulkString struct {
	Length   int
	Contents string
}

func parseBulkString(raw string) BulkString {
	var bs BulkString
	bs.Length = int(raw[1] - '0')
	bs.Contents = raw[4:bs.Length]
	return bs
}
