package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

func main() {
	inputInt := map[string]int64{
		"idx1": 1,
		"idx2": 2,
	}

	inputFloat := map[string]float64{
		"idx1": 10.1,
		"idx2": 10.22,
	}

	sumInt := sumInts(inputInt)
	fmt.Println(sumInt)

	sumFloats := sumFloats(inputFloat)
	fmt.Println(sumFloats)

	//input := map[string]string{
	//	"idx1": "1",
	//	"idx2": "2",
	//}

	//s := sumIntsOrFloats2[string, float64](inputFloat)
	//fmt.Println(sum)

	t := time.Date(2023, 2, 2, 10, 10, 11, 0, time.UTC)

	fmt.Println(t)

	a := atomic.Pointer[string]{}

	fmt.Println(a)

	var v atomic.Int64
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			v.Add(1)
		}()
	}
	wg.Wait()
	fmt.Println(v.Load())

	h := Hoge[string]{Name: "aaa"}
	hp := pointer[Hoge[string]]{}

	hp2 := pointer[Hoge2]{}
	hp = pointer[Hoge[string]](hp2)

	fmt.Println(h)
	fmt.Println(hp)
	fmt.Printf("%T\n", hp)

	s := sum[int](1, 2, 3)
	fmt.Println(s)

	type customNumber int

	vc1 := customNumber(1)
	vc2 := customNumber(2)
	fmt.Println(vc1, vc2)

	//fmt.Println(sum2[customNumber](vc1, vc2)) //コンパイルエラー

	// ifのテスト
	hoge := NewHoge[string]("name1")

	fmt.Println(hoge)

	// mapのテスト
	m := map[string]int{
		"idx1": 1,
		"idx2": 29,
	}
	fmt.Println(sumFromMap(m))

	st := "some string"
	i := 12
	h2 := Hoge2{Name: "a1"}
	fmt.Printf("%T\n", toPointer(st))
	fmt.Printf("%T\n", toPointer(i))
	fmt.Printf("%T\n", toPointer(h2))
}

type target interface {
	string | int | Hoge2
}

func toPointer[T target](v T) *T {
	return &v
}

type Hoge[T any] struct {
	Name string
}
type Hoge2 struct {
	Name string
}

func (h *Hoge[T]) greet(s T) string {
	fmt.Println(s)
	return "Hello"
}

type hogeIF[T any] interface {
	greet(params T) string
}

func NewHoge[T any](name string) hogeIF[T] {
	return &Hoge[T]{
		Name: name,
	}
}

type pointer[T any] struct {
	//_ [0]*T
	v unsafe.Pointer
}

func sum[T constraints.Integer](vs ...T) T {
	var sum T
	for _, v := range vs {
		sum += v
	}
	return sum
}

func sum2[T Number](vs ...T) T {
	var sum T
	for _, v := range vs {
		sum += v
	}
	return sum
}

type mapKey interface {
	int | string | int64
}

func sumFromMap[T mapKey, U constraints.Integer](m map[T]U) U {
	var s U
	for _, v := range m {
		s += v
	}
	return s
}

type Number interface {
	int64 | float64
}

type A comparable

func sumIntsOrFloats2[K string, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func sumIntsOrFloats[K string, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func sumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func sumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}
