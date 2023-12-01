package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_main(t *testing.T) {
	t.Run("readWrite", func(t *testing.T) {
		r := bytes.NewBuffer([]byte("test"))
		w := new(bytes.Buffer)
		//
		err := toUpper(w, r)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(w.Bytes())
		assert.Equal(t, w.String(), "TEST")
	})
}
