package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("========aaa")
	useWaitGroup()
}

func useWaitGroup() {
	wg := sync.WaitGroup{}

	go func() {
		test1()
		wg.Done()
	}()
	wg.Add(1)

	go func() {
		test2()
		wg.Done()
	}()
	wg.Add(1)

	wg.Wait()
}

func test1() {
	fmt.Println("test1 start")
	time.Sleep(1 * time.Second)
	fmt.Println("test1 end")
}

func test2() {
	fmt.Println("test2 start")
	time.Sleep(2 * time.Second)
	fmt.Println("test2 end")
}
