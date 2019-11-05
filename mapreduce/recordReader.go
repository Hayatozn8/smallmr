package mapreduce

import (
	inputSplit "github.com/Hayatozn8/smallmr/split"

)

type RecordReader interface {
	Initialize(split inputSplit.InputSplit, context TaskContext) error
	NextKeyValue() bool
	GetCuttentKey() (interface{}, error)
	GetCurrentValue() (interface{}, error)
	Err() error
	Close() error
}
