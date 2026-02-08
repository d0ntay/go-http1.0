package main

import (
	"github.com/d0ntay/go-http1/internal/api"
)

func main() {

	cfg := &api.Config{
		Addr: ":8080",
	}

	srv, err:= api.NewServer(*cfg)
	if err != nil {
		panic(err)
	}

	srv.Start()
}
