package util

import "strings"

type LineReader interface {
	//
	ReadLine(str *strings.Builder, maxLineLength int32, maxBytesToConsume int32) (int32, error)
	// Close stream
	// Close() error
}
