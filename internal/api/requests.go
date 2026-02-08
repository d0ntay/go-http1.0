package api

import (
	"bufio"
	"fmt"
	"strings"
)

type Request struct {
	Method string
	Uri string
	Version string
	Headers map[string]string
//	Body []byte
}

func parseRequest(r bufio.Reader) (*Request, error) {
		
	var raw []string

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("Failed to parse request")
		}
		raw = append(raw, line)
		if line == "\r\n"{
			break
		}
	}

	requestLine := strings.Fields(strings.TrimSpace(raw[0]))
	if len(requestLine) != 3 {
		return nil, fmt.Errorf("Invalid request line")
	}

	headers := make(map[string]string)
	for _,line := range raw[1:]{
		line = strings.TrimSpace(line)
		if line == " " {
			continue
		}

		parts := strings.SplitN(line,":",2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		headers[key] = val 
	}


	return &Request{
		Method: requestLine[0],
		Uri: requestLine[1],
		Version: requestLine[2],
		Headers: headers,
//		Body: body,
	},nil
}
