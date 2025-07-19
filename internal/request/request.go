package request

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	req, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read request: %w", err)
	}

	requestLineBytes := bytes.Split(req, []byte{'\r', '\n'})[0]
	requestLine, err := parseRequestLine(string(requestLineBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to parse request line: %w", err)
	}

	return &Request{
		RequestLine: requestLine,
	}, nil
}

func parseRequestLine(line string) (RequestLine, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return RequestLine{}, errors.New("invalid number of request line parts")
	}

	method := parts[0]
	requestTarget := parts[1]
	httpVersionParts := parts[2]

	for _, r := range strings.ToLower(method) {
		if r < 'a' || r > 'z' {
			return RequestLine{}, errors.New("method contains non-alpha chars")
		}
	}

	if !strings.Contains(requestTarget, "/") {
		return RequestLine{}, errors.New("invalid path")
	}

	httpVersion := strings.Split(httpVersionParts, "/")[1]
	if httpVersion != "1.1" {
		return RequestLine{}, errors.New("unsupported HTTP version")
	}

	return RequestLine{
		HttpVersion:   httpVersion,
		RequestTarget: requestTarget,
		Method:        method,
	}, nil
}
