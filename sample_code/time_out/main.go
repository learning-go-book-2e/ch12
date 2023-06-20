package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	result, err := timeLimit()
	fmt.Println(result, err)
}

func timeLimit() (int, error) {
	out := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go func() {
		out <- doSomeWork()
	}()
	select {
	case result := <-out:
		return result, nil
	case <-ctx.Done():
		return 0, errors.New("work timed out")
	}
}

func doSomeWork() int {
	if x := rand.Int(); x%2 == 0 {
		return x
	} else {
		time.Sleep(10 * time.Second)
		return 100
	}
}
