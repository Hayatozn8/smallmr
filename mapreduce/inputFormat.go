package mapreduce

//package inputformat
import (
	"github.com/Hayatozn8/smallmr/input/recordReader"
	intpuSplit "github.com/Hayatozn8/smallmr/input/split"
	//"github.com/Hayatozn8/smallmr/mapreduce"
)

type InputFormat interface {
	GetSplits(job JobContext) ([]intpuSplit.InputSplit, error)
	createRecordReader(split intpuSplit.InputSplit, context TaskContext) recordReader.RecordReader
}
