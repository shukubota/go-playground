package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

type member struct {
	id        int
	name      string
	createdAt *time.Time
}

func toPointer(t time.Time) *time.Time {
	return &t
}

var data = []member{
	{
		id:        1,
		name:      "one",
		createdAt: toPointer(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
	},
	{
		id:        2,
		name:      "two",
		createdAt: toPointer(time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)),
	},
	{
		id:        3,
		name:      "three",
		createdAt: toPointer(time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC)),
	},
	{
		id:        33,
		name:      "three33",
		createdAt: toPointer(time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC)),
	},
	{
		id:        44,
		name:      "four44",
		createdAt: toPointer(time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC)),
	},
	{
		id:        4,
		name:      "four",
		createdAt: toPointer(time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC)),
	},
	{
		id:        444,
		name:      "four444",
		createdAt: toPointer(time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC)),
	},
	{
		id:        4444,
		name:      "four444-",
		createdAt: toPointer(time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC)),
	},
}

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println(data)
	sort.Slice(data, func(i, j int) bool {
		return data[i].createdAt.After(*data[j].createdAt)
	})

	fmt.Println(data)
	people := []Person{
		{Name: "V", Age: 3},
		{Name: "K", Age: 3},
		{Name: "Y", Age: 3},
		{Name: "A", Age: 4},
		{Name: "E", Age: 3},
		{Name: "D", Age: 1},
		{Name: "C", Age: 3},
		{Name: "X", Age: 2},
		{Name: "B", Age: 3},
	}

	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Printf("NameでSort(Not-Stable):%+v\n", people)

	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Printf("AgeでSort(Not-Stable):%+v\n", people)

	a := math.Sin(math.Pi / 3.0)
	fmt.Println(a)
	fmt.Println(math.Sqrt(3) * 0.5)
}
