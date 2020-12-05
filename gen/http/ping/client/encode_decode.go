// Code generated by goa v3.2.5, DO NOT EDIT.
//
// ping HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/nabeen/goa-realworld-example-app/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildHealthCheckRequest instantiates a HTTP request object with method and
// path set to call the "ping" service "health_check" endpoint
func (c *Client) BuildHealthCheckRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: HealthCheckPingPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("ping", "health_check", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeHealthCheckResponse returns a decoder for responses returned by the
// ping health_check endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeHealthCheckResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("ping", "health_check", resp.StatusCode, string(body))
		}
	}
}
