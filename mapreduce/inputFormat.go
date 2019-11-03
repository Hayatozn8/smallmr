package mapreduce

//package inputformat
import (
	intpuSplit "github.com/Hayatozn8/smallmr/input/split"
	//"github.com/Hayatozn8/smallmr/mapreduce"
)

type InputFormat interface {
	GetSplits(context JobContext) ([]intpuSplit.InputSplit, error)
	//createRecordReader(InputSplit split, TaskAttemptContext context) (RecordReader<K,V>, error )
}
