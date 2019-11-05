package mapreduce

import (
	"fmt"
	"strings"
	"github.com/Hayatozn8/smallmr/config"
	"github.com/Hayatozn8/smallmr/split"
	"sync"
)

type Job struct{
	//inputFormatClass
	inputPaths []string
	conf *config.Configuration
}

func NewJob(conf *config.Configuration) JobContext{
	return &Job{
		conf:conf,
	}
}

func (job *Job) Submit() error{
	// use inoutformat
	// TODO : how to reflect????
	inputFormat := NewFileInputFormat()
	// 	GetSplits
	splitList, err := inputFormat.GetSplits(job)
	if err != nil{
		return err //TODO
	}

	// TODO: wait end
	var wg sync.WaitGroup
	wg.Add(len(splitList))
	fmt.Println("split count = ", len(splitList))
	fmt.Println("all split count = ", splitList)
	// switch to taskContext
	for i, splitInfo := range splitList{
		go func(splitInfo split.InputSplit, index int){
			defer wg.Done()
			reader := inputFormat.CreateRecordReader(splitInfo, job)
			err := reader.Initialize(splitInfo, job)
			if err != nil{
				fmt.Println("index=", index, "Job reader.Initialize:", err)
			}
					
			for reader.NextKeyValue(){
				// TODO
				ki,_ := reader.GetCuttentKey()
				vi,_ := reader.GetCurrentValue()

				v := vi.(strings.Builder)
				fmt.Println("index =", index, ", ",ki, ", ", v.String())
			}
			ki,_ := reader.GetCuttentKey()
			vi,_ := reader.GetCurrentValue()

			v := vi.(strings.Builder)
			fmt.Println("index =", index, ", ",ki, ", ", v.String())
		}(splitInfo, i)
	}

	wg.Wait()
	return nil
}

func (job *Job) SetInputFormatClass(InputFormat) {

}

//TODO
func (job *Job) GetNumReduceTasks() int{
	return 0
}

//TODO
func (job *Job) GetInputFormatClass() (InputFormat, error){
	return nil, nil
}

//TODO
func (job *Job) GetMapperClass() (Mapper, error){
	return nil,nil
}

//TODO
func (job *Job) GetReducerClass() (Reduce, error){
	return nil, nil
}

func (job *Job) SetInputPaths(paths ...string){
	job.inputPaths = paths
}

func (job *Job) GetInputPaths() []string{
	return job.inputPaths
}

func (job *Job) GetConfiguration() *config.Configuration{
	return job.conf
}


//TODO 
func (job *Job)NextKeyValue() (bool, error){
	return false,nil
}
func (job *Job)GetCuttentKey() (interface{}, error){
	return nil, nil
}
func (job *Job)GetCurrentValue() (interface{}, error){
	return nil, nil
}
func (job *Job)Write(outKey interface{}, outValue interface{}){
	
}
