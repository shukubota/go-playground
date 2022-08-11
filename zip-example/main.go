package main

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	fmt.Println("main")
	b := new(bytes.Buffer)
	zw := zip.NewWriter(b)

	raw := "iVBORw0KGgoAAAANSUhEUgAAAN4AAADeAQMAAABoqSz0AAAABlBMVEX///8AAABVwtN+AAAACXBIWXMAAA7EAAAOxAGVKw4bAAABTklEQVRYhe2Y6w3DIAyELTEAI7E6I2UAS67PBy1VFrhIsaqW+uNHTM8Pavbaa8+0gM1ceJ/ewzq+uyrEJ9yzpfPCejk1YQ9fjsF1RiYO6/ijXfYI2K6RJLmJw5JJCiQPnnq5aUgIMgUr//i656cQpLWsGX0y/24mBPHg+fjXQAj5Dr0MWehGgQQ35sL/NKQFed5LIFmM7RuHImwoD4PRMIjj4PXgBd4qEVHh0OtCF5avuty0CugoxnIwqrl5KRo/wi8ORUgVu60tXtt1YV9FoqG2JRmHTORgiiIdNfbOlYj96IJiMB2Nk05ugcYxppks7MGa8ctC04YobzVRxpqCXRXSWC149sxFTchxcjfk4Xv+1YQc1Y2Xtj0BuSzc18rYN6Gje4hCJiKKR0nbtWHU1MO5EiOwLLT114hzY0lbFu4UHLtpDLvnpwx87bWn2QdEUZa6xVXnqQAAAABJRU5ErkJggg=="
	body, _ := base64.StdEncoding.DecodeString(raw)

	zf, err := os.CreateTemp(".", "*.zip")
	defer zf.Close()
	defer os.Remove(zf.Name())
	fileInfo, err := zf.Stat()
	fileHeader, _ := zip.FileInfoHeader(fileInfo)
	fileHeader.Name = "output.png"
	writer, err := zw.CreateHeader(fileHeader)
	writer.Write(body)
	zw.Close()
	if err != nil {
		return
	}

	zf.Write(b.Bytes())
	fmt.Println(b.Bytes())
}
