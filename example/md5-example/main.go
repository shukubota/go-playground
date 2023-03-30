package main

import (
	"crypto/md5"
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	activationCode := "ac12345"
	udid := "udid123"

	data := []byte(activationCode + udid)
	fmt.Println(data)

	hash := md5.Sum(data)
	fmt.Println(hash)
	fmt.Println(fmt.Sprintf("%x", hash))
}
