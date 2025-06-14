package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"iter"
	"os"
	"strings"
)

func main() {
	err := ReadAndWriteCSV()
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func ReadAndWriteCSV() error {
	iterator := readCSV()

	for member, err := range iterator {
		if err != nil {
			return err
		}
		fmt.Println(member)
	}

	return nil
}

type Member struct {
	ID   string
	Name string
}

func readCSV() iter.Seq2[*Member, error] {
	b := strings.NewReader("1,仁志\n2,清水\n3,高橋\n4,清原\n")
	c := csv.NewReader(b)

	return func(yield func(v *Member, err error) bool) {
		for {
			record, err := c.Read()
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				if !yield(nil, err) {
					break
				}
				break
			}
			member := &Member{
				ID:   record[0],
				Name: record[1],
			}
			if !yield(member, nil) {
				break
			}
		}
	}
}
