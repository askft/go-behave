package util

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"sync"

	"github.com/askft/go-behave/core"
)

// Server ...
type Server struct {
	sync.Mutex
	Trees map[string]core.Node
}

// NewServer ...
func NewServer() *Server {
	return &Server{
		Trees: map[string]core.Node{},
	}
}

// Register ...
func (s *Server) Register(name string, node core.Node) {
	if _, ok := s.Trees[name]; ok {
		fmt.Printf("Warning: Overwriting node for %q!", name)
	}
	s.Trees[name] = node
}

// Run ...
func (s *Server) Run(port string) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	msg, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Message received:", string(msg))

	s.Lock()
	node, ok := s.Trees[msg]
	s.Unlock()
	if !ok {
		fmt.Printf("tree %q not found\n", msg)
		return
	}

	if err := gob.NewEncoder(conn).Encode(node); err != nil {
		panic(err)
	}
	fmt.Printf("Sent: %+v", node)
}
