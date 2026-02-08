package main

import (
	"github.com/d0ntay/go-http1/internal/api"
)

func main() {
	cfg, err := api.NewConfig(":8080")
	if err!= nil {
		panic(err)
	}
	srv, err:= api.NewServer(*cfg)
	if err != nil {
		panic(err)
	}

	srv.Start()
}
