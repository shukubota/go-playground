package main

import (
	"fmt"
	"iter"
)

func main() {
	fmt.Println("Hello, World!")

	for i := range iterator {
		fmt.Println(i)
	}

	for i := range numbers() {
		fmt.Println(i)
	}

	for name, score := range scores() {
		fmt.Println(name, score)
	}

	for prev, next := range Pair(numbers()) {
		fmt.Println(prev, next)
	}

}

func iterator(yield func(int) bool) {
	yield(1)
	yield(2)
	yield(3)
}

func numbers() iter.Seq[int] {
	return func(yield func(int) bool) {
		yield(1)
		yield(2)
		yield(3)
	}
}

func scores() iter.Seq2[string, int] {
	return func(yield func(string, int) bool) {
		yield("Alice", 100)
		yield("Bob", 90)
		yield("Charlie", 80)
	}
}

func Pair[V int](seq iter.Seq[V]) iter.Seq2[V, V] {
	return func(yield func(V, V) bool) {
		next, stop := iter.Pull(seq)
		defer stop()

		var prev V
		if v, ok := next(); ok {
			prev = v
		}

		for {
			if v, ok := next(); ok {
				if !yield(prev, v) {
					break
				}
				prev = v
			} else {
				break
			}
		}
	}
}
