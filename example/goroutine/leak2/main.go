package main

import (
	"fmt"
	"sync"
)

// Q2: 下記のコードはゴルーチンリークが発生しています。原因を特定し、修正してください。
// ただし※1のreadStreamチャネルを読み出す回数は変更しないでください。
func main() {
	var wg sync.WaitGroup
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for i := 0; i < 5; i++ {
				select {
				case r := <-randStream:
					fmt.Println(r)

				case <-ctx.Done():

				}
				//randStream <- i
			}
		}()
		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream) // ※1
	}

	wg.Wait()
	fmt.Println("Done")
}
