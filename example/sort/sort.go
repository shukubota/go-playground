package main

import (
	"fmt"
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
}

func main() {
	fmt.Println(data)
	sort.Slice(data, func(i, j int) bool {
		return data[i].createdAt.After(*data[j].createdAt)
	})

	fmt.Println(data)
}
