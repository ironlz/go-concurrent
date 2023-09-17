package concurrent

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	s := NewSemaphore(10)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			s.Acquire()
			defer s.Release()
			defer wg.Done()
			time.Sleep(time.Second)
			fmt.Println(n)
		}(i)
	}
	wg.Wait()
}
