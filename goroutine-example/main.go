package main

import (
	"fmt"
	"time"
)

var v = 0

func main() {
	ch1 := make(chan int)
	//ch2 := make(chan string)
	fmt.Println("--------main")
	go goroutine1(ch1)
	//go goroutine2(ch1)

	//var x int
	//var y string
	v, ok := <-ch1
	fmt.Println(v)
	fmt.Println(ok)
	v, ok = <-ch1
	fmt.Println(v)
	fmt.Println(ok)
	fmt.Println("-------000")
}

func goroutine1(ch chan int) {
	time.Sleep(time.Second * 2)
	ch <- 1000
	time.Sleep(time.Second * 2)
	ch <- 200
	fmt.Println("--------pppp")
}

func goroutine2(ch chan int) {
	ch <- 333
}
