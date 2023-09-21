package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	out := make(chan int, 3)
	go pipeStart(out)
	dispatcher(out)
}

func pipeStart(out chan int) {
	for i := 0; i < 110; i++ {
		fmt.Println("Send ", i)
		out <- i
		time.Sleep(5 * time.Millisecond)
	}
	close(out)
}

func dispatcher(input chan int) {
	var wg sync.WaitGroup
	for {
		for i := 0; i < 3; i++ {
			if num, ok := <-input; ok {
				wg.Add(1)
				go func(num int) {
					worker(num)
					wg.Done()
				}(num)
			} else {
				wg.Wait()
				return
			}
		}
		wg.Wait()
	}
}

func worker(in int) {
	fmt.Println("		StartProcess: ", in)
	time.Sleep(20 * time.Millisecond)
	fmt.Println("		DoneProcess: ", in)
}
