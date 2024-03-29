// Code generated by goa v3.0.2, DO NOT EDIT.
//
// file HTTP server
//
// Command:
// $ goa gen user-srv/design

package server

import (
	"context"
	"mime/multipart"
	"net/http"
	file "user-srv/gen/file"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the file service endpoint HTTP handlers.
type Server struct {
	Mounts []*MountPoint
	Upload http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// FileUploadDecoderFunc is the type to decode multipart request for the "file"
// service "upload" endpoint.
type FileUploadDecoderFunc func(*multipart.Reader, **file.UploadPayload) error

// New instantiates HTTP handlers for all the file service endpoints.
func New(
	e *file.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
	fileUploadDecoderFn FileUploadDecoderFunc,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Upload", "POST", "/file/upload"},
		},
		Upload: NewUploadHandler(e.Upload, mux, NewFileUploadDecoder(mux, fileUploadDecoderFn), enc, eh),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "file" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Upload = m(s.Upload)
}

// Mount configures the mux to serve the file endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountUploadHandler(mux, h.Upload)
}

// MountUploadHandler configures the mux to serve the "file" service "upload"
// endpoint.
func MountUploadHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/file/upload", f)
}

// NewUploadHandler creates a HTTP handler which loads the HTTP request and
// calls the "file" service "upload" endpoint.
func NewUploadHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeUploadRequest(mux, dec)
		encodeResponse = EncodeUploadResponse(enc)
		encodeError    = EncodeUploadError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "upload")
		ctx = context.WithValue(ctx, goa.ServiceKey, "file")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}
