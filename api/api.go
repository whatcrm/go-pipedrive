package api

import (
	"context"
	
	"github.com/valyala/fasthttp"
)

type api struct {
	url string
	request *fasthttp.Request
	accessToken string
}

type API interface {
	Get(ctx context.Context, endpoint string) (any, error)
	Post(ctx context.Context, endpoint string, data any) (any, error)
	Delete(ctx context.Context, endpoint string) (any, error)
}

func NewAPI(url string, request *fasthttp.Request, accessToken string) API {
	return &api{
		url: url,
		request: request,
		accessToken: accessToken,
	}
}

func (a *api) Delete(ctx context.Context, endpoint string) (any, error) {
	panic("unimplemented")
}

func (a *api) Get(ctx context.Context, endpoint string) (any, error) {
	panic("unimplemented")
}

func (a *api) Post(ctx context.Context, endpoint string, data any) (any, error) {
	panic("unimplemented")
}
