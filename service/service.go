package service

import (
	"context"

	"github.com/go-board/std/result"
	"github.com/go-board/std/slices"
)

// Service is the interface that provides the business logic for the service.
type Service[Req, Resp any] interface {
	// Call is the entry point for the service.
	Call(ctx context.Context, req Req) result.Result[Resp]
}

// ServiceFn is the function type that implements the Service interface.
type ServiceFn[Req, Resp any] func(context.Context, Req) result.Result[Resp]

func (fn ServiceFn[Req, Resp]) Call(ctx context.Context, req Req) result.Result[Resp] {
	return fn(ctx, req)
}

type Layer[S any] interface{ Then(s S) S }

type LayerFn[S any] func(s S) S

func (fn LayerFn[S]) Then(s S) S { return fn(s) }

func Chain[S any](ms ...Layer[S]) Layer[S] {
	return LayerFn[S](func(s S) S { return slices.FoldRight(ms, s, func(a S, m Layer[S]) S { return m.Then(a) }) })
}
