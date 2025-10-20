package main_test

import (
	"context"
	"example/synctest/app"
	"testing"
	"testing/synctest"
	"time"
)

func TestContext_withTimout(t *testing.T) {
	t.Parallel()

	synctest.Test(t, func(t *testing.T) {
		t.Log(time.Now())
		const timeout = 24 * time.Hour
		ctx, cancel := context.WithTimeout(t.Context(), timeout)
		defer cancel()

		time.Sleep(timeout / 4)

		t.Log(time.Now())

		if err := ctx.Err(); err != nil {
			t.Fatalf("unexpected error before timeout: %v", err)
		}

		time.Sleep(timeout)

		if err := ctx.Err(); err == nil {
			t.Fatalf("expected error after timeout, got nil")
		}

		err := app.Run()
		if err != nil {
			t.Fatalf("app.Run() failed: %v", err)
		}

	})

}
