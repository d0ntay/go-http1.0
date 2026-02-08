package api

import (
	"fmt"
	"net"
)

type Server struct {
	Cfg Config
}

func NewServer(c Config) (*Server, error) {
	return &Server{
		Cfg: c,
	},nil
}

func (srv *Server) Start() {
	l, err := net.Listen("tcp4",srv.Cfg.Addr)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server started: %s\n",l.Addr().String())
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go route(conn)
	}
}
