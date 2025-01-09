package util

import "context"

type WCtx[T any] struct {
	V   T
	Ctx context.Context
}
