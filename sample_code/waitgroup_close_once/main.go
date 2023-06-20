package main

import (
	"fmt"
	"sync"
)

func processAndGather[T, R any](in <-chan T, processor func(T) R, num int) []R {
	out := make(chan R, num)
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			for v := range in {
				out <- processor(v)
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	var result []R
	for v := range out {
		result = append(result, v)
	}
	return result
}

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
		}
		close(ch)
	}()
	results := processAndGather(ch, func(i int) int {
		return i * 2
	}, 3)
	fmt.Println(results)
}
