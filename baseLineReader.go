package smallMR

import (
	"io"
	"errors"
	"strconv"
)

type baseLineReader struct{
	in io.Reader
	bufferSize int
	buffer []byte
	// the number of bytes of real data in the buffer
	bufferLength int
	// the current position in the buffer
	// bytes before bufferPosn has been readed
	bufferPosn int
	// The line delimiter
	recordDelimiterBytes []byte
}

func newBaseLineReader(in io.Reader, bufferSize int, recordDelimiterBytes []byte) lineReader {
	blr := &baseLineReader{
		in:in,
		bufferSize:bufferSize,
		buffer : make([]byte, bufferSize),
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
func (blr *baseLineReader) fillBuffer() (int, error) {
	return blr.in.Read(blr.buffer)
}

// from org.apache.hadoop.util.LineReader
func (blr *baseLineReader) readLine(str string, maxLineLength int, maxBytesToConsume int) (int, error){
	if blr.recordDelimiterBytes == nil {
		return blr.readDefaultLine(str, maxLineLength, maxBytesToConsume)
	}else {
		return blr.readCustomLine(str, maxLineLength, maxBytesToConsume)
	}
}


/*
   * Read a line terminated by one of CR, LF, or CRLF.
   * from org.apache.hadoop.util.LineReader
*/
func (blr *baseLineReader) readDefaultLine(str string, maxLineLength int, maxBytesToConsume int) (int, error){
	str = ""
	txtLength := 0 //tracks str.getLength(), as an optimization
	newlineLength := 0 //length of terminating newline
	readLength := 0
	prevCharCR := false //true of prev char was CR
	var bytesConsumed int = 0
	
	startPosn := 0
	appendLength := 0
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
				if err == io.EOF{
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
				break;
			}
			
			//Mac:'\r'(CR)
			if (prevCharCR){
				newlineLength = 1
				break;
			}
			
			prevCharCR = (blr.buffer[blr.bufferPosn] == CR)
		}

		readLength = blr.bufferPosn - startPosn
		if prevCharCR && newlineLength == 0 { // buffer = ['a','b','\r']
			readLength-- //CR at the end of the buffer
		}

    	bytesConsumed += readLength
		appendLength = readLength - newlineLength // delete delimiter from str
		
		if appendLength > maxLineLength - txtLength {
			appendLength = maxLineLength - txtLength
		}

		if appendLength > 0 {
			// str.append(buffer, startPosn, appendLength);
			txtLength += appendLength
		}
	  
	  	//while (newlineLength == 0 && bytesConsumed < maxBytesToConsume);
		if newlineLength != 0 || bytesConsumed >= maxBytesToConsume {
			break
		}
	}

    if (bytesConsumed > INT_MAX) {
		return 0, errors.New("Too many bytes before newline: " + strconv.Itoa(bytesConsumed))
    }
    return bytesConsumed, nil
}

func (blr *baseLineReader) readCustomLine(str string, maxLineLength int, maxBytesToConsume int) (int, error){
	return 0, nil //TODO
}
