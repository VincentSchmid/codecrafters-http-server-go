package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpRequest_toBytes(t *testing.T) {
	request := HttpRequest{
		Method:        "GET",
		RequestTarget: "/",
		Headers:       map[string]string{"Host": "example.com", "Connection": "keep-alive"},
		Body:          []byte("Test body"),
	}
	expected := []byte("GET / HTTP/1.1\r\nHost: example.com\r\nConnection: keep-alive\r\n\r\nTest body")
	result := request.toBytes()

	assert.Equal(t, expected, result, "The byte representation of the HTTP request is incorrect")
}

func TestHttpResponse_toBytes(t *testing.T) {
	response := HttpResponse{
		Status:  OkStatus,
		Headers: map[string]string{"Content-Type": "text/html", "Content-Length": "12"},
		Body:    []byte("Hello World!"),
	}
	expected := []byte("HTTP/1.1 200 OK\r\nContent-Type: text/html\r\nContent-Length: 12\r\n\r\nHello World!")
	result := response.toBytes()

	assert.Equal(t, expected, result, "The byte representation of the HTTP response is incorrect")
}

func TestNewHttpRequest(t *testing.T) {
	request := []byte("POST /submit HTTP/1.1\r\nHost: example.com\r\nContent-Type: application/x-www-form-urlencoded\r\n\r\nname=John")
	expected := HttpRequest{
		Method:        "POST",
		RequestTarget: "/submit",
		Headers:       map[string]string{"Host": "example.com", "Content-Type": "application/x-www-form-urlencoded"},
		Body:          []byte("name=John"),
	}
	result := NewHttpRequest(request)

	assert.Equal(t, expected, result, "The parsed HTTP request is incorrect")
}
