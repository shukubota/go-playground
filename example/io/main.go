package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	src := flag.String("src", "", "source file path")
	in := flag.String("i", "flag", "")
	out := flag.String("o", "stdout", "")
	dest := flag.String("dest", "", "destination file path")

	flag.Parse()

	fmt.Println(*src)
	fmt.Println(*in)
	fmt.Println(*out)
	fmt.Println(*dest)

	var r io.Reader
	if *in == "txt" {
		file, err := os.Open(*src)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		r = file
	} else {
		sr := bytes.NewReader([]byte(*src))
		r = sr
	}

	var w io.Writer
	if *out == "txt" {
		stdout, err := os.Create(*dest)
		if err != nil {
			fmt.Println(err)
			return
		}
		w = stdout
	} else {
		f := os.Stdout
		defer f.Close()
		w = f
	}

	err := toUpper(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func toUpper(w io.Writer, r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	_, err = w.Write(bytes.ToUpper(data))
	if err != nil {
		return err
	}
	return nil
}
