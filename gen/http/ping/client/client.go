// Code generated by goa v3.2.5, DO NOT EDIT.
//
// ping client HTTP transport
//
// Command:
// $ goa gen github.com/nabeen/goa-realworld-example-app/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the ping service endpoint HTTP clients.
type Client struct {
	// HealthCheck Doer is the HTTP client used to make requests to the
	// health_check endpoint.
	HealthCheckDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the ping service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		HealthCheckDoer:     doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// HealthCheck returns an endpoint that makes HTTP requests to the ping service
// health_check server.
func (c *Client) HealthCheck() goa.Endpoint {
	var (
		decodeResponse = DecodeHealthCheckResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildHealthCheckRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.HealthCheckDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("ping", "health_check", err)
		}
		return decodeResponse(resp)
	}
}