package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {

	b := &bytes.Buffer{}
	w := csv.NewWriter(b)

	err := w.Write([]string{"a", "b", "c"})
	if err != nil {
		return err
	}

	w.Flush()
	w.Flush()

	fmt.Print(b.String())
	fmt.Println(w.Error())
	return nil
}
