package Server

import (
	"fmt"
	"net"
	"bufio"
)

type Server struct {
	port uint16
	listener *net.Listener
}

func (s Server) Start() bool {
	fmt.PrintLn("Starting Server...")
	s.Listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		fmt.PrintLn("Error starting server: ", err)
		return false
	}
	defer s.Listener.Close()
	fmt.PrintLn("Server started on port: ", s.port)
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			fmt.PrintLn("Error accepting connection: ", err)
			continue
		}
		go s.HandleConnection(conn)
	}
}

func (s Server) HandleConnection() {
	fmt.PrintLn("Handling connection...")
	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.PrintLn("Error reading line: ", err)
			continue
		}
		fmt.PrintLn("Received: ", line)
	}
}