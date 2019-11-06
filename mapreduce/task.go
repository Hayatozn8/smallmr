package mapreduce

import (
	"errors"
	"fmt"
	"reflect"

	inputSplit "github.com/Hayatozn8/smallmr/split"
	"github.com/Hayatozn8/smallmr/config"
	"github.com/Hayatozn8/smallmr/util"
)
// 1. job.submit
// 	loop splis Count 
// 		create TaskTracker
// 			go TaskTracker.run(Mapper)
// 				HeartCheck Mapper
// 			go TaskTracker.run(Reduce)
// 				HeartCheck Reduce
				
// 3. HeartCheck return:
// 	status:Rnning, End, Err
// 	task Report： task Precent
	
	
// 4. who is live？
// jobTracker
// 	n * HeartCheck

// n * TaskTackert and self.Work

type TaskTracker struct {
	taskSplit 		inputSplit.InputSplit
	taskType 		string
	mapperClass 		reflect.Type
	reducerClass 		reflect.Type
	conf           *config.Configuration
}

func NewTaskTracker(context JobContext, splitID uint32, taskType string) (TaskContext, error){
	// TODO: how to catch err??
	if taskType != util.TASK_TYPE_MAP && taskType != util.TASK_TYPE_REDUCE {
		return nil, errors.New(fmt.Sprintf("NewTaskTracker error: invalid taskClass:[%s]", taskType))
	}

	// TODO: how to catch err??
	mapperClass, err := context.GetMapperClass()
	if err != nil {
		return nil, err
	}

	// TODO: how to catch err??
	reducerClass, err := context.GetReducerClass()
	if err != nil {
		return nil, err
	}

	return &TaskTracker{
		taskSplit:context.,
		taskType:taskType,
		mapperClass:mapperClass,
		reducerClass:reducerClass,
		conf:context.GetConfiguration(),
	}, nil
}


//interface implements
func (this *TaskTracker) GetNumReduceTasks() int{
	return 0 //TODO
}

//interface implements
func (this *TaskTracker) GetInputFormatClass() (InputFormat, error){
	//TODO
	return nil,nil
}

//interface implements
func (this *TaskTracker) GetMapperClass() (reflect.Type, error){
	//TODO
	return nil,nil
}

//interface implements
func (this *TaskTracker) GetReducerClass() (reflect.Type, error){
	//TODO
	return nil,nil
}

//interface implements
func (this *TaskTracker) GetConfiguration() *config.Configuration{
	return this.conf
}

//interface implements
func (this *TaskTracker) NextKeyValue() (bool, error){
	//TODO
	return false, nil
}

//interface implements
func (this *TaskTracker) GetCuttentKey() (interface{}, error){
	//TODO
	return nil, nil
}

//interface implements
func (this *TaskTracker) GetCurrentValue() (interface{}, error){
	//TODO
	return nil, nil
}

//interface implements
func (this *TaskTracker) Write(outKey interface{}, outValue interface{}){
	//TODO
	
}
