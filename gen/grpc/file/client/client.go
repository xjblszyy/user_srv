// Code generated by goa v3.0.2, DO NOT EDIT.
//
// file gRPC client
//
// Command:
// $ goa gen user/design

package client

import (
	"context"
	filepb "user/gen/grpc/file/pb"

	goagrpc "goa.design/goa/v3/grpc"
	goapb "goa.design/goa/v3/grpc/pb"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli filepb.FileClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the file service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: filepb.NewFileClient(cc),
		opts:    opts,
	}
}

// Upload calls the "Upload" function in filepb.FileClient interface.
func (c *Client) Upload() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildUploadFunc(c.grpccli, c.opts...),
			EncodeUploadRequest,
			DecodeUploadResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *filepb.UploadFileUploadErrError:
				return nil, NewUploadFileUploadErrError(message)
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}
