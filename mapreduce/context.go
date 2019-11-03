package mapreduce

import (
	//"github.com/Hayatozn8/smallmr/input/inputformat"
	"github.com/Hayatozn8/smallmr/input/split"
)

type JobContext interface {
	GetNumReduceTasks() int
	GetInputFormatClass() (InputFormat, error)
	GetMapperClass() (Mapper, error)
	GetReducerClass() (Reduce, error)
}

type TaskContext interface {
	JobContext
	NextKeyValue() (bool, error)
	GetCuttentKey() (interface{}, error)
	GetCurrentValue() (interface{}, error)
	Write(outKey interface{}, outValue interface{})
	//TaskAttemptContext
}

type MapContext interface {
	TaskContext
	getIntputSplit() split.InputSplit
}

type ReduceContext interface {
	TaskContext
	nextKey() (bool, error)
	// getValues() ([])
}
