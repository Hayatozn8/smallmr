package smallmr

import (
	"io"
)

type lineRecordReader struct {
	start int64
	pos int64
	end int64
	recordDelimiterBytes []byte
	in io.Reader

}

func newLineRecordReader(start int64, end int64, recordDelimiterBytes []byte) recordReader{
	return &lineRecordReader{
		start:start
		recordDelimiterBytes:recordDelimiterBytes,
	}
}

func (lrr *lineRecordReader) initialize() error{

}


func (lrr *lineRecordReader) nextKeyValue() (bool, error){

}
