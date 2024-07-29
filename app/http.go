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
	Headers       []string
	Body          []byte
}

func (hr HttpRequest) toBytes() []byte {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s %s %s\r\n", hr.Method, hr.RequestTarget, hr.HttpVersion))

	for _, header := range hr.Headers {
		buffer.WriteString(fmt.Sprintf("%s\r\n", header))
	}

	buffer.WriteString("\r\n")
	buffer.Write(hr.Body)

	return buffer.Bytes()
}

type HttpResponse struct {
	HttpVersion   string
	StatusCode    string
	StatusMessage string
	Headers       []string
	Body          []byte
}

func (hr HttpResponse) toBytes() []byte {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s %s %s\r\n", hr.HttpVersion, hr.StatusCode, hr.StatusMessage))

	for _, header := range hr.Headers {
		buffer.WriteString(fmt.Sprintf("%s\r\n", header))
	}

	buffer.WriteString("\r\n")
	buffer.Write(hr.Body)

	return buffer.Bytes()
}

func NewHttpRequest(request []byte) HttpRequest {
	tmpReq := bytes.Split(request, []byte("\r\n\r\n"))
	metaData := strings.Split(string(tmpReq[0]), "\r\n")
	startLine := strings.Split(metaData[0], " ")

	return HttpRequest{
		Method:        startLine[0],
		RequestTarget: startLine[1],
		HttpVersion:   startLine[2],
		Headers:       metaData[1:],
		Body:          tmpReq[1],
	}
}
