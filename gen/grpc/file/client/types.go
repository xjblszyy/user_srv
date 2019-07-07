// Code generated by goa v3.0.2, DO NOT EDIT.
//
// file gRPC client types
//
// Command:
// $ goa gen user/design

package client

import (
	file "user/gen/file"
	filepb "user/gen/grpc/file/pb"
)

// NewUploadRequest builds the gRPC request type from the payload of the
// "upload" endpoint of the "file" service.
func NewUploadRequest(payload *file.UploadPayload) *filepb.UploadRequest {
	message := &filepb.UploadRequest{
		File: payload.File,
	}
	return message
}

// NewUploadResult builds the result type of the "upload" endpoint of the
// "file" service from the gRPC response type.
func NewUploadResult(message *filepb.UploadResponse) string {
	result := message.Field
	return result
}

// NewUploadFileUploadErrError builds the error type of the "upload" endpoint
// of the "file" service from the gRPC error response type.
func NewUploadFileUploadErrError(message *filepb.UploadFileUploadErrError) *file.FileUploadErr {
	er := &file.FileUploadErr{
		Message: message.Message_,
		ID:      message.Id,
	}
	return er
}