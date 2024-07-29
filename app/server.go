package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
)

type RequestHandler func(HttpRequest) HttpResponse

const (
	maxRequestSizeBytes = 1024
)

var (
	echoPath, _      = regexp.Compile("/echo/*")
	planePath, _     = regexp.Compile("/$")
	userAgentPath, _ = regexp.Compile("/user-agent$")
	filesPath, _     = regexp.Compile("/files/*")

	getHandlers = map[*regexp.Regexp]RequestHandler{
		echoPath:      getEchoHandler,
		planePath:     okResponse,
		userAgentPath: getUserAgentHandler,
		filesPath:     getFilesHandler,
	}

	postHandlers = map[*regexp.Regexp]RequestHandler{
		filesPath: postFilesHandler,
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

	if httpReq.Method == "GET" {
		for validTargetRegex, handler := range getHandlers {
			if validTargetRegex.MatchString(httpReq.RequestTarget) {
				requestHandler = handler
				break
			}
		}
	}

	if httpReq.Method == "POST" {
		for validTargetRegex, handler := range postHandlers {
			if validTargetRegex.MatchString(httpReq.RequestTarget) {
				requestHandler = handler
				break
			}
		}
	}

	resp := requestHandler(httpReq)

	_, err := conn.Write(resp.toBytes())
	if err != nil {
		fmt.Println("error writing to connection:", err.Error())
	}
}
