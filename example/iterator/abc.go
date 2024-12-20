package main

import (
	"fmt"
)

func main() {
	for c := range Alphabet {
		fmt.Printf("%c\n", c)
		if c == 'B' {
			break
		}
	}
}

func Alphabet(yield func(rune) bool) {
	l := []rune{'A', 'B', 'C'}
	for _, c := range l {
		if !yield(c) {
			break
		}
	}
}
