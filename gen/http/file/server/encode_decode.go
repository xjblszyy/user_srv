// Code generated by goa v3.0.2, DO NOT EDIT.
//
// file HTTP server encoders and decoders
//
// Command:
// $ goa gen user-srv/design

package server

import (
	"context"
	"net/http"
	file "user-srv/gen/file"
	fileviews "user-srv/gen/file/views"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeUploadResponse returns an encoder for responses returned by the file
// upload endpoint.
func EncodeUploadResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*fileviews.ResponseData)
		enc := encoder(ctx, w)
		body := NewUploadResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUploadRequest returns a decoder for requests sent to the file upload
// endpoint.
func DecodeUploadRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var payload *file.UploadPayload
		if err := decoder(r).Decode(&payload); err != nil {
			return nil, goa.DecodePayloadError(err.Error())
		}

		return payload, nil
	}
}

// NewFileUploadDecoder returns a decoder to decode the multipart request for
// the "file" service "upload" endpoint.
func NewFileUploadDecoder(mux goahttp.Muxer, fileUploadDecoderFn FileUploadDecoderFunc) func(r *http.Request) goahttp.Decoder {
	return func(r *http.Request) goahttp.Decoder {
		return goahttp.EncodingFunc(func(v interface{}) error {
			mr, merr := r.MultipartReader()
			if merr != nil {
				return merr
			}
			p := v.(**file.UploadPayload)
			if err := fileUploadDecoderFn(mr, p); err != nil {
				return err
			}
			return nil
		})
	}
}

// EncodeUploadError returns an encoder for errors returned by the upload file
// endpoint.
func EncodeUploadError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "file_upload_err":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewUploadFileUploadErrResponseBody(res)
			w.Header().Set("goa-error", "file_upload_err")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}
