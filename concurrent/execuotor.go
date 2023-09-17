package concurrent

type ExecutorService interface {
}

type Runnable interface {
	Run()
}

type Callable interface {
	Call() (result interface{}, err error)
}

func NewExecutorService(parallelism int, resultChan chan interface{}) ExecutorService {
	return executor{
		semaphore: NewSemaphore(parallelism),
	}
}

type executor struct {
	semaphore Semaphore
}

func (e *executor) Submit(callable Callable) {
	go callable.Call()
}
