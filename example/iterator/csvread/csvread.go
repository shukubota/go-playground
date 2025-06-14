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

	var w io.Writer
	// w = os.Stdout
	f, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	w = f

	for member, err := range iterator {
		if err != nil {
			return err
		}
		_, err = w.Write([]byte(member.ToLine()))
		if err != nil {
			return err
		}
	}

	return nil
}

type Member struct {
	ID   string
	Name string
}

func (m *Member) ToLine() string {
	return fmt.Sprintf("%s,%s\n", m.ID, m.Name)
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
