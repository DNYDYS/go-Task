package main

import (
	"fmt"
	"sync"
)

// Channel 第一题
func producer(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		//fmt.Println("生产: ", i)
	}
	close(ch)
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("消费：", v)
	}

}

// Channel 第二题
func producer2(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i
		//fmt.Println("生产: ", i)
	}
	fmt.Println("发送结束")
	close(ch)
}

func consumer2(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("消费：", v)
	}

}

func main() {
	//Channel 第一题
	var wg sync.WaitGroup
	var ch = make(chan int)
	wg.Add(1)
	go producer(ch)
	go consumer(ch, &wg)
	wg.Wait()

	//Channel 第二题 有缓冲
	var wg1 sync.WaitGroup
	var ch1 = make(chan int, 100)
	wg1.Add(1)
	go producer2(ch1)
	go consumer2(ch1, &wg1)
	wg1.Wait()

}
