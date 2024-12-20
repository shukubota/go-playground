package main

import (
	"fmt"
	"iter"
)

func main() {
	s := []int{1, 2, 3}

	seq := sliceToIter[int](s)
	for i, v := range seq {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}

	seq2 := sliceToIter[string]([]string{"aaa", "bbb"})

	for i, v := range seq2 {
		fmt.Printf("index: %d, value: %s\n", i, v)
	}

	// convert
	for v := range convert(seq) {
		fmt.Println(v)
	}
}

func sliceToIter[V any](s []V) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		for i, v := range s {
			if !yield(i, v) {
				break
			}
		}
	}
}

// 値が2このiteratorから値が一つのiteratorに変換
func convert[K, V any](seq iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range seq {
			if !yield(v) {
				break
			}
		}
	}
}
