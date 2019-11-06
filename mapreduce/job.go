package mapreduce

import (
	"fmt"
	"strings"
	"sync"
	"reflect"

	"github.com/Hayatozn8/smallmr/config"
	"github.com/Hayatozn8/smallmr/split"
)

type Job struct {
	inputFormatObj InputFormat //TODO should be type, (delete)
	inputFormatClass reflect.Type //system should set defalut inputFormatClass when execute Submit()
	mapperClass 	reflect.Type
	reducerClass 	reflect.Type
	inputPaths     []string
	conf           *config.Configuration
	fileSplits     []split.InputSplit
}

func NewJob(conf *config.Configuration) *Job {
	return &Job{
		conf: conf,
	}
}

//TODO
// interface implement
func (job *Job) GetNumReduceTasks() int {
	return 0
}

//TODO
// interface implement
func (job *Job) GetInputFormatClass() (InputFormat, error) {
	return nil, nil
}

//TODO
func (job *Job) SetInputFormatClass(format InputFormat) {
	job.inputFormatObj = format
}

//TODO
// interface implement
func (job *Job) GetMapperClass() (reflect.Type, error) {
	return nil, nil
}

//TODO job or config ????
func (job *Job) SetMapperClass(nilMapObj Mapper) {
	t := reflect.TypeOf(nilMapObj)
	if t.Kind() == reflect.Ptr{
		t = t.Elem()
	}

	// TODO:
	// reflect.New(t).Interface()
	// check func(x *X) and func(x X) 
	job.mapperClass = t
}

//TODO
// interface implement
func (job *Job) GetReducerClass() (reflect.Type, error) {
	return nil, nil
}

// interface implement
func (job *Job) GetConfiguration() *config.Configuration {
	return job.conf
}

func (job *Job) Submit() error {
	// use inoutformat
	// TODO : how to reflect????
	// inputFormat := NewFileInputFormat()

	// 	GetSplits
	splitList, err := job.inputFormatObj.GetSplits(job)
	if err != nil {
		return err //TODO
	}

	// TODO:len(split) == 0 ????
	job.fileSplits = splitList

	// TODO: wait end
	var wg sync.WaitGroup
	wg.Add(len(job.fileSplits))
	fmt.Println("split count = ", len(job.fileSplits))
	fmt.Println("all split count = ", job.fileSplits)

	job.fileSplits = MAX_TASK_COUNT

	NewTaskTracker()



	// switch to taskContext
	for i, splitInfo := range job.fileSplits {
		go func(splitInfo split.InputSplit, index int) {
			defer wg.Done()
			reader := job.inputFormatObj.CreateRecordReader(splitInfo, job)
			err := reader.Initialize(splitInfo, job)
			if err != nil {
				fmt.Println("index=", index, "Job reader.Initialize:", err)
			}

			for reader.NextKeyValue() {
				// TODO
				// ki,_ := reader.GetCuttentKey()
				// vi,_ := reader.GetCurrentValue()

				// v := vi.(strings.Builder)
				// fmt.Println("index =", index, ", ",ki, ", ", v.String())
			}
			ki, _ := reader.GetCuttentKey()
			vi, _ := reader.GetCurrentValue()

			v := vi.(strings.Builder)
			fmt.Println("index =", index, ", ", ki, ", ", v.String())
		}(splitInfo, i)
	}

	wg.Wait()
	return nil
}





//TODO ????????????????????????///
func (job *Job) NextKeyValue() (bool, error) {
	return false, nil
}
func (job *Job) GetCuttentKey() (interface{}, error) {
	return nil, nil
}
func (job *Job) GetCurrentValue() (interface{}, error) {
	return nil, nil
}
func (job *Job) Write(outKey interface{}, outValue interface{}) {

}
