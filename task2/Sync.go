package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// 锁机制 第一题
	var mutex sync.Mutex
	var wg sync.WaitGroup
	var counter int

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println("counter的最终值：", counter)

	// 锁机制 第二题

	var wg1 sync.WaitGroup
	var counter1 int32

	wg1.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg1.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&counter1, 1)
			}
		}()
	}

	wg1.Wait()
	fmt.Println("counter1的最终值：", counter1)
}
