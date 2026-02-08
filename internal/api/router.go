package api

import (
	"fmt"
	"log"
	"net"
)

func route(c net.Conn) {
	defer c.Close()
	
	msg := make([]byte, 1024)
	n, err := c.Read(msg)
	if err != nil {
		log.Fatal("error reading request")
	}
	
	req := string(msg[:n])
	if req[:16] == "GET / HTTP/1.1\r\n"{
		c.Write([]byte("home"))
		c.Close()
	}
	if req[:22] == "GET /health HTTP/1.1\r\n" {
		c.Write([]byte("ok"))
		c.Close()
	}
	fmt.Println(req)
	c.Close()
}
