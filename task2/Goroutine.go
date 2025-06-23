package main

import (
	"fmt"
	"sync"
	"time"
)

// 打印1-10的奇偶数字
func printNum() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i = i + 2 {
			fmt.Println("奇数打印: ", i)
		}
	}()

	//wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i = i + 2 {
			fmt.Println("偶数打印: ", i)
		}
	}()
	// 等待所有goroutine完成
	wg.Wait()
}

// Goroutine 第二题
type Task func()

func excuteTasks(tasks []Task) {
	var wg sync.WaitGroup
	wg.Add(3)
	for i, task := range tasks {
		go func(i int, t Task) {
			defer wg.Done()
			start := time.Now()
			t() // 执行任务
			duration := time.Since(start)
			fmt.Println("任务", i+1, "耗时", duration, "秒")
		}(i, task)
	}
	wg.Wait() // 等待所有任务完成
}

func main() {
	//Goroutine 第一题
	printNum()

	// Goroutine 第二题
	tasks := []Task{
		func() {
			time.Sleep(time.Second * 1)
			fmt.Println("任务一")
		},
		func() {
			time.Sleep(time.Second * 2)
			fmt.Println("任务二")
		},
		func() {
			time.Sleep(time.Second * 4)
			fmt.Println("任务三")
		},
	}
	excuteTasks(tasks)

}
