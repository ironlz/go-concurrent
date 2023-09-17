package concurrent

type Semaphore interface {
	Acquire()
	Release()
}

func NewSemaphore(permits int) Semaphore {
	return &semaphoreImpl{
		permits:   permits,
		blockChan: make(chan interface{}, permits),
	}
}

type semaphoreImpl struct {
	permits   int
	blockChan chan interface{}
}

func (s *semaphoreImpl) Acquire() {
	s.blockChan <- nil
}

func (s *semaphoreImpl) Release() {
	<-s.blockChan
}
