package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

func main() {
	fmt.Println("========aaa")
	//useErrorGroup()
	//useWaitGroup()
	//useChannel()
}

func useErrorGroup() {
	fmt.Println("===========iii")
	eg, ctx := errgroup.WithContext(context.TODO())

	fmt.Println(ctx)

	for i := 0; i < 10; i++ {
		i := i
		fmt.Println(i)
		fmt.Println("======before eg.Go")
		eg.Go(func() error {
			fmt.Println(i)
			if i == 2 || i == 3 {
				return fmt.Errorf("error: hoge i: %d", i)
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Printf("error :%+v\n", err)
	}
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
