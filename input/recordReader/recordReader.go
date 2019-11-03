package recordReader

import (
	inputSplit "github.com/Hayatozn8/smallmr/input/split"
)

type RecordReader interface {
	Initalize(split inputSplit.InputSplit) error
	NextKeyValue() (bool, error)
	GetCuttentKey() (interface{}, error)
	GetCurrentValue() (interface{}, error)
}
