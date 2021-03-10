package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// WaitGroup
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()

	// Channels
	ch := make(chan int)
	go func() {
		ch <- doSomething(5)
	}()
	fmt.Println(<-ch)

	// Race Condition
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	const gs = 100
	var wg sync.WaitGroup

	wg.Add(gs)

	var mutex sync.Mutex

	counter := 0
	for i := 0; i < gs; i++ {
		go func() {
			mutex.Lock()

			v := counter
			// time.sleep(time.Second) - could use time to delay also!
			runtime.Gosched() // tells the CPU to 'go run something else' or yield control
			v++
			counter = v

			mutex.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines\t", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Counter:", counter)
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}

func doSomething(x int) int {
	return x * 5
}
