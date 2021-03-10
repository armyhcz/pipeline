package start

import "context"

type Result struct {
	Id int
}

func Start(ctx context.Context, concurrency int) chan *Result {
	ch := make(chan *Result, 1024)
	for i := 0; i < concurrency; i++ {
		go doStart(ctx, ch)
	}
	return ch
}

func doStart(ctx context.Context, result chan<- *Result) {
	// doStart
}
