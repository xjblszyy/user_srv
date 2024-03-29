// Code generated by goa v3.0.2, DO NOT EDIT.
//
// user HTTP client encoders and decoders
//
// Command:
// $ goa gen user-srv/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	user "user-srv/gen/user"
	userviews "user-srv/gen/user/views"

	goahttp "goa.design/goa/v3/http"
)

// BuildRetrieveRequest instantiates a HTTP request object with method and path
// set to call the "user" service "retrieve" endpoint
func (c *Client) BuildRetrieveRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RetrieveUserPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "retrieve", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeRetrieveRequest returns an encoder for requests sent to the user
// retrieve server.
func EncodeRetrieveRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.RetrievePayload)
		if !ok {
			return goahttp.ErrInvalidType("user", "retrieve", "*user.RetrievePayload", v)
		}
		req.Header.Set("Authorization", p.Token)
		return nil
	}
}

// DecodeRetrieveResponse returns a decoder for responses returned by the user
// retrieve endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeRetrieveResponse may return the following errors:
//	- "invalide_token" (type *goa.ServiceError): http.StatusBadRequest
//	- "not_found" (type *goa.ServiceError): http.StatusNotFound
//	- error: internal error
func DecodeRetrieveResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body RetrieveResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "retrieve", err)
			}
			p := NewRetrieveResponseDataOK(&body)
			view := "default"
			vres := &userviews.ResponseData{p, view}
			if err = userviews.ValidateResponseData(vres); err != nil {
				return nil, goahttp.ErrValidationError("user", "retrieve", err)
			}
			res := user.NewResponseData(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body RetrieveInvalideTokenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "retrieve", err)
			}
			err = ValidateRetrieveInvalideTokenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "retrieve", err)
			}
			return nil, NewRetrieveInvalideToken(&body)
		case http.StatusNotFound:
			var (
				body RetrieveNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "retrieve", err)
			}
			err = ValidateRetrieveNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "retrieve", err)
			}
			return nil, NewRetrieveNotFound(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "retrieve", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "user" service "create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateUserPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the user create
// server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.AddUser)
		if !ok {
			return goahttp.ErrInvalidType("user", "create", "*user.AddUser", v)
		}
		body := NewCreateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("user", "create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the user
// create endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeCreateResponse may return the following errors:
//	- "email_exist" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body CreateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "create", err)
			}
			p := NewCreateResponseDataCreated(&body)
			view := "default"
			vres := &userviews.ResponseData{p, view}
			if err = userviews.ValidateResponseData(vres); err != nil {
				return nil, goahttp.ErrValidationError("user", "create", err)
			}
			res := user.NewResponseData(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body CreateEmailExistResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "create", err)
			}
			err = ValidateCreateEmailExistResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "create", err)
			}
			return nil, NewCreateEmailExist(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildSigninRequest instantiates a HTTP request object with method and path
// set to call the "user" service "signin" endpoint
func (c *Client) BuildSigninRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SigninUserPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "signin", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSigninRequest returns an encoder for requests sent to the user signin
// server.
func EncodeSigninRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.Signin)
		if !ok {
			return goahttp.ErrInvalidType("user", "signin", "*user.Signin", v)
		}
		body := NewSigninRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("user", "signin", err)
		}
		return nil
	}
}

// DecodeSigninResponse returns a decoder for responses returned by the user
// signin endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeSigninResponse may return the following errors:
//	- "email_or_password_error" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeSigninResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body SigninResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "signin", err)
			}
			p := NewSigninResponseDataOK(&body)
			view := "default"
			vres := &userviews.ResponseData{p, view}
			if err = userviews.ValidateResponseData(vres); err != nil {
				return nil, goahttp.ErrValidationError("user", "signin", err)
			}
			res := user.NewResponseData(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body SigninEmailOrPasswordErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "signin", err)
			}
			err = ValidateSigninEmailOrPasswordErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "signin", err)
			}
			return nil, NewSigninEmailOrPasswordError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "signin", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateRequest instantiates a HTTP request object with method and path
// set to call the "user" service "update" endpoint
func (c *Client) BuildUpdateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateUserPath()}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateRequest returns an encoder for requests sent to the user update
// server.
func EncodeUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.UpdateUser)
		if !ok {
			return goahttp.ErrInvalidType("user", "update", "*user.UpdateUser", v)
		}
		req.Header.Set("Authorization", p.Token)
		body := NewUpdateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("user", "update", err)
		}
		return nil
	}
}

// DecodeUpdateResponse returns a decoder for responses returned by the user
// update endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeUpdateResponse may return the following errors:
//	- "invalide_token" (type *goa.ServiceError): http.StatusBadRequest
//	- "not_found" (type *goa.ServiceError): http.StatusNotFound
//	- error: internal error
func DecodeUpdateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusBadRequest:
			var (
				body UpdateInvalideTokenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "update", err)
			}
			err = ValidateUpdateInvalideTokenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "update", err)
			}
			return nil, NewUpdateInvalideToken(&body)
		case http.StatusNotFound:
			var (
				body UpdateNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "update", err)
			}
			err = ValidateUpdateNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "update", err)
			}
			return nil, NewUpdateNotFound(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "update", resp.StatusCode, string(body))
		}
	}
}

// BuildSendEmailRequest instantiates a HTTP request object with method and
// path set to call the "user" service "sendEmail" endpoint
func (c *Client) BuildSendEmailRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SendEmailUserPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "sendEmail", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSendEmailRequest returns an encoder for requests sent to the user
// sendEmail server.
func EncodeSendEmailRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.SendEmail)
		if !ok {
			return goahttp.ErrInvalidType("user", "sendEmail", "*user.SendEmail", v)
		}
		body := NewSendEmailRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("user", "sendEmail", err)
		}
		return nil
	}
}

// DecodeSendEmailResponse returns a decoder for responses returned by the user
// sendEmail endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeSendEmailResponse may return the following errors:
//	- "email_not_found" (type *goa.ServiceError): http.StatusNotFound
//	- error: internal error
func DecodeSendEmailResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body SendEmailResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "sendEmail", err)
			}
			p := NewSendEmailResponseDataOK(&body)
			view := "default"
			vres := &userviews.ResponseData{p, view}
			if err = userviews.ValidateResponseData(vres); err != nil {
				return nil, goahttp.ErrValidationError("user", "sendEmail", err)
			}
			res := user.NewResponseData(vres)
			return res, nil
		case http.StatusNotFound:
			var (
				body SendEmailEmailNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "sendEmail", err)
			}
			err = ValidateSendEmailEmailNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "sendEmail", err)
			}
			return nil, NewSendEmailEmailNotFound(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "sendEmail", resp.StatusCode, string(body))
		}
	}
}

// BuildActiveEmailRequest instantiates a HTTP request object with method and
// path set to call the "user" service "activeEmail" endpoint
func (c *Client) BuildActiveEmailRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		code string
	)
	{
		p, ok := v.(*user.EmailCode)
		if !ok {
			return nil, goahttp.ErrInvalidType("user", "activeEmail", "*user.EmailCode", v)
		}
		code = p.Code
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ActiveEmailUserPath(code)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("user", "activeEmail", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeActiveEmailResponse returns a decoder for responses returned by the
// user activeEmail endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeActiveEmailResponse may return the following errors:
//	- "code_invalide" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeActiveEmailResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ActiveEmailResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "activeEmail", err)
			}
			p := NewActiveEmailResponseDataOK(&body)
			view := "default"
			vres := &userviews.ResponseData{p, view}
			if err = userviews.ValidateResponseData(vres); err != nil {
				return nil, goahttp.ErrValidationError("user", "activeEmail", err)
			}
			res := user.NewResponseData(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ActiveEmailCodeInvalideResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("user", "activeEmail", err)
			}
			err = ValidateActiveEmailCodeInvalideResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("user", "activeEmail", err)
			}
			return nil, NewActiveEmailCodeInvalide(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("user", "activeEmail", resp.StatusCode, string(body))
		}
	}
}
