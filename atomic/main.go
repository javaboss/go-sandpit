package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
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

	var counter int64

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			runtime.Gosched() // tells the CPU to 'go run something else' or yield control - not really needed here!
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
