package main

import (
	"fmt"
	"sync"
)

type counter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *counter) Inc(wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	c.v["wakuwaku"]++
	c.mu.Unlock()
}

func main() {
	fmt.Println("---------9999")
	wg := sync.WaitGroup{}
	c := counter{v: make(map[string]int)}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go c.Inc(&wg)
	}

	fmt.Println("%+v", c.v)
}
