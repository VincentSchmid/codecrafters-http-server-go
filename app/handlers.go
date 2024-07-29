package main

import (
	"fmt"
	"strings"
)

func okResponse(_ HttpRequest) HttpResponse {
	return HttpResponse{
		HttpVersion: httpVersion,
		StatusCode: "200",
		StatusMessage: "OK",
	}
}

func notFoundResponse(_ HttpRequest) HttpResponse {
	return HttpResponse{
		HttpVersion: httpVersion,
		StatusCode: "404",
		StatusMessage: "Not Found",
	}
}

func echoHandler(request HttpRequest) HttpResponse {
	echo := strings.Split(request.RequestTarget, "/echo/")[1]

	return HttpResponse{
		HttpVersion: httpVersion,
		StatusCode: "200",
		StatusMessage: "OK",
		Headers: []string{
			"Content-Type: text/plain",
			fmt.Sprintf("Content-Length: %d", len(echo)),
		},
		Body: []byte(echo),
	}
}
