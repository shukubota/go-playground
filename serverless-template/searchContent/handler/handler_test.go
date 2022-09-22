package handler_test

import (
	"github.com/shukubota/go-api-template/serverless-template/searchContent/handler"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	r, err := handler.Handler()
	assert.NoError(t, err)
	assert.Equal(t, true, r)
}
