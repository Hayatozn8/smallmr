package recordReader

import (
	inputSplit "github.com/Hayatozn8/smallmr/input/split"
	"github.com/Hayatozn8/smallmr/mapreduce"
)

type RecordReader interface {
	Initalize(split inputSplit.InputSplit, context mapreduce.TaskContext) error
	NextKeyValue() bool
	GetCuttentKey() (interface{}, error)
	GetCurrentValue() (interface{}, error)
	Err() error
	Close() error
}
