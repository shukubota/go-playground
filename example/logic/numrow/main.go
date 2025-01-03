package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var reader io.Reader

	// For test
	// input := []string{
	// 	"4",
	// 	"2 1",
	// 	"2 5",
	// 	"3 1",
	// 	"1 2",
	// }
	// reader = strings.NewReader(strings.Join(input, "\n"))

	file, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	reader = file

	sc := bufio.NewScanner(reader)

	err = run(sc)
	if err != nil {
		os.Exit(1)
	}
}

func run(sc *bufio.Scanner) error {
	sc.Scan()
	rowCount, err := strconv.Atoi(sc.Text())
	if err != nil {
		return err
	}

	all := make([]int, 0)

	for i := 0; i < rowCount; i++ {
		sc.Scan()
		r := strings.Split(sc.Text(), " ")

		c, err := strconv.Atoi(r[0])
		if err != nil {
			return err
		}
		value, err := strconv.Atoi(r[1])
		if err != nil {
			return err
		}
		for j := 0; j < c; j++ {
			all = append(all, value)
		}
	}

	count := len(all) / 2

	result := 0

	for i := 0; i < count; i++ {
		v := all[i] - all[i+count]
		if v < 0 {
			v = -v
		}
		result = result + v
	}

	fmt.Println(result)

	return nil
}
