package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("=============aaaa")
	a := flag.Bool("b", false, "aaaa")
	b := flag.Int("int", 0, "aaaa")
	flag.Parse()

	fmt.Println(*a)
	fmt.Println(*b)
}
