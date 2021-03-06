// Code generated by goa v3.2.5, DO NOT EDIT.
//
// ping client
//
// Command:
// $ goa gen github.com/nabeen/goa-realworld-example-app/design

package ping

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "ping" service client.
type Client struct {
	HealthCheckEndpoint goa.Endpoint
}

// NewClient initializes a "ping" service client given the endpoints.
func NewClient(healthCheck goa.Endpoint) *Client {
	return &Client{
		HealthCheckEndpoint: healthCheck,
	}
}

// HealthCheck calls the "health_check" endpoint of the "ping" service.
func (c *Client) HealthCheck(ctx context.Context) (err error) {
	_, err = c.HealthCheckEndpoint(ctx, nil)
	return
}
