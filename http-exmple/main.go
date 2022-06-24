package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io"
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
	// exampleにあるけどこれだと動かない serverHttpを用意しないといけない
	// customHandler := func(w http.ResponseWriter, _ *http.Request) {
	// 	io.WriteString(w, "Hello from a HandleFunc #1!\n")
	// }

	customHandler, _ := newCustomHandler()

	http.Handle("/foo", customHandler)
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		body := r.Body
		defer body.Close()

		buf := new(bytes.Buffer)
		io.Copy(buf, body)

		var hello helloJSON
		json.Unmarshal(buf.Bytes(), &hello)
		fmt.Println(r.Body)
		fmt.Println(r.Method)
		fmt.Println(hello)

		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	fmt.Println("listening on :8080")
	http.ListenAndServe(":8080", nil)
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
