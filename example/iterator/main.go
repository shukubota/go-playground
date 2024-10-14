package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	for i := range iterator {
		fmt.Println(i)
	}

}

func iterator(yield func(int) bool) {
	yield(1)
	yield(2)
	yield(3)
}
