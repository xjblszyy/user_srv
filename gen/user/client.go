// Code generated by goa v3.0.2, DO NOT EDIT.
//
// user client
//
// Command:
// $ goa gen user-srv/design

package user

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "user" service client.
type Client struct {
	RetrieveEndpoint    goa.Endpoint
	CreateEndpoint      goa.Endpoint
	SigninEndpoint      goa.Endpoint
	UpdateEndpoint      goa.Endpoint
	SendEmailEndpoint   goa.Endpoint
	ActiveEmailEndpoint goa.Endpoint
}

// NewClient initializes a "user" service client given the endpoints.
func NewClient(retrieve, create, signin, update, sendEmail, activeEmail goa.Endpoint) *Client {
	return &Client{
		RetrieveEndpoint:    retrieve,
		CreateEndpoint:      create,
		SigninEndpoint:      signin,
		UpdateEndpoint:      update,
		SendEmailEndpoint:   sendEmail,
		ActiveEmailEndpoint: activeEmail,
	}
}

// Retrieve calls the "retrieve" endpoint of the "user" service.
// Retrieve may return the following errors:
//	- "not_found" (type *goa.ServiceError): User not found
//	- "invalide_token" (type *goa.ServiceError): invalide token
//	- error: internal error
func (c *Client) Retrieve(ctx context.Context, p *RetrievePayload) (res *ResponseData, err error) {
	var ires interface{}
	ires, err = c.RetrieveEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ResponseData), nil
}

// Create calls the "create" endpoint of the "user" service.
// Create may return the following errors:
//	- "email_exist" (type *goa.ServiceError): email exist
//	- error: internal error
func (c *Client) Create(ctx context.Context, p *AddUser) (res *ResponseData, err error) {
	var ires interface{}
	ires, err = c.CreateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ResponseData), nil
}

// Signin calls the "signin" endpoint of the "user" service.
// Signin may return the following errors:
//	- "email_or_password_error" (type *goa.ServiceError): email or password error
//	- error: internal error
func (c *Client) Signin(ctx context.Context, p *Signin) (res *ResponseData, err error) {
	var ires interface{}
	ires, err = c.SigninEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ResponseData), nil
}

// Update calls the "update" endpoint of the "user" service.
// Update may return the following errors:
//	- "not_found" (type *goa.ServiceError): user not found
//	- "invalide_token" (type *goa.ServiceError): invalide token
//	- error: internal error
func (c *Client) Update(ctx context.Context, p *UpdateUser) (err error) {
	_, err = c.UpdateEndpoint(ctx, p)
	return
}

// SendEmail calls the "sendEmail" endpoint of the "user" service.
// SendEmail may return the following errors:
//	- "email_not_found" (type *goa.ServiceError): email not found
//	- error: internal error
func (c *Client) SendEmail(ctx context.Context, p *SendEmail) (res *ResponseData, err error) {
	var ires interface{}
	ires, err = c.SendEmailEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ResponseData), nil
}

// ActiveEmail calls the "activeEmail" endpoint of the "user" service.
// ActiveEmail may return the following errors:
//	- "code_invalide" (type *goa.ServiceError): code_invalide
//	- error: internal error
func (c *Client) ActiveEmail(ctx context.Context, p *EmailCode) (res *ResponseData, err error) {
	var ires interface{}
	ires, err = c.ActiveEmailEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ResponseData), nil
}
