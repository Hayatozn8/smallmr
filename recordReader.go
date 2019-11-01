package smallMR

type recordReader interface{
	//接收切片信息，对文件进行指定切割读取
	nextKeyValue() (bool, error)
}
