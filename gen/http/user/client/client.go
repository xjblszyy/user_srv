// Code generated by goa v3.0.2, DO NOT EDIT.
//
// user client HTTP transport
//
// Command:
// $ goa gen user-srv/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the user service endpoint HTTP clients.
type Client struct {
	// Retrieve Doer is the HTTP client used to make requests to the retrieve
	// endpoint.
	RetrieveDoer goahttp.Doer

	// Create Doer is the HTTP client used to make requests to the create endpoint.
	CreateDoer goahttp.Doer

	// Signin Doer is the HTTP client used to make requests to the signin endpoint.
	SigninDoer goahttp.Doer

	// Update Doer is the HTTP client used to make requests to the update endpoint.
	UpdateDoer goahttp.Doer

	// SendEmail Doer is the HTTP client used to make requests to the sendEmail
	// endpoint.
	SendEmailDoer goahttp.Doer

	// ActiveEmail Doer is the HTTP client used to make requests to the activeEmail
	// endpoint.
	ActiveEmailDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the user service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		RetrieveDoer:        doer,
		CreateDoer:          doer,
		SigninDoer:          doer,
		UpdateDoer:          doer,
		SendEmailDoer:       doer,
		ActiveEmailDoer:     doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Retrieve returns an endpoint that makes HTTP requests to the user service
// retrieve server.
func (c *Client) Retrieve() goa.Endpoint {
	var (
		encodeRequest  = EncodeRetrieveRequest(c.encoder)
		decodeResponse = DecodeRetrieveResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRetrieveRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RetrieveDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("user", "retrieve", err)
		}
		return decodeResponse(resp)
	}
}

// Create returns an endpoint that makes HTTP requests to the user service
// create server.
func (c *Client) Create() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateRequest(c.encoder)
		decodeResponse = DecodeCreateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("user", "create", err)
		}
		return decodeResponse(resp)
	}
}

// Signin returns an endpoint that makes HTTP requests to the user service
// signin server.
func (c *Client) Signin() goa.Endpoint {
	var (
		encodeRequest  = EncodeSigninRequest(c.encoder)
		decodeResponse = DecodeSigninResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSigninRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SigninDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("user", "signin", err)
		}
		return decodeResponse(resp)
	}
}

// Update returns an endpoint that makes HTTP requests to the user service
// update server.
func (c *Client) Update() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateRequest(c.encoder)
		decodeResponse = DecodeUpdateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUpdateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("user", "update", err)
		}
		return decodeResponse(resp)
	}
}

// SendEmail returns an endpoint that makes HTTP requests to the user service
// sendEmail server.
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
			return nil, goahttp.ErrRequestError("user", "sendEmail", err)
		}
		return decodeResponse(resp)
	}
}

// ActiveEmail returns an endpoint that makes HTTP requests to the user service
// activeEmail server.
func (c *Client) ActiveEmail() goa.Endpoint {
	var (
		decodeResponse = DecodeActiveEmailResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildActiveEmailRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ActiveEmailDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("user", "activeEmail", err)
		}
		return decodeResponse(resp)
	}
}