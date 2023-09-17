package executor

type Runnable func() (result interface{}, err error)

type ExecutorService interface {
	Submit(r Runnable) AsyncResult
}
