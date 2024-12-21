package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := []string{
		"3 3",
		"...",
		"..#",
		"##.",
	}

	sc := bufio.NewScanner(strings.NewReader(strings.Join(input, "\n")))

	err := run(sc)
	if err != nil {
		log.Fatalf("err: %+v", err)
		os.Exit(1)
	}
	os.Exit(0)
}

type Grid struct {
	hgrid int // 1, 2 ... H
	wgrid int // 1,2 ... W
}

func run(sc *bufio.Scanner) error {

	// toBoard
	sc.Scan()

	header := strings.Split(sc.Text(), " ")

	h, err := strconv.Atoi(header[0])
	if err != nil {
		return err
	}
	w, err := strconv.Atoi(header[1])
	if err != nil {
		return err
	}

	board := make([][]bool, 0, h)
	for i := 0; i < h; i++ {
		row := make([]bool, 0, w)
		sc.Scan()
		r := sc.Text()
		for idx := range w {
			row = append(row, string(r[idx]) == ".")
		}
		board = append(board, row)
	}

	fmt.Println(board)
	return nil
}

// 盤面が与えられたときに(1,1)から(H, W)まで通れるかを判定する関数
func canReachGoal(盤面 any) bool {
	return true
}

// 邪魔になって取り除く候補のgridのsliceを返す関数
func getBlockers(盤面 any) []Grid {
	return []Grid{}
}
