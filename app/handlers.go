package main

import (
	"strconv"
	"strings"
)

func okResponse(_ HttpRequest) HttpResponse {
	return HttpResponse{
		HttpVersion:   httpVersion,
		StatusCode:    "200",
		StatusMessage: "OK",
	}
}

func notFoundResponse(_ HttpRequest) HttpResponse {
	return HttpResponse{
		HttpVersion:   httpVersion,
		StatusCode:    "404",
		StatusMessage: "Not Found",
	}
}

func getEchoHandler(request HttpRequest) HttpResponse {
	echo := strings.Split(request.RequestTarget, "/echo/")[1]

	return HttpResponse{
		HttpVersion:   httpVersion,
		StatusCode:    "200",
		StatusMessage: "OK",
		Headers: map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": strconv.Itoa(len(echo)),
		},
		Body: []byte(echo),
	}
}

func getUserAgentHandler(request HttpRequest) HttpResponse {
	return HttpResponse{
		HttpVersion: httpVersion,
		StatusCode: "200",
		StatusMessage: "OK",
		Headers: map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": strconv.Itoa(len(request.Headers["User-Agent"])),
		},
		Body: []byte(request.Headers["User-Agent"]),
	}
}

func getFilesHandler(request HttpRequest) HttpResponse {
	filePath := strings.Split(request.RequestTarget, "/files/")[1]

	httpResponse := notFoundResponse(request)

	body, err := GetFileContent(filePath)
	if err == nil {
		httpResponse = HttpResponse{
			HttpVersion: httpVersion,
			StatusCode: "200",
			StatusMessage: "OK",
			Headers: map[string]string{
				"Content-Type": "application/octet-stream",
				"Content-Length": strconv.Itoa(len(body)),
			},
			Body: body,
		}
	} 

	return httpResponse
}

func postFilesHandler(request HttpRequest) HttpResponse {
	filePath := strings.Split(request.RequestTarget, "/files/")[1]

	contentSize, err := strconv.Atoi(request.Headers["Content-Length"])
	if err != nil {
		return HttpResponse{
			HttpVersion: httpVersion,
			StatusCode: "400",
			StatusMessage: "Bad Request",
		}
	}

	err = writeFile(filePath, request.Body, contentSize)
	if err != nil {
		return HttpResponse{
			HttpVersion: httpVersion,
			StatusCode: "500",
			StatusMessage: "Internal Server Error",
		}
	} 

	return HttpResponse{
		HttpVersion: httpVersion,
		StatusCode: "201",
		StatusMessage: "Created",
	}
}

