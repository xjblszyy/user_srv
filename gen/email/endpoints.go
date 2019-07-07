// Code generated by goa v3.0.2, DO NOT EDIT.
//
// email endpoints
//
// Command:
// $ goa gen user/design

package email

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "email" service endpoints.
type Endpoints struct {
	Active    goa.Endpoint
	SendEmail goa.Endpoint
}

// NewEndpoints wraps the methods of the "email" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Active:    NewActiveEndpoint(s),
		SendEmail: NewSendEmailEndpoint(s),
	}
}

// Use applies the given middleware to all the "email" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Active = m(e.Active)
	e.SendEmail = m(e.SendEmail)
}

// NewActiveEndpoint returns an endpoint function that calls the method
// "active" of service "email".
func NewActiveEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ActivePayload)
		return s.Active(ctx, p)
	}
}

// NewSendEmailEndpoint returns an endpoint function that calls the method
// "send_email" of service "email".
func NewSendEmailEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SendEmailPayload)
		return s.SendEmail(ctx, p)
	}
}