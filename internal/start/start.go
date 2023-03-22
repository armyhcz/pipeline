package start

import "context"

type Result struct {
	Id int
}

func Start(ctx context.Context, concurrency int) chan *Result {
	ch := make(chan *Result, 1024)
	wg := &sync.WaitGroup{}
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go doStart(ctx, ch)
	}
	
	go func() {
		wg.Wait()
		close(ch)
	}()
	
	return ch
}

func doStart(ctx context.Context, result chan<- *Result) {
	defer wg.Done()
	// doStart
}
