package raw

import "context"

// RawService is the interface that provides the business logic for the service.
type RawService[Req, Resp any] interface {
	// RawCall is the entry point for the service.
	RawCall(ctx context.Context, req Req) (Resp, error)
}

// RawServiceFn is the function type that implements the Service interface.
type RawServiceFn[Req, Resp any] func(context.Context, Req) (Resp, error)

func (self RawServiceFn[Req, Resp]) RawCall(ctx context.Context, req Req) (Resp, error) {
	return self(ctx, req)
}

type RawLayer[Req, Resp any, S RawService[Req, Resp]] interface {
	RawNext(service S) S
}

type RawLayerFn[Req, Resp any, S RawService[Req, Resp]] func(service S) S

func (self RawLayerFn[Req, Resp, S]) RawNext(service S) S { return self(service) }

func ComposeRawLayers[Req, Resp any, S RawService[Req, Resp], L RawLayer[Req, Resp, S]](service S, layers ...L) S {
	for _, layer := range layers {
		service = layer.RawNext(service)
	}
	return service
}
