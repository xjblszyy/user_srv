// Code generated by goa v3.0.2, DO NOT EDIT.
//
// email gRPC server
//
// Command:
// $ goa gen user/design

package server

import (
	"context"
	email "user/gen/email"
	emailpb "user/gen/grpc/email/pb"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

// Server implements the emailpb.EmailServer interface.
type Server struct {
	ActiveH    goagrpc.UnaryHandler
	SendEmailH goagrpc.UnaryHandler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the email service endpoints.
func New(e *email.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		ActiveH:    NewActiveHandler(e.Active, uh),
		SendEmailH: NewSendEmailHandler(e.SendEmail, uh),
	}
}

// NewActiveHandler creates a gRPC handler which serves the "email" service
// "active" endpoint.
func NewActiveHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeActiveRequest, EncodeActiveResponse)
	}
	return h
}

// Active implements the "Active" method in emailpb.EmailServer interface.
func (s *Server) Active(ctx context.Context, message *emailpb.ActiveRequest) (*emailpb.ActiveResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "active")
	ctx = context.WithValue(ctx, goa.ServiceKey, "email")
	resp, err := s.ActiveH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "invalide_code":
				er := err.(*email.EmailInvalide)
				return nil, goagrpc.NewStatusError(codes.NotFound, err, NewActiveInvalideCodeError(er))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*emailpb.ActiveResponse), nil
}

// NewSendEmailHandler creates a gRPC handler which serves the "email" service
// "send_email" endpoint.
func NewSendEmailHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeSendEmailRequest, EncodeSendEmailResponse)
	}
	return h
}

// SendEmail implements the "SendEmail" method in emailpb.EmailServer interface.
func (s *Server) SendEmail(ctx context.Context, message *emailpb.SendEmailRequest) (*emailpb.SendEmailResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "send_email")
	ctx = context.WithValue(ctx, goa.ServiceKey, "email")
	resp, err := s.SendEmailH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "invalide_email":
				er := err.(*email.EmailInvalide)
				return nil, goagrpc.NewStatusError(codes.NotFound, err, NewSendEmailInvalideEmailError(er))
			case "email_not_found":
				er := err.(*email.EmailNotFound)
				return nil, goagrpc.NewStatusError(codes.NotFound, err, NewSendEmailEmailNotFoundError(er))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*emailpb.SendEmailResponse), nil
}