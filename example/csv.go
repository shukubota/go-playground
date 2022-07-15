package main

import (
	"fmt"
	"log"
	"os"
)

func afterRead(file *os.File) {
	file.Close()
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
	fmt.Println("test")
}

func main() {
	readCSV()
}
