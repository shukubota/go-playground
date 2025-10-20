package app

import (
	"fmt"
	"time"
)

func Run() error {
	fmt.Println("aaa")
	time.Sleep(5 * time.Second)
	fmt.Println("bbb")
	return nil
}
