package main

import (
	"context"
	"fmt"
)

func countTo(ctx context.Context, max int) <-chan int {
	ch := make(chan int)
	go func() {
	loop:
		for i := 0; i < max; i++ {
			select {
			case <-ctx.Done():
				break loop
			case ch <- i:
			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := countTo(ctx, 10)
	for i := range ch {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
	cancel()
}
