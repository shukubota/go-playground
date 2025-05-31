package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")

	var r io.ReadSeeker
	file, err := os.Open("./read.csv")
	if err != nil {
		log.Fatal(err)
	}

	r = file

	tmpFile, err := os.CreateTemp("", "push")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := tmpFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	name := tmpFile.Name()

	fmt.Println(name)

	_, err = io.Copy(tmpFile, r)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	currentPos, err := f.Seek(0, io.SeekCurrent)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(currentPos)

	cr := csv.NewReader(f)
	header, err := cr.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header)

	//bs, err := io.ReadAll(cr)
	//if err != nil {
	//	log.Fatal(err)
	//}

	for {
		rs, err := cr.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(rs)
	}

	// csv読むパターン
	buf := &bytes.Buffer{}
	file, err = os.Open("./read.csv")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(buf, file)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())

	// strings読むパターン
	buf2 := &bytes.Buffer{}
	st := strings.NewReader("member_id,params\n111,aaa\n")

	_, err = io.Copy(buf2, st)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf2.String())

	fmt.Println(buf.String() == buf2.String())
}
