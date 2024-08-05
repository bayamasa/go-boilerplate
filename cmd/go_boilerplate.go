package main

import (
	"fmt"

	"github.com/bayamasa/go-boilerplate/server"
)

func main() {
	s := server.NewServer("8080")
	if err := s.Run(); err != nil {
		fmt.Println(err)
	}
}
