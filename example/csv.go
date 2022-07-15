package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func afterRead(file *os.File) {
	file.Close()
}

var UTF8_BOM = []byte{239, 187, 191}

func hasBOM(in []byte) bool {
	return bytes.HasPrefix(in, UTF8_BOM)
}

func stripBOM(in []byte) []byte {
	return bytes.TrimPrefix(in, UTF8_BOM)
}

func readCSV() {
	filename := "/Users/shu.kubota/Desktop/pushtest.csv"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		log.Panic("failed to read file")
	}
	defer afterRead(file)

	fmt.Println(file)

	// nil, 0 (not bom)
	fmt.Println("00sssss")
	var row []byte
	row, err = io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		log.Panic("failed to read file")
	}

	hasBom := hasBOM(row)
	fmt.Printf("hasbom: %v\n", hasBom)

	var rs string
	if hasBom {
		rs = string(stripBOM(row))
	} else {
		rs = string(row)
	}

	fmt.Println(row)
	fmt.Println(rs)

	a := strings.Split(rs, "\r\n")
	fmt.Println(a)
	fmt.Println(len(a))
	for _, memberId := range a {
		fmt.Println(memberId)
		fmt.Println(memberId == "Tansu1111member")
	}

}

func main() {
	readCSV()
}
