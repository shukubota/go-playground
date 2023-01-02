package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("==============oooo")
	input := []MyInt{1, 2, 3, 4}
	r := f1(input)
	fmt.Println(r)
	fmt.Println(r[0])
	fmt.Printf("%T\n", r[0])
	fmt.Println(r[0] == "1")
}

func f1[T Stringer](xs []T) []string {
	var result []string
	for _, x := range xs {
		result = append(result, x.String())
	}
	return result
}

//func f2(xs []Stringer) []string {
//	var result []string
//	for _, x := range xs {
//		result = append(result, x.String())
//	}
//	return result
//}

type Stringer interface {
	String() string
}

type MyInt int

func (i MyInt) String() string {
	return strconv.Itoa(int(i))
}
