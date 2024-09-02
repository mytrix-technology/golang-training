package contexts

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		time.Sleep(1 * time.Second)
		fmt.Println("doing work...")
	}
}

func main() {
	bgCtx := context.Background()
	innerCtx, cancel := context.WithCancel(bgCtx)

	go doWork(innerCtx)         // call goroutine
	time.Sleep(3 * time.Second) // do work in main

	// well, if `doWork` is still not done, just cancel it
	cancel()
}
