// Code generated by goa v3.0.2, DO NOT EDIT.
//
// email client HTTP transport
//
// Command:
// $ goa gen user/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the email service endpoint HTTP clients.
type Client struct {
	// Active Doer is the HTTP client used to make requests to the active endpoint.
	ActiveDoer goahttp.Doer

	// SendEmail Doer is the HTTP client used to make requests to the send_email
	// endpoint.
	SendEmailDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the email service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		ActiveDoer:          doer,
		SendEmailDoer:       doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Active returns an endpoint that makes HTTP requests to the email service
// active server.
func (c *Client) Active() goa.Endpoint {
	var (
		decodeResponse = DecodeActiveResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildActiveRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ActiveDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("email", "active", err)
		}
		return decodeResponse(resp)
	}
}

// SendEmail returns an endpoint that makes HTTP requests to the email service
// send_email server.
func (c *Client) SendEmail() goa.Endpoint {
	var (
		encodeRequest  = EncodeSendEmailRequest(c.encoder)
		decodeResponse = DecodeSendEmailResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSendEmailRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SendEmailDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("email", "send_email", err)
		}
		return decodeResponse(resp)
	}
}