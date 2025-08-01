package request

import (
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

var ERROR_BAD_START_LINE = fmt.Errorf("ERROR_BAD_START_LINE")
var SEPARATOR = "\r\n"

func parseRequestLine(b string) (*Request, string, error) {
	idx := strings.Index(b, SEPARATOR)
	if idx == -1 {
		return nil, b , nil
	}

	startLine := b[:idx]
	restOfMsg := b[idx+len(SEPARATOR):]

	parts := strings.Split(startLine, " ")
	if len(parts)!= 3 {
		return nil, restOfMsg, ERROR_BAD_START_LINE
	}

	return &RequestLine{
		Method: parts[0],
		RequestTarget: parts[1],
		HttpVersion: parts[2],
	}, restOfMsg, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {
}
