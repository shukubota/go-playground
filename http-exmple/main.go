package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
)

type CustomHandler struct {
}

func (h *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w)
	fmt.Println(r)
	io.WriteString(w, "CustomHandler\n")
}

func newCustomHandler() (*CustomHandler, error) {
	h := &CustomHandler{}
	return h, nil
}

type helloJSON struct {
	UserName string `json:"user_name"`
	Content  string `json:"content"`
}

func main() {
	// Handleでcustom handlerを登録する場合
	customHandler, _ := newCustomHandler()
	http.Handle("/foo", customHandler)

	// handler関数を直接指定する場合
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		body := r.Body
		defer body.Close()

		buf := &bytes.Buffer{}
		io.Copy(buf, body)

		var hello helloJSON
		json.Unmarshal(buf.Bytes(), &hello)
		fmt.Println(r.Body)
		fmt.Println(r.Method)
		fmt.Println(hello)

		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	fmt.Println("listening on :18080")
	log.Fatal(http.ListenAndServe("127.0.0.1:18080", nil))
}
