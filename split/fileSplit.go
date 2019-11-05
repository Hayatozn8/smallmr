package split

import (
	"strconv"
)

type FileSplit struct {
	path   string
	start  int64
	length int64
	//TODO
	//ignore host and hostInfors
	//hosts []string
	//hostInfos
}

func NewFileSplit(path string, start int64, length int64) InputSplit {
	return &FileSplit{
		path:   path,
		start:  start,
		length: length,
	}
}

func (fsplit *FileSplit) GetPath() string {
	return fsplit.path
}

func (fsplit *FileSplit) GetStart() int64 {
	return fsplit.start
}

func (fsplit *FileSplit) GetLength() int64 {
	return fsplit.length
}

func (fsplit *FileSplit) String() string {
	// analyze file TODO
	return fsplit.path + ":" +
		strconv.FormatInt(fsplit.start, 10) +
		"+" + strconv.FormatInt(fsplit.length, 10)
}
