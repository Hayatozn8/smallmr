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
	"strings"
	//"reflect"
	"github.com/Hayatozn8/smallmr/input/split"
	"github.com/Hayatozn8/smallmr/input/recordReader"
)

func main() {
	path := `test.txt`
	// var start int64= 20
	// fileInfo, _ := os.Stat(p)
	// fmt.Println(fileInfo.)
	// offs := []int64{0, 45, 75}
	offs := []int64{0, 60, 120}
	status := make(chan struct{}, 3)
	for _, off := range offs {
	 	go func (path string, start int64) {
			fsplit := split.NewFileSplit(path, start, 10)

			reader := recordReader.NewLineRecordReader(nil, 100)

			err := reader.Initalize(fsplit)
			if err !=nil{
				return
			}

			for reader.NextKeyValue(){
				ki,_ := reader.GetCuttentKey()
				vi,_ := reader.GetCurrentValue()

				// k := ki.(int64)
				v := vi.(strings.Builder)
				fmt.Println("start =", start, ", ",ki, ", ", v.String())
			}

			reader.Close()
			status <- struct{}{}
		}(path, off)
	}

	<-status
	<-status
	<-status
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
