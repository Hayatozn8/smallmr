package mapreduce

type Reduce interface {
	SetUp(context ReduceContext) error
	Reduce(key interface{}, values []interface{}, context ReduceContext) error
	CleanUp(context ReduceContext) error
}

type ReduceRunner interface {
	Run(context ReduceContext) error
}

type BaseReduce struct {
}

func (r *BaseReduce) SetUp(context ReduceContext) error {
	return nil
}

func (r *BaseReduce) Reduce(key interface{}, values []interface{}, context ReduceContext) error {
	for _, value := range values {
		context.Write(key, value)
	}
	return nil
}

func (r *BaseReduce) CleanUp(context ReduceContext) error {
	return nil
}
