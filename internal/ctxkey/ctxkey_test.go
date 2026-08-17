package ctxkey

import (
	"context"
	"testing"
)

func TestSameKey(t *testing.T) {
	ctx := Set(context.Background(), "alice")
	v, ok := Get(ctx)
	if !ok || v != "alice" {
		t.Fatalf("ok=%v v=%q", ok, v)
	}
}
