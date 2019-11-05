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
