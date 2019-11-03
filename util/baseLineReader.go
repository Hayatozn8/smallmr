package util

import (
	"errors"
	"io"
	"math"
	"os"
	"strconv"
)

type BaseLineReader struct {
	in         os.File
	bufferSize int32
	buffer     []byte
	// the number of bytes of real data in the buffer
	bufferLength int32
	// the current position in the buffer
	// bytes before bufferPosn has been readed
	bufferPosn int32
	// The line delimiter
	recordDelimiterBytes []byte
}

func NewBaseLineReader(in os.File, bufferSize int32, recordDelimiterBytes []byte) LineReader {
	blr := &BaseLineReader{
		in:         in,
		bufferSize: bufferSize,
		buffer:     make([]byte, bufferSize),
	}

	if recordDelimiterBytes != nil {
		blr.recordDelimiterBytes = recordDelimiterBytes //TODO
	}

	return blr
}

// func (blr *baseLineReader) Close() error{
// 	return blr.in.Close()
// }

// from org.apache.hadoop.util.LineReader
func (blr *BaseLineReader) fillBuffer() (int32, error) {
	n, err := blr.in.Read(blr.buffer)
	return int32(n), err
}

// from org.apache.hadoop.util.LineReader
func (blr *BaseLineReader) ReadLine(str *string, maxLineLength int32, maxBytesToConsume int32) (int32, error) {
	if blr.recordDelimiterBytes == nil {
		return blr.readDefaultLine(str, maxLineLength, maxBytesToConsume)
	} else {
		return blr.readCustomLine(str, maxLineLength, maxBytesToConsume)
	}
}

/*
 * Read a line terminated by one of CR, LF, or CRLF.
 * from org.apache.hadoop.util.LineReader
 */
func (blr *BaseLineReader) readDefaultLine(str *string, maxLineLength int32, maxBytesToConsume int32) (int32, error) {
	*str = ""
	var txtLength int32 = 0     //tracks str.getLength(), as an optimization
	var newlineLength int32 = 0 //length of terminating newline
	var readLength int32 = 0
	prevCharCR := false //true of prev char was CR
	var bytesConsumed int64 = 0

	var startPosn int32 = 0
	var appendLength int32 = 0
	var err error = nil
	for {
		startPosn = blr.bufferPosn //starting from where we left off the last time
		// if buffer is finished, reload
		if blr.bufferPosn >= blr.bufferLength {
			startPosn = 0
			blr.bufferPosn = 0

			if prevCharCR {
				bytesConsumed++ //account for CR from previous read
			}

			blr.bufferLength, err = blr.fillBuffer()
			if err != nil { // EOF or other Error
				if err == io.EOF {
					break
				} else {
					return 0, err
				}
			}
		}

		for ; blr.bufferPosn < blr.bufferLength; blr.bufferPosn++ { //search for newline
			if blr.buffer[blr.bufferPosn] == LF {
				if prevCharCR {
					newlineLength = 2 //Windows:'\r\n' (CR)(LF)
				} else {
					newlineLength = 1 //UNIX:'\n' (LF)
				}

				blr.bufferPosn++ // at next invocation proceed from following byte
				break
			}

			//Mac:'\r'(CR)
			if prevCharCR {
				newlineLength = 1
				break
			}

			prevCharCR = (blr.buffer[blr.bufferPosn] == CR)
		}

		readLength = blr.bufferPosn - startPosn
		if prevCharCR && newlineLength == 0 { // buffer = ['a','b','\r']
			readLength-- //CR at the end of the buffer
		}

		bytesConsumed += int64(readLength)
		appendLength = readLength - newlineLength // delete delimiter from str

		if appendLength > maxLineLength-txtLength {
			appendLength = maxLineLength - txtLength
		}

		if appendLength > 0 {
			*str += string(blr.buffer[startPosn : startPosn+appendLength])
			txtLength += appendLength
		}

		//while (newlineLength == 0 && bytesConsumed < maxBytesToConsume);
		if newlineLength != 0 || bytesConsumed >= int64(maxBytesToConsume) {
			break
		}
	}

	if bytesConsumed > math.MaxUint32 {
		return 0, errors.New("Too many bytes before newline: " + strconv.FormatInt(bytesConsumed, 10))
	}
	return int32(bytesConsumed), nil
}

func (blr *BaseLineReader) readCustomLine(str *string, maxLineLength int32, maxBytesToConsume int32) (int32, error) {
	return 0, nil //TODO
}
