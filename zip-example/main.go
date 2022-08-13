package main

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"example/hello/zip-example/adaptor"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("main")
	b := new(bytes.Buffer)

	raw := "iVBORw0KGgoAAAANSUhEUgAAAN4AAADeAQMAAABoqSz0AAAABlBMVEX///8AAABVwtN+AAAACXBIWXMAAA7EAAAOxAGVKw4bAAABTklEQVRYhe2Y6w3DIAyELTEAI7E6I2UAS67PBy1VFrhIsaqW+uNHTM8Pavbaa8+0gM1ceJ/ewzq+uyrEJ9yzpfPCejk1YQ9fjsF1RiYO6/ijXfYI2K6RJLmJw5JJCiQPnnq5aUgIMgUr//i656cQpLWsGX0y/24mBPHg+fjXQAj5Dr0MWehGgQQ35sL/NKQFed5LIFmM7RuHImwoD4PRMIjj4PXgBd4qEVHh0OtCF5avuty0CugoxnIwqrl5KRo/wi8ORUgVu60tXtt1YV9FoqG2JRmHTORgiiIdNfbOlYj96IJiMB2Nk05ugcYxppks7MGa8ctC04YobzVRxpqCXRXSWC149sxFTchxcjfk4Xv+1YQc1Y2Xtj0BuSzc18rYN6Gje4hCJiKKR0nbtWHU1MO5EiOwLLT114hzY0lbFu4UHLtpDLvnpwx87bWn2QdEUZa6xVXnqQAAAABJRU5ErkJggg=="
	body, _ := base64.StdEncoding.DecodeString(raw)

	fl := make([]string, 2, 2)
	fl[0] = "test1.png"
	fl[1] = "test2.png"

	var bs []byte

	//zf, err := os.Create("fuga.zip")
	zw := zip.NewWriter(b)
	for _, fileName := range fl {

		//// これでファイルごと作成
		//f, err := zw.Create(fileName)
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}
		//_, err = f.Write(body)
		//
		//fmt.Println(fileName)
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}

		zf, err := os.Create(fileName)
		fileInfo, err := zf.Stat()
		if err != nil {
			fmt.Println(err)
			return
		}
		fileHeader, _ := zip.FileInfoHeader(fileInfo)
		fileHeader.Name = fileName
		writer, err := zw.CreateHeader(fileHeader)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = writer.Write(body)
		if err != nil {
			fmt.Println(err)
			return
		}
		//bs = append(bs, b.Bytes()...)
		zf.Write(b.Bytes())
		zf.Close()
	}

	zw.Close()

	fmt.Println(bs)
	fmt.Println("================bs")
	//zf.Write(bs)

	//defer os.Remove(zf.Name())

	fmt.Println("bytes-----------")
	fmt.Println(b.Bytes())

	s, err := adaptor.NewAdapter()
	if err != nil {
		fmt.Println(err)
	}

	content := strings.NewReader(string(b.Bytes()))

	fmt.Println(content)
	fmt.Println(s)
	//zf.Close()
	// １回閉じてから開く
	//file, err := os.Open(zf.Name())
	s.Upload(content)
}
