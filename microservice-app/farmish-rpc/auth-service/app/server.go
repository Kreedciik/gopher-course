package app

import (
	"farmish/pkg/handler"
	"fmt"
	"log/slog"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Server struct {
	AuthRPC *handler.Handler
}

func (s *Server) Run(port string) error {

	if err := rpc.Register(s.AuthRPC); err != nil {
		return err
	}

	l, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			slog.Error(fmt.Sprintf("rpc-server: %s", err.Error()))
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
