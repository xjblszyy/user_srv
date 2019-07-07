// Code generated by goa v3.0.2, DO NOT EDIT.
//
// email HTTP server
//
// Command:
// $ goa gen user/design

package server

import (
	"context"
	"net/http"
	email "user/gen/email"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the email service endpoint HTTP handlers.
type Server struct {
	Mounts    []*MountPoint
	Active    http.Handler
	SendEmail http.Handler
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

// New instantiates HTTP handlers for all the email service endpoints.
func New(
	e *email.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Active", "GET", "/email/active/{code}"},
			{"SendEmail", "POST", "/email/active"},
		},
		Active:    NewActiveHandler(e.Active, mux, dec, enc, eh),
		SendEmail: NewSendEmailHandler(e.SendEmail, mux, dec, enc, eh),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "email" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Active = m(s.Active)
	s.SendEmail = m(s.SendEmail)
}

// Mount configures the mux to serve the email endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountActiveHandler(mux, h.Active)
	MountSendEmailHandler(mux, h.SendEmail)
}

// MountActiveHandler configures the mux to serve the "email" service "active"
// endpoint.
func MountActiveHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/email/active/{code}", f)
}

// NewActiveHandler creates a HTTP handler which loads the HTTP request and
// calls the "email" service "active" endpoint.
func NewActiveHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeActiveRequest(mux, dec)
		encodeResponse = EncodeActiveResponse(enc)
		encodeError    = EncodeActiveError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "active")
		ctx = context.WithValue(ctx, goa.ServiceKey, "email")
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

// MountSendEmailHandler configures the mux to serve the "email" service
// "send_email" endpoint.
func MountSendEmailHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/email/active", f)
}

// NewSendEmailHandler creates a HTTP handler which loads the HTTP request and
// calls the "email" service "send_email" endpoint.
func NewSendEmailHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeSendEmailRequest(mux, dec)
		encodeResponse = EncodeSendEmailResponse(enc)
		encodeError    = EncodeSendEmailError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "send_email")
		ctx = context.WithValue(ctx, goa.ServiceKey, "email")
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
