package main

import "fmt"

func encodeAsBulkString(message string) []byte {
	encodedMessage := fmt.Sprintf("$%d\r\n%s\r\n", len(message), message)
	return []byte(encodedMessage)
}

func encodeAsSimpleString(message string) []byte {
	return []byte(fmt.Sprintf("+%s\r\n", message))
}
