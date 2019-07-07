// Code generated by goa v3.0.2, DO NOT EDIT.
//
// userProfile gRPC client types
//
// Command:
// $ goa gen user/design

package client

import (
	"unicode/utf8"
	user_profilepb "user/gen/grpc/user_profile/pb"
	userprofile "user/gen/user_profile"
	userprofileviews "user/gen/user_profile/views"

	goa "goa.design/goa/v3/pkg"
)

// NewRetrieveRequest builds the gRPC request type from the payload of the
// "retrieve" endpoint of the "userProfile" service.
func NewRetrieveRequest() *user_profilepb.RetrieveRequest {
	message := &user_profilepb.RetrieveRequest{}
	return message
}

// NewRetrieveResult builds the result type of the "retrieve" endpoint of the
// "userProfile" service from the gRPC response type.
func NewRetrieveResult(message *user_profilepb.RetrieveResponse) *userprofileviews.UserProfileView {
	result := &userprofileviews.UserProfileView{
		ID:       &message.Id,
		Email:    &message.Email,
		Password: &message.Password,
	}
	if message.Avator != "" {
		result.Avator = &message.Avator
	}
	if message.Nickname != "" {
		result.Nickname = &message.Nickname
	}
	return result
}

// NewRetrieveNotFoundError builds the error type of the "retrieve" endpoint of
// the "userProfile" service from the gRPC error response type.
func NewRetrieveNotFoundError(message *user_profilepb.RetrieveNotFoundError) *userprofile.NotFound {
	er := &userprofile.NotFound{
		Message: message.Message_,
		ID:      message.Id,
	}
	return er
}

// NewRetrieveInvalideTokenError builds the error type of the "retrieve"
// endpoint of the "userProfile" service from the gRPC error response type.
func NewRetrieveInvalideTokenError(message *user_profilepb.RetrieveInvalideTokenError) *userprofile.TokenInvalide {
	er := &userprofile.TokenInvalide{
		Message: message.Message_,
		ID:      message.Id,
	}
	return er
}

// NewCreateRequest builds the gRPC request type from the payload of the
// "create" endpoint of the "userProfile" service.
func NewCreateRequest(payload *userprofile.AddUser) *user_profilepb.CreateRequest {
	message := &user_profilepb.CreateRequest{
		Email:    payload.Email,
		Password: payload.Password,
	}
	return message
}

// NewCreateResult builds the result type of the "create" endpoint of the
// "userProfile" service from the gRPC response type.
func NewCreateResult(message *user_profilepb.CreateResponse) string {
	result := message.Field
	return result
}

// NewCreateEmailExistError builds the error type of the "create" endpoint of
// the "userProfile" service from the gRPC error response type.
func NewCreateEmailExistError(message *user_profilepb.CreateEmailExistError) *userprofile.EmailExist {
	er := &userprofile.EmailExist{
		Message: message.Message_,
		ID:      message.Id,
	}
	return er
}

// NewSigninRequest builds the gRPC request type from the payload of the
// "signin" endpoint of the "userProfile" service.
func NewSigninRequest(payload *userprofile.SigninPayload) *user_profilepb.SigninRequest {
	message := &user_profilepb.SigninRequest{
		Email:    payload.Email,
		Password: payload.Password,
	}
	return message
}

// NewSigninResult builds the result type of the "signin" endpoint of the
// "userProfile" service from the gRPC response type.
func NewSigninResult(message *user_profilepb.SigninResponse) string {
	result := message.Field
	return result
}

// NewSigninEmailOrPasswordErrorError builds the error type of the "signin"
// endpoint of the "userProfile" service from the gRPC error response type.
func NewSigninEmailOrPasswordErrorError(message *user_profilepb.SigninEmailOrPasswordErrorError) *userprofile.EmailOrPasswordError {
	er := &userprofile.EmailOrPasswordError{
		Message: message.Message_,
		ID:      message.Id,
	}
	return er
}

// NewUpdateRequest builds the gRPC request type from the payload of the
// "update" endpoint of the "userProfile" service.
func NewUpdateRequest(payload *userprofile.UpdateUser) *user_profilepb.UpdateRequest {
	message := &user_profilepb.UpdateRequest{
		Avator:   payload.Avator,
		Nickname: payload.Nickname,
	}
	return message
}

// NewUpdateInvalideTokenError builds the error type of the "update" endpoint
// of the "userProfile" service from the gRPC error response type.
func NewUpdateInvalideTokenError(message *user_profilepb.UpdateInvalideTokenError) *userprofile.TokenInvalide {
	er := &userprofile.TokenInvalide{
		Message: message.Message_,
		ID:      message.Id,
	}
	return er
}

// ValidateRetrieveResponse runs the validations defined on RetrieveResponse.
func ValidateRetrieveResponse(message *user_profilepb.RetrieveResponse) (err error) {
	err = goa.MergeErrors(err, goa.ValidatePattern("message.email", message.Email, "\\S+@\\S+\\.\\S+"))
	if utf8.RuneCountInString(message.Password) < 6 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.password", message.Password, utf8.RuneCountInString(message.Password), 6, true))
	}
	return
}