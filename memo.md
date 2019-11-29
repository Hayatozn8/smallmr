* func method(long maxSplitSize)
* 加载文件
* 获取文件大小
* 计算分片数量
    * 对文件进行逻辑划分
* go func 读取各个分片
    
    * lineRecordReader --读取start～end的内容
        * LineReader --从缓存中读取某一行
        * 读取的每一行内容放入chan
    * Mapper
        * 转化为csv格式
        * 构造key value
* 创建CsvInputFormat来处理每一行的解析和切片设定
********************************************
package main
import (
	//"log"
	"fmt"
	"os"
	//"strings"
	//"reflect"
	"github.com/Hayatozn8/smallmr/config"
	// "github.com/Hayatozn8/smallmr/util"
	"github.com/Hayatozn8/smallmr/mapreduce"
	"time"
)

func main(){
	start := time.Now().Unix()

	path := `C:\Users\liujinsuo\Desktop\mygo\src\learn\test.txt`
	fileInfo, _ := os.Stat(path)
	fmt.Println(fileInfo.Size())

	conf := config.NewConfiguration()
	// conf.SetInt64(config.SPLIT_MAXSIZE, util.STORAGE_UNIT_1MB * 100)
	job:=mapreduce.NewJob(conf)
	// job.
	job.SetInputPaths(path)
	job.Submit()

	end := time.Now().Unix()
	fmt.Print(end - start)
	
}


aaaaaaaaaaaaaaaaaaa
bbbbbbbbbbbbbbbbbbb
ccccccccccccccccccc
ddddddddddddddddddd
eeeeeeeeeeeeeeeeeee
fffffffffffffffffff
ggggggggggggggggggg
hhhhhhhhhhhhhhhhhhh
iiiiiiiiiiiiiiiiiii



////////////////////////////////////////////////
package main

import (
	"github.com/Hayatozn8/excelize"
	"fmt"
)

func main(){
	job := NewJobContextImpl(``, "", 4)
	job.Submit()
}
//////////////////////////////////////////////////////////////
type Context interface{
	NextKeyValue() bool
	GetCuttentKey() interface{}
	GetCurrentValue() interface{}
	// Write(outKey interface{}, outValue interface{})
}
//////////////////////////////////////////////////////////////
type JobContext interface{
	Submit()
}

type JobContextImpl struct{
	filePath string
	sheetName string
	effictRowStartIndex uint
}

func NewJobContextImpl(filePath string, sheetName string, effictRowStartIndex uint) JobContext{
	return &JobContextImpl{
		filePath:filePath,
		sheetName:sheetName,
		effictRowStartIndex:effictRowStartIndex,
	}
}

func(job *JobContextImpl) Submit(){
	inputFormat, _ := NewExcelSheetLayoutReader(job.filePath, job.sheetName, job.effictRowStartIndex)
	mapContext := NewMapContext(inputFormat)
	mapper := NewMapper(mapContext)
	mapper.Run()
}


//////////////////////////////////////////////////////////////
type MapContext struct{
	inputFormat InputFormat
}

func NewMapContext(inputFormat InputFormat) Context{
	return &MapContext{
		inputFormat : inputFormat,
	}
}

func(mc *MapContext) NextKeyValue() bool {
	return mc.inputFormat.NextKeyValue()
}

func(mc *MapContext) GetCuttentKey() interface{} {
	return mc.inputFormat.GetCuttentKey()
}

func(mc *MapContext) GetCurrentValue() interface{} {
	return mc.inputFormat.GetCurrentValue()
}

//////////////////////////////////////////////////////////////
type MapperInterface interface{
	// SetUp(context MapContext) error
	// Map(key interface{}, value interface{}, context MapContext) error
	// CleanUp(context MapContext) error
	// Run(context MapContext) error
	
	Run() error
	Map(key interface{}, value interface{}) error
}

type Mapper struct{
	context Context
}

func NewMapper(context Context) MapperInterface{
	return &Mapper{
		context : context,
	}
}

func(m *Mapper) Run() error {
	for m.context.NextKeyValue() {
		err := m.Map(m.context.GetCuttentKey(), m.context.GetCurrentValue())
		if err != nil{
			return err
		}
	}

	return nil //TODO
}

func(m *Mapper) Map(key interface{}, value interface{}) error {
	a := value.([][]string)
	for _, aa := range a{
		for _, ax := range aa{
			fmt.Print(ax, ", ")
		}
	}
	fmt.Println()

	return nil //TODO
}

//////////////////////////////////////////////////////////////
type InputFormat interface{
	NextKeyValue() bool
	GetCuttentKey() (interface{})
	GetCurrentValue() (interface{})
}

type ExcelSheetLayoutReader struct{
	filePath string
	rowReader *excelize.Rows
	currentKey int
	currentValue [][]string
}

func NewExcelSheetLayoutReader(filePath string, sheetName string, effictRowStartIndex uint) (InputFormat, error){
	f, err := excelize.OpenFile(filePath)
	if err != nil{
		return nil, err
	}

	rowReader, err := f.RowsFromStartIndex(sheetName, int(effictRowStartIndex))

	if err != nil{
		return nil, err
	}

	return &ExcelSheetLayoutReader{
		filePath : filePath,
		rowReader : rowReader,
	}, nil
}

func(reader *ExcelSheetLayoutReader) NextKeyValue() bool{
	if reader.rowReader.Next() {
		row, _ := reader.rowReader.ColumnsWithIndexes("B","C","D","E","F","G","K")
		reader.currentValue = make([][]string,2)
		reader.currentValue = append(reader.currentValue, row)
		return true
	} else {
		return false
	}
}

func(reader *ExcelSheetLayoutReader) GetCuttentKey() (interface{}){
	return reader.currentKey
}

func(reader *ExcelSheetLayoutReader) GetCurrentValue() (interface{}){
	return reader.currentValue
}
