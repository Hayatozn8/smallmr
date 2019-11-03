package mapreduce

//package inputformat

import (
	intpuSplit "github.com/Hayatozn8/smallmr/input/split"
	"github.com/Hayatozn8/smallmr/mapreduce"
)

const (
	SPLIT_SLOP = 1.1
	MInSplitSize
)

//single file
type FileInputFormat struct {
}

// not have PathFilter
func (fif *FileInputFormat) GetSplits(mapreduce.JobContext context) ([]intpuSplit.InputSplit, error) {

	var splits []intpuSplit.InputSplit
	return nil, nil
}

//createRecordReader(InputSplit split, TaskAttemptContext context) (RecordReader<K,V>, error )
