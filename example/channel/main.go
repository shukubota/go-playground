package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("channel")

	ch1 := make(chan int)

	ch2 := make(chan int)

	close(ch1)
	close(ch2)

	var c1count, c2count int

	for i := 0; i < 1000; i++ {
		select {
		case i, more := <-ch1:
			fmt.Println("ch1", i, more)
			c1count++
		case i, more := <-ch2:
			fmt.Println("ch2", i, more)
			c2count++
		}
	}

	fmt.Println("ch1 count", c1count)
	fmt.Println("ch2 count", c2count)

	//c := make(chan string)
	//
	//go do1(c)
	//go do2(c)
	//
	//msg := <-c
	//fmt.Println(msg)
	//msg = <-c
	//fmt.Println(msg)
}

func do1(c chan string) {
	fmt.Println("start do1")
	time.Sleep(time.Second * 2)
	fmt.Println("end do1")
	c <- "do1 end from channel"
}

func do2(c chan string) {
	fmt.Println("start do2")
	time.Sleep(time.Second * 1)
	fmt.Println("end do2")
	c <- "do2 end from channel"
}
