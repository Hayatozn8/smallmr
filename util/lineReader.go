package util

type LineReader interface {
	//
	ReadLine(str *string, maxLineLength int32, maxBytesToConsume int32) (int32, error)
	// Close stream
	// Close() error
}
