package main

import "fmt"

func encodeAsBulkString(message string) []byte {
	if message == "" {
		return nullBulkString()
	}
	encodedMessage := fmt.Sprintf("$%d\r\n%s\r\n", len(message), message)
	return []byte(encodedMessage)
}

func nullBulkString() []byte {
	return []byte(fmt.Sprintf("$-1\r\n"))
}

func encodeAsSimpleString(message string) []byte {
	return []byte(fmt.Sprintf("+%s\r\n", message))
}
