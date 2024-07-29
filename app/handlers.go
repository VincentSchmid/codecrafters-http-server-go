package main

import (
	"strconv"
	"strings"
)

func okResponse(_ HttpRequest) HttpResponse {
	return HttpResponse{
		Status: OkStatus,
	}
}

func notFoundResponse(_ HttpRequest) HttpResponse {
	return HttpResponse{
		Status: NotFoundStatus,
	}
}

func getEchoHandler(request HttpRequest) HttpResponse {
	echo := strings.Split(request.RequestTarget, "/echo/")[1]

	return HttpResponse{
		Status: OkStatus,
		Headers: map[string]string{
			"Content-Type":   "text/plain",
		},
		Body: []byte(echo),
	}
}

func getUserAgentHandler(request HttpRequest) HttpResponse {
	return HttpResponse{
		Status: OkStatus,
		Headers: map[string]string{
			"Content-Type":   "text/plain",
		},
		Body: []byte(request.Headers["User-Agent"]),
	}
}

func getFilesHandler(request HttpRequest) HttpResponse {
	filePath := strings.Split(request.RequestTarget, "/files/")[1]

	httpResponse := HttpResponse{
		Status: NotFoundStatus,
	}

	body, err := GetFileContent(filePath)
	if err == nil {
		httpResponse = HttpResponse{
			Status: OkStatus,
			Headers: map[string]string{
				"Content-Type":   "application/octet-stream",
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
			Status: BadRequestStatus,
		}
	}

	err = writeFile(filePath, request.Body, contentSize)
	if err != nil {
		return HttpResponse{
			Status: InternalServerErrorStatus,
		}
	}

	return HttpResponse{
		Status: CreatedStatus,
	}
}
