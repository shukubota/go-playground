package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
)

func main() {
	log.Print("start")

	ctx := context.WithValue(context.Background(), "key", "value")

	slog.Info("start", "aaa", "bbb", "sss", struct {
		A string
		B string
	}{A: "aaa", B: "bbb"})

	slog.Error("start", "aaa", "bbb", "sss", struct {
		A string
		B string
	}{A: "aaa", B: "bbb"})

	slog.Warn("start", "aaa", "bbb", "sss", struct {
		A string
		B string
	}{A: "aaa", B: "bbb"})

	slog.ErrorContext(ctx, "start", "aaa", "bbb", "sss", struct {
		A string
		B string
	}{A: "aaa", B: "bbb"})

	a := map[string]any{
		"key": "value",
	}
	slog.Info(fmt.Sprintf("aaa %+v", a))
}
