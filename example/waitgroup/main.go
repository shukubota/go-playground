package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
	"time"
)

func main() {
	sem := semaphore.NewWeighted(3)
	ctx := context.TODO()
	//wg := sync.WaitGroup{}
	eg, ctx := errgroup.WithContext(ctx)
	//for i := 0; i < 20; i++ {
	//	//eg.Add(1)
	//	i := i
	//	eg.Go(func() error {
	//		err := waitFunc(ctx, i)
	//		if err != nil {
	//			return err
	//		}
	//		return nil
	//	})
	//
	//}
	////wg.Wait()
	//if err := eg.Wait(); err != nil {
	//	fmt.Printf("err: %+v\n", err)
	//}

	for i := 0; i < 20; i++ {
		i := i
		if err := sem.Acquire(ctx, 1); err != nil {
			fmt.Printf("Failed to acquire semaphore: %v", err)
			break
		}
		eg.Go(func() error {
			defer sem.Release(1)
			err := waitFunc(ctx, i)
			if err != nil {
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		fmt.Printf("err: %+v\n", err)
	}
	fmt.Println("------finished")
}

func waitFunc(ctx context.Context, i int) error {
	//defer eg.Done()

	fmt.Printf("before sleep i: %d\n", i)
	// iが90以上ならエラーを返す
	if i >= 19 {
		return fmt.Errorf("error: i is %d\n", i)
	}

	time.Sleep(time.Millisecond * 100)
	fmt.Println("after sleep")
	return nil
}
