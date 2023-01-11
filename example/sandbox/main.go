package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println(time.Now())

	random := rand.Intn(3)
	fmt.Println(random)

	fmt.Println(split(9))

	fmt.Println("==============func")
	input := []MyInt{1, 2, 3, 4}
	r := f1(input)
	fmt.Println(r)
	fmt.Println(r[0])
	fmt.Printf("%T\n", r[0])
	fmt.Println(r[0] == "1")

	fmt.Println("==============stack")
	s := New[string]()
	s.Push("hello")
	fmt.Println(s)
}

func split(sum int) (x, y int) {
	xx := sum * 4 / 9
	yy := sum - xx
	fmt.Println(x)
	return xx, yy
}

type Stack[T any] []T

func New[T any]() *Stack[T] {
	v := make(Stack[T], 0)
	return &v
}

func (s *Stack[T]) Push(x T) {
	*s = append(*s, x)
}

func (s *Stack[T]) Pop() T {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func f1[A Stringer](xs []A) []string {
	var result []string
	for _, x := range xs {
		fmt.Printf("%T\n", x)
		result = append(result, x.String())
	}
	return result
}

// これだとコンパイルできない
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
