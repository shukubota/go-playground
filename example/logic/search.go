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
		"2 2",
		".#",
		"#.",
		// "...",
		// "..#",
		// "##.",
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

func (g *Grid) GetNextGrids(history []Grid, board [][]bool) ([]Grid, bool, []Grid) {
	candidates := []Grid{
		{
			hgrid: g.hgrid,
			wgrid: g.wgrid + 1,
		},
		{
			hgrid: g.hgrid,
			wgrid: g.wgrid - 1,
		},
		{
			hgrid: g.hgrid + 1,
			wgrid: g.wgrid,
		},
		{
			hgrid: g.hgrid - 1,
			wgrid: g.wgrid,
		},
	}

	next := make([]Grid, 0)
	blockers := make([]Grid, 0)

	for _, c := range candidates {
		exists := false
		for _, v := range history {
			if v.hgrid == c.hgrid && v.wgrid == c.wgrid {
				exists = true
			}
		}

		if exists {
			continue
		}

		if c.hgrid > len(board) || c.hgrid < 1 || c.wgrid > len(board[0]) || c.wgrid < 1 {
			continue
		}

		if board[c.hgrid-1][c.wgrid-1] {
			next = append(next, c)
		} else {
			blockers = append(blockers, c)
		}
	}
	return next, len(next) == 0, blockers
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

	count := 0
	current := Grid{
		hgrid: 1,
		wgrid: 1,
	}
	history := []Grid{
		current,
	}

	// 1回通れるかチェック
	canReach, blockers := canReachGoal(board, history, current)

	if canReach {
		return nil
	}

	// blockersの重複を排除する
	flattedBlockers := make([]Grid, 0)
	for _, v := range blockers {
		exists := false
		for _, f := range flattedBlockers {
			if f.hgrid == v.hgrid && f.wgrid == v.wgrid {
				exists = true
				break
			}
		}

		if !exists {
			flattedBlockers = append(flattedBlockers, v)
		}
	}

	for _, v := range flattedBlockers {
		// 盤面を変える
		newBoard := make([][]bool, len(board), len(board))

		for i := 0; i < len(board); i++ {
			row := make([]bool, len(board[0]), len(board[0]))
			for j := 0; j < len(board[0]); j++ {
				if i == v.hgrid-1 && j == v.wgrid-1 {
					row[j] = true
					continue
				}
				row[j] = board[i][j]
			}
			newBoard[i] = row
		}

		current := Grid{
			hgrid: 1,
			wgrid: 1,
		}
		history := []Grid{
			current,
		}

		canReach, _ := canReachGoal(newBoard, history, current)

		if canReach {
			count++
		}
	}

	fmt.Println(count)

	return nil
}

// 盤面が与えられたときに(1,1)から(H, W)まで通れるかを判定する関数
func canReachGoal(board [][]bool, history []Grid, current Grid) (bool, []Grid) {
	allBlockers := make([]Grid, 0)

	next, isTerminal, blockers := current.GetNextGrids(history, board)

	// 終端なら到達できない
	if isTerminal {
		return false, blockers
	}

	for _, v := range next {
		current = v
		history = append(history, v)

		if current.hgrid == len(board) && current.wgrid == len(board[0]) {
			return true, allBlockers
		}

		canReach, bs := canReachGoal(board, history, current)
		allBlockers = append(allBlockers, bs...)

		if canReach {
			return true, bs
		}
	}

	return false, allBlockers
}
