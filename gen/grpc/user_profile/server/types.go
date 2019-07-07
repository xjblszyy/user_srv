// Code generated by goa v3.0.2, DO NOT EDIT.
//
// userProfile gRPC server types
//
// Command:
// $ goa gen user/design

package server

import (
	"unicode/utf8"
	user_profilepb "user/gen/grpc/user_profile/pb"
	userprofile "user/gen/user_profile"
	userprofileviews "user/gen/user_profile/views"

	goa "goa.design/goa/v3/pkg"
)

// NewRetrievePayload builds the payload of the "retrieve" endpoint of the
// "userProfile" service from the gRPC request type.
func NewRetrievePayload(token string) *userprofile.RetrievePayload {
	v := &userprofile.RetrievePayload{}
	v.Token = token
	return v
}

// NewRetrieveResponse builds the gRPC response type from the result of the
// "retrieve" endpoint of the "userProfile" service.
func NewRetrieveResponse(result *userprofileviews.UserProfileView) *user_profilepb.RetrieveResponse {
	message := &user_profilepb.RetrieveResponse{}
	if result.ID != nil {
		message.Id = *result.ID
	}
	if result.Email != nil {
		message.Email = *result.Email
	}
	if result.Password != nil {
		message.Password = *result.Password
	}
	if result.Avator != nil {
		message.Avator = *result.Avator
	}
	if result.Nickname != nil {
		message.Nickname = *result.Nickname
	}
	return message
}

// NewRetrieveNotFoundError builds the gRPC error response type from the error
// of the "retrieve" endpoint of the "userProfile" service.
func NewRetrieveNotFoundError(er *userprofile.NotFound) *user_profilepb.RetrieveNotFoundError {
	message := &user_profilepb.RetrieveNotFoundError{
		Message_: er.Message,
		Id:       er.ID,
	}
	return message
}

// NewRetrieveInvalideTokenError builds the gRPC error response type from the
// error of the "retrieve" endpoint of the "userProfile" service.
func NewRetrieveInvalideTokenError(er *userprofile.TokenInvalide) *user_profilepb.RetrieveInvalideTokenError {
	message := &user_profilepb.RetrieveInvalideTokenError{
		Message_: er.Message,
		Id:       er.ID,
	}
	return message
}

// NewCreatePayload builds the payload of the "create" endpoint of the
// "userProfile" service from the gRPC request type.
func NewCreatePayload(message *user_profilepb.CreateRequest) *userprofile.AddUser {
	v := &userprofile.AddUser{
		Email:    message.Email,
		Password: message.Password,
	}
	return v
}

// NewCreateResponse builds the gRPC response type from the result of the
// "create" endpoint of the "userProfile" service.
func NewCreateResponse(result string) *user_profilepb.CreateResponse {
	message := &user_profilepb.CreateResponse{}
	message.Field = result
	return message
}

// NewCreateEmailExistError builds the gRPC error response type from the error
// of the "create" endpoint of the "userProfile" service.
func NewCreateEmailExistError(er *userprofile.EmailExist) *user_profilepb.CreateEmailExistError {
	message := &user_profilepb.CreateEmailExistError{
		Message_: er.Message,
		Id:       er.ID,
	}
	return message
}

// NewSigninPayload builds the payload of the "signin" endpoint of the
// "userProfile" service from the gRPC request type.
func NewSigninPayload(message *user_profilepb.SigninRequest) *userprofile.SigninPayload {
	v := &userprofile.SigninPayload{
		Email:    message.Email,
		Password: message.Password,
	}
	return v
}

// NewSigninResponse builds the gRPC response type from the result of the
// "signin" endpoint of the "userProfile" service.
func NewSigninResponse(result string) *user_profilepb.SigninResponse {
	message := &user_profilepb.SigninResponse{}
	message.Field = result
	return message
}

// NewSigninEmailOrPasswordErrorError builds the gRPC error response type from
// the error of the "signin" endpoint of the "userProfile" service.
func NewSigninEmailOrPasswordErrorError(er *userprofile.EmailOrPasswordError) *user_profilepb.SigninEmailOrPasswordErrorError {
	message := &user_profilepb.SigninEmailOrPasswordErrorError{
		Message_: er.Message,
		Id:       er.ID,
	}
	return message
}

// NewUpdatePayload builds the payload of the "update" endpoint of the
// "userProfile" service from the gRPC request type.
func NewUpdatePayload(message *user_profilepb.UpdateRequest, token string) *userprofile.UpdateUser {
	v := &userprofile.UpdateUser{
		Avator:   message.Avator,
		Nickname: message.Nickname,
	}
	v.Token = token
	return v
}

// NewUpdateResponse builds the gRPC response type from the result of the
// "update" endpoint of the "userProfile" service.
func NewUpdateResponse() *user_profilepb.UpdateResponse {
	message := &user_profilepb.UpdateResponse{}
	return message
}

// NewUpdateInvalideTokenError builds the gRPC error response type from the
// error of the "update" endpoint of the "userProfile" service.
func NewUpdateInvalideTokenError(er *userprofile.TokenInvalide) *user_profilepb.UpdateInvalideTokenError {
	message := &user_profilepb.UpdateInvalideTokenError{
		Message_: er.Message,
		Id:       er.ID,
	}
	return message
}

// ValidateCreateRequest runs the validations defined on CreateRequest.
func ValidateCreateRequest(message *user_profilepb.CreateRequest) (err error) {
	err = goa.MergeErrors(err, goa.ValidatePattern("message.email", message.Email, "\\S+@\\S+\\.\\S+"))
	if utf8.RuneCountInString(message.Password) < 6 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.password", message.Password, utf8.RuneCountInString(message.Password), 6, true))
	}
	return
}