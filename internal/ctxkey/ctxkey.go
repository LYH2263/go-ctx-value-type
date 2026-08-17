package ctxkey

import "context"

type keyType string

const userKey keyType = "user"

func Set(ctx context.Context, v string) context.Context {
	// BUG: string key
	return context.WithValue(ctx, "user", v)
}

func Get(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(userKey).(string)
	return v, ok
}
