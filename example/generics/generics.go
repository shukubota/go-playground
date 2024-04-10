package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/exp/constraints"
	"sync/atomic"
)

func main() {
	fmt.Println(someFunc[Price](1))
	fmt.Println(someFunc[SKU]("sku1"))

	var v atomic.Int64
	fmt.Println(v)

	//j := "'{\"name\": \"taro\"}"
	u := user{Name: "taro"}

	r, err := json.Marshal(u)
	if err != nil {
		fmt.Errorf("error: %v", err)
		fmt.Println(err)
		return
	}

	fmt.Println(r)

	du := user{}

	err = json.Unmarshal(r, du)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(du)
}

type user struct {
	Name string `json:"name"`
}

type SKU string
type Price int

func someFunc[T ~string | ~int | constraints.Integer](x T) T {
	return x
}
