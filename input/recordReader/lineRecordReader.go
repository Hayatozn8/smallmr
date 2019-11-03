package recordReader

import (
	"math"
	"os"

	inputSplit "github.com/Hayatozn8/smallmr/input/split"
	"github.com/Hayatozn8/smallmr/util"
)

type LineRecordReader struct {
	start                int64
	pos                  int64
	end                  int64
	recordDelimiterBytes []byte
	in                   util.LineReader
	fileIn               *os.File
	key                  int64
	value                string //TODO
	maxLineLength        int32
	// ignore zip file TODO

}

func NewLineRecordReader(recordDelimiterBytes []byte) RecordReader {
	return &LineRecordReader{
		recordDelimiterBytes: recordDelimiterBytes,
	}
}

// implements
func (reader *LineRecordReader) Initalize(split inputSplit.InputSplit) error {
	fsplit, error := split.(inputSplit.FileSplit)
	reader.start = fsplit.GetStart()
	reader.end = start + fsplit.GetLength()

	//reader.in = util.NewBaseLineReader()
	reader.fileIn, _ = os.Open(fsplit.GetPath())

	if reader.start != 0 {
		reader.start += in.readLine(reader.value, 0, reader.maxBytesToConsume(reader.start))
	}
	return nil //TODO
}

// implements
func (reader *LineRecordReader) NextKeyValue() (bool, error) {
	reader.key = reader.pos

	var newSize int32 = 0

	for reader.pos <= reader.end {
		if reader.pos == 0 {
			newSize = reader.skipUtfByteOrderMark()
		} else {
			newSize, err := reader.in.ReadLine(&reader.value, reader.maxLineLength, reader.maxBytesToConsume(reader.pos))
			reader.pos += int64(newSize)
		}

		// EOF:newSize==0,
		if newSize == 0 || newSize < reader.maxLineLength {
			break
		}
	}

	if newSize == 0 {
		//reader.key =
		return false, nil
	} else {
		return true, nil
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
	reader.pos += int64(newSize)

	textLength := len(reader.value)

	// check BOM
	if textLength > 3 && reader.value[0] == 0xEF && reader.value[1] == 0xBB && reader.value[2] == 0xBF {
		textLength -= 3
		newSize -= 3
		if textLength > 0 {
			reader.value = reader.value[3:]
		} else {
			reader.value = "" //TODO
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

// func (reader *LineRecordReader) getFilePosition() (int64, error){

// }
