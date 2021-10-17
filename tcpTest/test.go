// This module contains the communication functions
// Taken from class slides
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
)

type Message struct {
	Sender   string
	Receiver string
	Body     interface{}
}

// Send any data to desired ip
func Send(data interface{}, ip string) error {
	var conn net.Conn
	var err error
	var encoder *gob.Encoder

	conn, err = net.Dial("tcp", ip)

	if err != nil {
		panic("Client connection error")
	}

	binBuffer := new(bytes.Buffer)

	encoder = gob.NewEncoder(binBuffer)
	err = encoder.Encode(data)

	conn.Write(binBuffer.Bytes())

	return err
}

// listen for messages
func Receive(data interface{}, listener *net.Listener) error {
	var conn net.Conn
	var err error
	var decoder *gob.Decoder

	conn, err = (*listener).Accept()
	if err != nil {
		panic("Server accept connection error")
	}

	decoder = gob.NewDecoder(conn)

	err = decoder.Decode(data)

	return err
}

func main() {
	msg := Message{
		Sender:   "test",
		Receiver: "P0",
		Body:     "Hola",
	}
	fmt.Println(msg.Body)
	Send(msg, "127.0.0.1:7000")
}
