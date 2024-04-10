package main

import "fmt"

func main() {
	h := hoge{name: "hoge"}
	hp := &h

	fmt.Println(h)
	fmt.Println(hp)

	doSomething(h, hp, []hoge{h}, []*hoge{hp})

	fmt.Println(h)
	fmt.Println(hp)
}

type hoge struct {
	name string
}

func doSomething(h hoge, hp *hoge, hl []hoge, hpl []*hoge) {
	fmt.Println(h)
	fmt.Println(hp)
	fmt.Println(hl)
	fmt.Println(hpl)

	//a := hl[0]
	a := hpl[0]
	a.name = "new name"
}
