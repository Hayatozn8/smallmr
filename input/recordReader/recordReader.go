package recordReader

import (
	inputSplit "github.com/Hayatozn8/smallmr/input/split"
)

type RecordReader interface {
	Initalize(split inputSplit.InputSplit) error
	NextKeyValue() bool
	GetCuttentKey() (interface{}, error)
	GetCurrentValue() (interface{}, error)
	Err() error
	Close() error
}
