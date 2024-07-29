package main

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	HttpVersion = "HTTP/1.1"
)

type HttpStatus struct {
	StatusCode    string
	StatusMessage string
}

var (
	OkStatus = HttpStatus{
		StatusCode:    "200",
		StatusMessage: "OK",
	}
	CreatedStatus = HttpStatus{
		StatusCode:    "201",
		StatusMessage: "Created",
	}
	NotFoundStatus = HttpStatus{
		StatusCode:    "404",
		StatusMessage: "Not Found",
	}
	BadRequestStatus = HttpStatus{
		StatusCode:    "400",
		StatusMessage: "Bad Request",
	}
	InternalServerErrorStatus = HttpStatus{
		StatusCode:    "500",
		StatusMessage: "Internal Server Error",
	}
)

type HttpRequest struct {
	Method        string
	RequestTarget string
	Headers       map[string]string
	Body          []byte
}

func NewHttpRequest(request []byte) HttpRequest {
	tmpReq := bytes.Split(request, []byte("\r\n\r\n"))
	metaData := strings.Split(string(tmpReq[0]), "\r\n")
	startLine := strings.Split(metaData[0], " ")
	headers := make(map[string]string)

	for _, headerValue := range metaData[1:] {
		splitHeaders := strings.Split(headerValue, ": ")
		headers[splitHeaders[0]] = splitHeaders[1]
	}

	return HttpRequest{
		Method:        startLine[0],
		RequestTarget: startLine[1],
		Headers:       headers,
		Body:          tmpReq[1],
	}
}

func (hr HttpRequest) toBytes() []byte {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s %s %s\r\n", hr.Method, hr.RequestTarget, HttpVersion))

	for key, value := range hr.Headers {
		buffer.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}

	buffer.WriteString("\r\n")
	buffer.Write(hr.Body)

	return buffer.Bytes()
}

type HttpResponse struct {
	Status  HttpStatus
	Headers map[string]string
	Body    []byte
}

func (hr HttpResponse) toBytes() []byte {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s %s %s\r\n", HttpVersion, hr.Status.StatusCode, hr.Status.StatusMessage))

	for key, value := range hr.Headers {
		buffer.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}

	buffer.WriteString("\r\n")
	buffer.Write(hr.Body)

	return buffer.Bytes()
}
