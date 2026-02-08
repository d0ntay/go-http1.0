package api

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func route(c net.Conn) {
	defer c.Close()

	reader := bufio.NewReader(c)
	writer := bufio.NewWriter(c)
	defer writer.Flush()

	req, err := parseRequest(*reader)
	if err != nil {
		fmt.Printf("%q",err)
	}

	switch req.Uri{
		case "/":
			log.Printf("Method: %s | Uri: %s | Version: %s | Headers: %s \n",req.Method, req.Uri ,req.Version, req.Headers)
			homeHandler(writer, req)
		case "/health":
			log.Printf("Method: %s | Uri: %s | Version: %s | Headers: %s \n",req.Method, req.Uri ,req.Version, req.Headers)
			healthChechHandler(writer, req)
		default:
			log.Printf("Method: %s | Uri: %s | Version: %s | Headers: %s \n",req.Method, req.Uri ,req.Version, req.Headers)
			notFoundHandler(writer)
	}
}
