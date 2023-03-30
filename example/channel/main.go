package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("channel")
	c := make(chan string)

	go do1(c)
	go do2(c)

	msg := <-c
	fmt.Println(msg)
	msg = <-c
	fmt.Println(msg)
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
