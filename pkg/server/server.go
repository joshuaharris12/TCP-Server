package server

import (
	"fmt"
	"io"
	"net"

	"github.com/google/uuid"
)

type Server struct {
	Listener net.Listener
}

func NewServer(port string) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

	if err != nil {
		return nil, fmt.Errorf("LOG: Failed to listen on port %s: %w", port, err)
	}

	return &Server{
		Listener: listener,
	}, nil
}

func (s *Server) Run() error {
	defer s.Listener.Close()
	for {
		con, err := s.Listener.Accept()

		if err != nil {
			return fmt.Errorf("LOG: Failed to accept new connection: %w", err)
		}

		fmt.Println("LOG: Ready to accept new connections")

		go handleConnection(con, uuid.NewString())
	}
}

func handleConnection(con net.Conn, uid string) {
	defer con.Close()

	address := con.RemoteAddr().String()
	buffer := make([]byte, 1024)

	for {
		n, err := con.Read(buffer)

		if err == io.EOF {
			fmt.Printf("Client %s is disconnected", address)
			con.Close()
			return
		}

		if err != nil {
			fmt.Printf("LOG: Failed to read data from TCP connection for address %s: %w", address, err)
			con.Close()
			return
		}

		bytes := buffer[:n]
		data := string(bytes)
		fmt.Printf("Client %s: %s", uid, data)
	}
}
