package mapreduce

type Mapper interface {
	SetUp(context MapContext) error
	Map(key interface{}, value interface{}, context MapContext) error
	CleanUp(context MapContext) error
	Run(context MapContext) error
}
