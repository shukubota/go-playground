package main

import (
	"fmt"
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
}

func goroutine1(ch chan int) {
	//time.Sleep(time.Second * 3)
	ch <- 1000
	ch <- 200
}

func goroutine2(ch chan int) {
	ch <- 333
}
