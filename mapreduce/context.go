package mapreduce

import (
	"reflect"
	"github.com/Hayatozn8/smallmr/config"
	"github.com/Hayatozn8/smallmr/split"
)

/*
 submit check
 paths
*/
type JobContext interface {
	GetNumReduceTasks() int
	
	// SetInputPaths(paths ...string)
	// GetInputPaths() []string //TODO
	GetInputFormatClass() (InputFormat, error)
	// SetInputFormatClass(format InputFormat)

	GetMapperClass() (reflect.Type, error)
	// SetMapperClass(nilMapObj Mapper)

	GetReducerClass() (reflect.Type, error)
	
	GetConfiguration() *config.Configuration
	// Submit() error
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
