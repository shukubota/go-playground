package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("========aaa")
	//useWaitGroup()
	useChannel()
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

func useChannel() {
	ch1 := make(chan *testResponse)
	ch2 := make(chan *testResponse)
	go test1WithChan(ch1)
	go test2WithChan(ch2)

	res1 := <-ch1
	fmt.Println(res1)
	res2 := <-ch2
	fmt.Println(res2)

}

type testResponse struct {
	value any
}

func test1() *testResponse {
	fmt.Println("test1 start")
	time.Sleep(1 * time.Second)
	fmt.Println("test1 end")

	res := &testResponse{
		value: "end test1",
	}

	return res
}

func test2() *testResponse {
	fmt.Println("test2 start")
	time.Sleep(2 * time.Second)
	fmt.Println("test2 end")

	res := &testResponse{
		value: "end test2",
	}

	return res
}

func test1WithChan(ch chan *testResponse) {
	fmt.Println("test1 start")
	time.Sleep(1 * time.Second)
	fmt.Println("test1 end")
	ch <- &testResponse{
		value: "test1 end",
	}
}

func test2WithChan(ch chan *testResponse) {
	fmt.Println("test2 start")
	time.Sleep(2 * time.Second)
	fmt.Println("test2 end")
	ch <- &testResponse{
		value: "test2 end",
	}
}
