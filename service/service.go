package service

import (
	"context"

	"github.com/go-board/std/result"
)

type Service[Req, Resp any] interface {
	Call(ctx context.Context, req Req) result.Result[Resp]
}

type ServiceFn[Req, Resp any] func(context.Context, Req) result.Result[Resp]

func (self ServiceFn[Req, Resp]) Call(ctx context.Context, req Req) result.Result[Resp] {
	return self(ctx, req)
}

type Layer[Req, Resp any, S Service[Req, Resp]] interface {
	Next(service S) S
}

type LayerFn[Req, Resp any, S Service[Req, Resp]] func(service S) S

func (self LayerFn[Req, Resp, S]) Next(service S) S { return self(service) }

func ComposeLayers[Req, Resp any, S Service[Req, Resp], L Layer[Req, Resp, S]](service S, layers ...L) S {
	for _, layer := range layers {
		service = layer.Next(service)
	}
	return service
}
