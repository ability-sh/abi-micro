package micro

type Executor interface {
	Exec(ctx Context, name string, data interface{}) (interface{}, error)
}
