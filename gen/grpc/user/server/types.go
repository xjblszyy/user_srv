// Code generated by goa v3.0.2, DO NOT EDIT.
//
// user gRPC server types
//
// Command:
// $ goa gen user-srv/design

package server

import (
	"unicode/utf8"
	userpb "user-srv/gen/grpc/user/pb"
	user "user-srv/gen/user"
	userviews "user-srv/gen/user/views"

	goa "goa.design/goa/v3/pkg"
)

// NewRetrievePayload builds the payload of the "retrieve" endpoint of the
// "user" service from the gRPC request type.
func NewRetrievePayload(token string) *user.RetrievePayload {
	v := &user.RetrievePayload{}
	v.Token = token
	return v
}

// NewRetrieveResponse builds the gRPC response type from the result of the
// "retrieve" endpoint of the "user" service.
func NewRetrieveResponse(result *userviews.ResponseDataView) *userpb.RetrieveResponse {
	message := &userpb.RetrieveResponse{}
	if result.Code != nil {
		message.Code = int32(*result.Code)
	}
	if result.Message != nil {
		message.Message_ = *result.Message
	}
	if result.Data != nil {
		message.Data = *result.Data
	}
	if result.Code == nil {
		message.Code = 200
	}
	if result.Message == nil {
		message.Message_ = "success"
	}
	return message
}

// NewCreatePayload builds the payload of the "create" endpoint of the "user"
// service from the gRPC request type.
func NewCreatePayload(message *userpb.CreateRequest) *user.AddUser {
	v := &user.AddUser{
		Email:    message.Email,
		Password: message.Password,
	}
	return v
}

// NewCreateResponse builds the gRPC response type from the result of the
// "create" endpoint of the "user" service.
func NewCreateResponse(result *userviews.ResponseDataView) *userpb.CreateResponse {
	message := &userpb.CreateResponse{}
	if result.Code != nil {
		message.Code = int32(*result.Code)
	}
	if result.Message != nil {
		message.Message_ = *result.Message
	}
	if result.Data != nil {
		message.Data = *result.Data
	}
	if result.Code == nil {
		message.Code = 200
	}
	if result.Message == nil {
		message.Message_ = "success"
	}
	return message
}

// NewSigninPayload builds the payload of the "signin" endpoint of the "user"
// service from the gRPC request type.
func NewSigninPayload(message *userpb.SigninRequest) *user.Signin {
	v := &user.Signin{
		Email:    message.Email,
		Password: message.Password,
	}
	return v
}

// NewSigninResponse builds the gRPC response type from the result of the
// "signin" endpoint of the "user" service.
func NewSigninResponse(result *userviews.ResponseDataView) *userpb.SigninResponse {
	message := &userpb.SigninResponse{}
	if result.Code != nil {
		message.Code = int32(*result.Code)
	}
	if result.Message != nil {
		message.Message_ = *result.Message
	}
	if result.Data != nil {
		message.Data = *result.Data
	}
	if result.Code == nil {
		message.Code = 200
	}
	if result.Message == nil {
		message.Message_ = "success"
	}
	return message
}

// NewUpdatePayload builds the payload of the "update" endpoint of the "user"
// service from the gRPC request type.
func NewUpdatePayload(message *userpb.UpdateRequest, token string) *user.UpdateUser {
	v := &user.UpdateUser{
		Avatar:   message.Avatar,
		Nickname: message.Nickname,
	}
	v.Token = token
	return v
}

// NewUpdateResponse builds the gRPC response type from the result of the
// "update" endpoint of the "user" service.
func NewUpdateResponse() *userpb.UpdateResponse {
	message := &userpb.UpdateResponse{}
	return message
}

// NewSendEmailPayload builds the payload of the "sendEmail" endpoint of the
// "user" service from the gRPC request type.
func NewSendEmailPayload(message *userpb.SendEmailRequest) *user.SendEmail {
	v := &user.SendEmail{
		Email: message.Email,
	}
	return v
}

// NewSendEmailResponse builds the gRPC response type from the result of the
// "sendEmail" endpoint of the "user" service.
func NewSendEmailResponse(result *userviews.ResponseDataView) *userpb.SendEmailResponse {
	message := &userpb.SendEmailResponse{}
	if result.Code != nil {
		message.Code = int32(*result.Code)
	}
	if result.Message != nil {
		message.Message_ = *result.Message
	}
	if result.Data != nil {
		message.Data = *result.Data
	}
	if result.Code == nil {
		message.Code = 200
	}
	if result.Message == nil {
		message.Message_ = "success"
	}
	return message
}

// NewActiveEmailPayload builds the payload of the "activeEmail" endpoint of
// the "user" service from the gRPC request type.
func NewActiveEmailPayload(message *userpb.ActiveEmailRequest) *user.EmailCode {
	v := &user.EmailCode{
		Code: message.Code,
	}
	return v
}

// NewActiveEmailResponse builds the gRPC response type from the result of the
// "activeEmail" endpoint of the "user" service.
func NewActiveEmailResponse(result *userviews.ResponseDataView) *userpb.ActiveEmailResponse {
	message := &userpb.ActiveEmailResponse{}
	if result.Code != nil {
		message.Code = int32(*result.Code)
	}
	if result.Message != nil {
		message.Message_ = *result.Message
	}
	if result.Data != nil {
		message.Data = *result.Data
	}
	if result.Code == nil {
		message.Code = 200
	}
	if result.Message == nil {
		message.Message_ = "success"
	}
	return message
}

// ValidateCreateRequest runs the validations defined on CreateRequest.
func ValidateCreateRequest(message *userpb.CreateRequest) (err error) {
	err = goa.MergeErrors(err, goa.ValidatePattern("message.email", message.Email, "\\S+@\\S+\\.\\S+"))
	if utf8.RuneCountInString(message.Password) < 6 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.password", message.Password, utf8.RuneCountInString(message.Password), 6, true))
	}
	return
}

// ValidateSendEmailRequest runs the validations defined on SendEmailRequest.
func ValidateSendEmailRequest(message *userpb.SendEmailRequest) (err error) {
	err = goa.MergeErrors(err, goa.ValidatePattern("message.email", message.Email, "\\S+@\\S+\\.\\S+"))
	return
}
