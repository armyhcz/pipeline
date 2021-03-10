package end

import (
	"duomai/pipeline/internal/next"
	"sync"
)

type Result struct {
	Id      int
	Status  bool
	Message string
}

func End(c <-chan *next.Result, concurrency int) chan *Result {
	wg := &sync.WaitGroup{}
	ch := make(chan *Result, 1024)
	for i := 1; i < concurrency; i++ {
		wg.Add(1)
		go doEnd(c, ch, wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}

func doEnd(c <-chan *next.Result, result chan<- *Result, wg *sync.WaitGroup) {
	for v := range c {
		result <- &Result{
			Id:      v.Id,
			Status:  true,
			Message: "success",
		}
	}
	wg.Done()
}
