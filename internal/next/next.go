package next

import (
	"duomai/pipeline/internal/start"
	"sync"
)

type Result struct {
	Id int
}

func Next(task <-chan *start.Result, concurrency int) chan *Result {
	wg := &sync.WaitGroup{}
	ch := make(chan *Result, 1024)
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go doNext(task, ch, wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}

func doNext(task <-chan *start.Result, result chan<- *Result, wg *sync.WaitGroup) {
	for v := range task {
		result <- &Result{
			Id: v.Id,
		}
	}

	wg.Done()
}
