package mapreduce

import (
	"errors"
	"math"
	"os"
	"strings"

	inputSplit "github.com/Hayatozn8/smallmr/split"
	"github.com/Hayatozn8/smallmr/util"
	"github.com/Hayatozn8/smallmr/config"
)

type LineRecordReader struct {
	start                int64
	pos                  int64
	end                  int64
	recordDelimiterBytes []byte
	in                   util.LineReader
	fileIn               *os.File
	key                  int64
	value                strings.Builder //TODO
	maxLineLength        int32
	err                  error
	// ignore zip file TODO

}

// maxLineLength TODO
func NewLineRecordReader(recordDelimiterBytes []byte) RecordReader {
	result := &LineRecordReader{
		recordDelimiterBytes: recordDelimiterBytes,
	}

	return result
}

func (reader *LineRecordReader) Err() error {
	// if reader.err == io.EOF {
	// 	return nil
	// }
	return reader.err
}

// implements
func (reader *LineRecordReader) Initialize(split inputSplit.InputSplit, context TaskContext) error {
	reader.maxLineLength = context.GetConfiguration().GetInt32(config.MAX_LINE_LENGTH)
	
	fsplit, ok := split.(*inputSplit.FileSplit)
	if !ok {
		return errors.New("LineRecordReader.Initalize: split is not a object of inputSplit.FileSplit")
	}

	reader.start = fsplit.GetStart()
	reader.pos = reader.start
	reader.end = reader.start + fsplit.GetLength()

	var err error
	reader.fileIn, reader.err = os.Open(fsplit.GetPath()) //TODO
	if err != nil {
		return err // listen file Open error // TODO
	}
	reader.fileIn.Seek(reader.start, 0)

	reader.in = util.NewBaseLineReader(reader.fileIn, 200, nil)

	if reader.start != 0 {
		num, err := reader.in.ReadLine(&reader.value, 0, reader.maxBytesToConsume(reader.start))
		if err != nil {
			return err
		}

		reader.start += int64(num)
	}

	return nil
}

// implements
func (reader *LineRecordReader) NextKeyValue() bool {
	reader.key = reader.pos

	var newSize int32 = 0

	for reader.pos <= reader.end {
		if reader.pos == 0 {
			newSize, reader.err = reader.skipUtfByteOrderMark()
		} else {
			newSize, reader.err = reader.in.ReadLine(&reader.value, reader.maxLineLength, reader.maxBytesToConsume(reader.pos))
			reader.pos += int64(newSize)
		}

		if reader.err != nil {
			return false
		}

		// EOF:newSize==0,
		if newSize == 0 || newSize < reader.maxLineLength {
			break
		}
	}

	if newSize == 0 {
		reader.value.Reset()
		return false
	} else {
		return true
	}
}

// implements
func (reader *LineRecordReader) GetCuttentKey() (interface{}, error) {
	return reader.key, nil
}

// implements
func (reader *LineRecordReader) GetCurrentValue() (interface{}, error) {
	return reader.value, nil
}

// others
/*
	Strip BOM(Byte Order Mark)
	Text only support UTF-8, we only need to check UTF-8 BOM
	(0xEF,0xBB,0xBF) at the start of the text stream.
*/
func (reader *LineRecordReader) skipUtfByteOrderMark() (int32, error) {
	var newMaxLineLength int32 = int32(
		util.MinInt64(int64(reader.maxLineLength)+3, math.MaxInt64))

	newSize, err := reader.in.ReadLine(&reader.value, newMaxLineLength, reader.maxBytesToConsume(reader.pos))
	if err != nil {
		return 0, err
	}

	reader.pos += int64(newSize)

	textBytes := []byte(reader.value.String())
	textLength := reader.value.Len()

	// check BOM
	if textLength > 3 && textBytes[0] == 0xEF && textBytes[1] == 0xBB && textBytes[2] == 0xBF {
		textLength -= 3
		newSize -= 3
		if textLength > 0 {
			reader.value.Write(textBytes[3:])
		} else {
			reader.value.Reset() //TODO
		}
	}

	return newSize, nil
}

func (reader *LineRecordReader) maxBytesToConsume(pos int64) int32 {
	return int32(
		util.MaxInt64(
			util.MinInt64(reader.end-pos, math.MaxInt64),
			int64(reader.maxLineLength)))
}

func (reader *LineRecordReader) Close() error {
	return reader.fileIn.Close()
}

// func (reader *LineRecordReader) getFilePosition() (int64, error){

// }
