// Code generated by goa v3.0.2, DO NOT EDIT.
//
// file HTTP server types
//
// Command:
// $ goa gen user/design

package server

import (
	file "user/gen/file"

	goa "goa.design/goa/v3/pkg"
)

// UploadRequestBody is the type of the "file" service "upload" endpoint HTTP
// request body.
type UploadRequestBody struct {
	// file
	File *string `form:"file,omitempty" json:"file,omitempty" xml:"file,omitempty"`
}

// UploadFileUploadErrResponseBody is the type of the "file" service "upload"
// endpoint HTTP response body for the "file_upload_err" error.
type UploadFileUploadErrResponseBody struct {
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
	// ID of missing user
	ID string `form:"id" json:"id" xml:"id"`
}

// NewUploadFileUploadErrResponseBody builds the HTTP response body from the
// result of the "upload" endpoint of the "file" service.
func NewUploadFileUploadErrResponseBody(res *file.FileUploadErr) *UploadFileUploadErrResponseBody {
	body := &UploadFileUploadErrResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewUploadPayload builds a file service upload endpoint payload.
func NewUploadPayload(body *UploadRequestBody) *file.UploadPayload {
	v := &file.UploadPayload{
		File: *body.File,
	}
	return v
}

// ValidateUploadRequestBody runs the validations defined on UploadRequestBody
func ValidateUploadRequestBody(body *UploadRequestBody) (err error) {
	if body.File == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("file", "body"))
	}
	return
}
