package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	counter := 0
	wg.Add(100)

	//	var mlock sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			//			mlock.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			//			mlock.Unlock()
			fmt.Println("Func Counter:", counter)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
