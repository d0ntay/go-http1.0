package api

import (
	"bufio"
	"fmt"
)

func notFoundHandler(w *bufio.Writer) {
	body := "404"
	w.WriteString("HTTP/1.0 404 NOT FOUND\r\n")
	w.WriteString("Content-Type: text/plain\r\n")
	w.WriteString(fmt.Sprintf("Content-Length: %d\r\n",len(body)))
	w.WriteString("\r\n")
	w.WriteString(body)
}

func homeHandler(w *bufio.Writer) {
	body := "home"
	w.WriteString("HTTP/1.0 200 OK\r\n")
	w.WriteString("Content-Type: text/plain\r\n")
	w.WriteString(fmt.Sprintf("Content-Length: %d\r\n",len(body)))
	w.WriteString("\r\n")
	w.WriteString(body)
}

func healthChechHandler(w *bufio.Writer) {
	body := `{"status":"ok"}`
	w.WriteString("HTTP/1.0 200 OK\r\n")
	w.WriteString("Content-Type: application/json\r\n")
	w.WriteString(fmt.Sprintf("Content-Length: %d\r\n",len(body)))
	w.WriteString("\r\n")
	w.Write([]byte(body))
}
