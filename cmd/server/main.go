package main

import (
	"fmt"

	"github.com/askft/go-behave/util"
)

func main() {
	s := util.NewServer()
	port := ":8081"
	fmt.Println("Starting server on http://127.0.0.1" + port)
	err := s.Run(port)
	if err != nil {
		panic(err)
	}
}
