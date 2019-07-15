// Code generated by goa v3.0.2, DO NOT EDIT.
//
// file client HTTP transport
//
// Command:
// $ goa gen user-srv/design

package client

import (
	"context"
	"mime/multipart"
	"net/http"
	file "user-srv/gen/file"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the file service endpoint HTTP clients.
type Client struct {
	// Upload Doer is the HTTP client used to make requests to the upload endpoint.
	UploadDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// FileUploadEncoderFunc is the type to encode multipart request for the "file"
// service "upload" endpoint.
type FileUploadEncoderFunc func(*multipart.Writer, *file.UploadPayload) error

// NewClient instantiates HTTP clients for all the file service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		UploadDoer:          doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Upload returns an endpoint that makes HTTP requests to the file service
// upload server.
func (c *Client) Upload(fileUploadEncoderFn FileUploadEncoderFunc) goa.Endpoint {
	var (
		encodeRequest  = EncodeUploadRequest(NewFileUploadEncoder(fileUploadEncoderFn))
		decodeResponse = DecodeUploadResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUploadRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UploadDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("file", "upload", err)
		}
		return decodeResponse(resp)
	}
}
