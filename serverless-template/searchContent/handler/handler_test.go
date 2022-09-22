package handler_test

import (
	"example/hello/serverless-template/searchContent/handler"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	r, err := handler.Handler()
	assert.NoError(t, err)
	assert.Equal(t, true, r)
}
