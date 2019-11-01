package smallmr

type lineReader interface {
	// 
	readLine(str string, maxLineLength int, maxBytesToConsume int) (int, error)
	// Close stream
	// Close() error
}

