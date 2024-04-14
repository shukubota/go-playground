package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
)

type data struct {
	Test string `json:"test"`
}

var secret = "testtest"

// secret testtestと {"test":"test"} をHMAC-SHA256でハッシュ化して比較する
// hashは9ca66abefa1254df726385f07432ed60a648b853a0d71ea1c375865252dabf9dになる

func main() {
	d := data{
		Test: "test",
	}

	body, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}

	expected := "9ca66abefa1254df726385f07432ed60a648b853a0d71ea1c375865252dabf9d"
	verified, err := verify(body, []byte(expected))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(verified)
}

// see https://pkg.go.dev/crypto/hmac
func generate(body []byte) ([]byte, error) {
	hm := hmac.New(sha256.New, []byte(secret))

	hm.Write(body)

	expectedMAC := hm.Sum(nil)
	hash := make([]byte, hex.EncodedLen(len(expectedMAC)))

	hex.Encode(hash, expectedMAC)

	return hash, nil
}

func verify(body []byte, hash []byte) (bool, error) {
	fromBody, err := generate(body)
	if err != nil {
		return false, err
	}

	return hmac.Equal(fromBody, hash), nil
}
