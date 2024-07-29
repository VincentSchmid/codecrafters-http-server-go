package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
)

type RequestHandler func(HttpRequest) HttpResponse

const (
	httpVersion = "HTTP/1.1"
	maxRequestSizeBytes = 1024
)

var (
	echoPath, _ = regexp.Compile("/echo/*")
	planePath, _ = regexp.Compile("/$")
	userAgentPath, _ = regexp.Compile("/user-agent$")
	filesPath, _ = regexp.Compile("/files/*")

	handlers = map[*regexp.Regexp]RequestHandler{
		echoPath:  echoHandler,
		planePath: okResponse,
		userAgentPath: userAgentHandler,
		filesPath: filesHandler,
	}

	filesDir = "assets"
)

func main() {
	if len(os.Args) >= 3 && os.Args[1] == "--directory" {
		filesDir = os.Args[2]
	}
	
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	fmt.Printf("server listening on port 4221")

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
	req := make([]byte, maxRequestSizeBytes)
	conn.Read(req)
	httpReq := NewHttpRequest(req)

	requestHandler := notFoundResponse

	for validTargetRegex, handler := range handlers {
		if validTargetRegex.MatchString(httpReq.RequestTarget)  {
			requestHandler = handler
			break
		}
	}

	resp := requestHandler(httpReq)

	_, err := conn.Write(resp.toBytes())
	if err != nil {
		fmt.Println("error writing to connection:", err.Error())
	}
}
