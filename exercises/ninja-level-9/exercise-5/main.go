package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {

	var counter int64
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Func Counter:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
