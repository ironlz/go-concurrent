package executor

type AsyncResult interface {
	Result() (interface{}, error)
	Error() error
	IsCompleted() bool
}
