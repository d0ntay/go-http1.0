package api

import (
	"bufio"
	"log"
	"net"
)

func route(c net.Conn) {
	defer c.Close()
	reader := bufio.NewReader(c)
	writer := bufio.NewWriter(c)
	defer writer.Flush()

	line, err := reader.ReadString('\n')
	log.Printf("request: %s",line)
	if err != nil {
		log.Fatal("error reading request")
	}
	switch line{
		case "GET / HTTP/1.1\r\n":
			homeHandler(writer)
		case "GET /health HTTP/1.1\r\n":
			healthChechHandler(writer)
	}
}
