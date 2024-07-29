package main

import (
	"bytes"
	"fmt"
	"strings"
)

type HttpRequest struct {
	Method        string
	RequestTarget string
	HttpVersion   string
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
		HttpVersion:   startLine[2],
		Headers:       headers,
		Body:          tmpReq[1],
	}
}

func (hr HttpRequest) toBytes() []byte {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s %s %s\r\n", hr.Method, hr.RequestTarget, hr.HttpVersion))

	for key, value := range hr.Headers {
		buffer.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}

	buffer.WriteString("\r\n")
	buffer.Write(hr.Body)

	return buffer.Bytes()
}

type HttpResponse struct {
	HttpVersion   string
	StatusCode    string
	StatusMessage string
	Headers       map[string]string
	Body          []byte
}

func (hr HttpResponse) toBytes() []byte {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s %s %s\r\n", hr.HttpVersion, hr.StatusCode, hr.StatusMessage))

	for key, value := range hr.Headers {
		buffer.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}

	buffer.WriteString("\r\n")
	buffer.Write(hr.Body)

	return buffer.Bytes()
}
