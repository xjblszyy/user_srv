// Code generated by goa v3.0.2, DO NOT EDIT.
//
// user endpoints
//
// Command:
// $ goa gen user-srv/design

package user

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "user" service endpoints.
type Endpoints struct {
	Retrieve    goa.Endpoint
	Create      goa.Endpoint
	Signin      goa.Endpoint
	Update      goa.Endpoint
	SendEmail   goa.Endpoint
	ActiveEmail goa.Endpoint
}

// NewEndpoints wraps the methods of the "user" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Retrieve:    NewRetrieveEndpoint(s, a.JWTAuth),
		Create:      NewCreateEndpoint(s),
		Signin:      NewSigninEndpoint(s),
		Update:      NewUpdateEndpoint(s, a.JWTAuth),
		SendEmail:   NewSendEmailEndpoint(s),
		ActiveEmail: NewActiveEmailEndpoint(s),
	}
}

// Use applies the given middleware to all the "user" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Retrieve = m(e.Retrieve)
	e.Create = m(e.Create)
	e.Signin = m(e.Signin)
	e.Update = m(e.Update)
	e.SendEmail = m(e.SendEmail)
	e.ActiveEmail = m(e.ActiveEmail)
}

// NewRetrieveEndpoint returns an endpoint function that calls the method
// "retrieve" of service "user".
func NewRetrieveEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RetrievePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.Retrieve(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedResponseData(res, "default")
		return vres, nil
	}
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "user".
func NewCreateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AddUser)
		res, err := s.Create(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedResponseData(res, "default")
		return vres, nil
	}
}

// NewSigninEndpoint returns an endpoint function that calls the method
// "signin" of service "user".
func NewSigninEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Signin)
		res, err := s.Signin(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedResponseData(res, "default")
		return vres, nil
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "user".
func NewUpdateEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdateUser)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.Update(ctx, p)
	}
}

// NewSendEmailEndpoint returns an endpoint function that calls the method
// "sendEmail" of service "user".
func NewSendEmailEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SendEmail)
		res, err := s.SendEmail(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedResponseData(res, "default")
		return vres, nil
	}
}

// NewActiveEmailEndpoint returns an endpoint function that calls the method
// "activeEmail" of service "user".
func NewActiveEmailEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*EmailCode)
		res, err := s.ActiveEmail(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedResponseData(res, "default")
		return vres, nil
	}
}
