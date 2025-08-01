package request

import "fmt"

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion string
	RequestTarget string
	Method string
}

var ERROR_BAD_START_LINE = fmt.Errorf("ERROR_BAD_START_LINE")

func RequestFromReader (reader io.Reader) (*Request, error) {
}
