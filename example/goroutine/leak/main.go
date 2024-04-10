package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Q2: 下記のコードはゴルーチンリークが発生しています。原因を特定し、修正してください。
// ただし※1のreadStreamチャネルを読み出す回数は変更しないでください。
func main() {
	var wg sync.WaitGroup

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	newRandStream := func(ctx context.Context, done <-chan any) <-chan int {
		randStream := make(chan int)
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)

			for i := 0; i < 5; i++ {
				select {
				case randStream <- i:
					break
				case <-done:
					fmt.Println("done")
					return
					//case <-ctx.Done():
					//	fmt.Println("ctx.Done")
					//	return
				}
			}

			fmt.Println("------------aaaa")
		}()
		return randStream
	}

	d := make(chan any)

	randStream := newRandStream(ctx, d)
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream) // ※1
	}

	//cancel()
	d <- struct{}{}
	wg.Wait()

	fmt.Println("Done")
}
