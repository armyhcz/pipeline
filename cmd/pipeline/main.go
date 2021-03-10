package main

import (
	"context"
	"duomai/pipeline/internal/end"
	"duomai/pipeline/internal/next"
	"duomai/pipeline/internal/start"
	"log"
)

func main() {
	ctx := context.Background()
	s := start.Start(ctx, 1)
	n := next.Next(s, 8)
	e := end.End(n, 8)
	for message := range e {
		if !message.Status {
			log.Printf("id: %d; error: %s", message.Id, message.Message)
		}
	}
}
