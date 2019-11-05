package mapreduce

//package inputformat
import (
	intpuSplit "github.com/Hayatozn8/smallmr/split"
	//"github.com/Hayatozn8/smallmr/mapreduce"
)

type InputFormat interface {
	GetSplits(job JobContext) ([]intpuSplit.InputSplit, error)
	CreateRecordReader(split intpuSplit.InputSplit, context TaskContext) RecordReader
}
