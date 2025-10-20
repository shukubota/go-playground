package main

import (
	"fmt"
	"example/synctest/app"
)

func main() {
	if err := app.Run(); err != nil {
		fmt.Println(err)
	}
}
