package main

import (
	"fmt"
	"os"
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

	sum := sumIntsOrFloats2[string, float64](inputFloat)
	fmt.Println(sum)

	var a animal

	a = &dog{name: "dog"}

	fmt.Println(a)

	var n os.File

}

type Number interface {
	int64 | float64
}

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
