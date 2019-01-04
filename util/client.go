package util

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"

	"github.com/askft/go-behave/core"
)

// Client ...
type Client struct{}

// Run ...
func (c *Client) Run() {
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		panic(fmt.Errorf("Connection error: %v", err))
	}
	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Request node: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")

		var data core.Node

		dec := gob.NewDecoder(conn)
		if err := dec.Decode(data); err != nil {
			panic(err)
		}
		fmt.Println(data)
	}
}
