package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"bytes"
)

var (
	availablePaths = [...]string{"/", "/index.html"}
)

type HttpRequest struct {
	Method string
	RequestTarget string
	HttpVersion string
	Headers []string
	MessageBody []byte
}

func NewHttpRequest(request []byte) HttpRequest {
	tmpReq := bytes.Split(request, []byte("\r\n\r\n"))
	metaData := strings.Split(string(tmpReq[0]), "\r\n")
	startLine := strings.Split(metaData[0], " ")

	return HttpRequest{
		Method: startLine[0],
		RequestTarget: startLine[1],
		HttpVersion: startLine[2],
		Headers: metaData[1:],
		MessageBody: tmpReq[1],
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	req := make([]byte, 1024)
	conn.Read(req)
	httpReq := NewHttpRequest(req)

	if validRequestTarget(httpReq.RequestTarget) {
		_, err := conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
		if err != nil {
			fmt.Println("error writing to connection:", err.Error())
		}
	} else {
		_, err := conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		if err != nil {
			fmt.Println("error writing to connection:", err.Error())
		}
	}
}

func validRequestTarget(requestTarget string) bool {
	for _, validTarget := range availablePaths {
        if validTarget == requestTarget {
            return true
        }
    }
    return false
}
