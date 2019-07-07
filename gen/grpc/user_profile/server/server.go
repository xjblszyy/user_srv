// Code generated by goa v3.0.2, DO NOT EDIT.
//
// userProfile gRPC server
//
// Command:
// $ goa gen user/design

package server

import (
	"context"
	user_profilepb "user/gen/grpc/user_profile/pb"
	userprofile "user/gen/user_profile"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

// Server implements the user_profilepb.UserProfileServer interface.
type Server struct {
	RetrieveH goagrpc.UnaryHandler
	CreateH   goagrpc.UnaryHandler
	SigninH   goagrpc.UnaryHandler
	UpdateH   goagrpc.UnaryHandler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the userProfile service endpoints.
func New(e *userprofile.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		RetrieveH: NewRetrieveHandler(e.Retrieve, uh),
		CreateH:   NewCreateHandler(e.Create, uh),
		SigninH:   NewSigninHandler(e.Signin, uh),
		UpdateH:   NewUpdateHandler(e.Update, uh),
	}
}

// NewRetrieveHandler creates a gRPC handler which serves the "userProfile"
// service "retrieve" endpoint.
func NewRetrieveHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeRetrieveRequest, EncodeRetrieveResponse)
	}
	return h
}

// Retrieve implements the "Retrieve" method in
// user_profilepb.UserProfileServer interface.
func (s *Server) Retrieve(ctx context.Context, message *user_profilepb.RetrieveRequest) (*user_profilepb.RetrieveResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "retrieve")
	ctx = context.WithValue(ctx, goa.ServiceKey, "userProfile")
	resp, err := s.RetrieveH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "not_found":
				er := err.(*userprofile.NotFound)
				return nil, goagrpc.NewStatusError(codes.NotFound, err, NewRetrieveNotFoundError(er))
			case "invalide_token":
				er := err.(*userprofile.TokenInvalide)
				return nil, goagrpc.NewStatusError(codes.InvalidArgument, err, NewRetrieveInvalideTokenError(er))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*user_profilepb.RetrieveResponse), nil
}

// NewCreateHandler creates a gRPC handler which serves the "userProfile"
// service "create" endpoint.
func NewCreateHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeCreateRequest, EncodeCreateResponse)
	}
	return h
}

// Create implements the "Create" method in user_profilepb.UserProfileServer
// interface.
func (s *Server) Create(ctx context.Context, message *user_profilepb.CreateRequest) (*user_profilepb.CreateResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "create")
	ctx = context.WithValue(ctx, goa.ServiceKey, "userProfile")
	resp, err := s.CreateH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "email_exist":
				er := err.(*userprofile.EmailExist)
				return nil, goagrpc.NewStatusError(codes.AlreadyExists, err, NewCreateEmailExistError(er))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*user_profilepb.CreateResponse), nil
}

// NewSigninHandler creates a gRPC handler which serves the "userProfile"
// service "signin" endpoint.
func NewSigninHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeSigninRequest, EncodeSigninResponse)
	}
	return h
}

// Signin implements the "Signin" method in user_profilepb.UserProfileServer
// interface.
func (s *Server) Signin(ctx context.Context, message *user_profilepb.SigninRequest) (*user_profilepb.SigninResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "signin")
	ctx = context.WithValue(ctx, goa.ServiceKey, "userProfile")
	resp, err := s.SigninH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "email_or_password_error":
				er := err.(*userprofile.EmailOrPasswordError)
				return nil, goagrpc.NewStatusError(codes.InvalidArgument, err, NewSigninEmailOrPasswordErrorError(er))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*user_profilepb.SigninResponse), nil
}

// NewUpdateHandler creates a gRPC handler which serves the "userProfile"
// service "update" endpoint.
func NewUpdateHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeUpdateRequest, EncodeUpdateResponse)
	}
	return h
}

// Update implements the "Update" method in user_profilepb.UserProfileServer
// interface.
func (s *Server) Update(ctx context.Context, message *user_profilepb.UpdateRequest) (*user_profilepb.UpdateResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "update")
	ctx = context.WithValue(ctx, goa.ServiceKey, "userProfile")
	resp, err := s.UpdateH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "invalide_token":
				er := err.(*userprofile.TokenInvalide)
				return nil, goagrpc.NewStatusError(codes.InvalidArgument, err, NewUpdateInvalideTokenError(er))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*user_profilepb.UpdateResponse), nil
}
