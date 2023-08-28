package main

import "fmt"

/*
2. Create a function that launches two goroutines.
Each goroutine writes 10 numbers to its own channel.
Use a +for-select+ loop to read from both channels, printing out the number and the goroutine that wrote the value.
Make sure that your function exits after all values are read and that none of your goroutines leak.
*/
func ProcessData() {
	ch := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	go func() {
		for i := 0; i < 10; i++ {
			ch2 <- i*100 + 1
		}
		close(ch2)
	}()
	// once a channel is closed, ok will return false. Use that information to set the channel variable to nil,
	// disabling the case. When both cases are disabled, you are done.
	count := 2
	for count != 0 {
		select {
		case v, ok := <-ch:
			if !ok {
				ch = nil
				count--
				break
			}
			fmt.Println(v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				count--
				break
			}
			fmt.Println(v)
		}
	}
}

func main() {
	ProcessData()
}
