package main

import (
	"context"
	"log"
	"net"
	"net/url"
	"sync"
	email "user/gen/email"
	file "user/gen/file"
	emailpb "user/gen/grpc/email/pb"
	emailsvr "user/gen/grpc/email/server"
	filepb "user/gen/grpc/file/pb"
	filesvr "user/gen/grpc/file/server"
	user_profilepb "user/gen/grpc/user_profile/pb"
	userprofilesvr "user/gen/grpc/user_profile/server"
	userprofile "user/gen/user_profile"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcmdlwr "goa.design/goa/v3/grpc/middleware"
	"goa.design/goa/v3/middleware"
	"google.golang.org/grpc"
)

// handleGRPCServer starts configures and starts a gRPC server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleGRPCServer(ctx context.Context, u *url.URL, userProfileEndpoints *userprofile.Endpoints, emailEndpoints *email.Endpoints, fileEndpoints *file.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = middleware.NewLogger(logger)
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to gRPC requests and
	// responses.
	var (
		userProfileServer *userprofilesvr.Server
		emailServer       *emailsvr.Server
		fileServer        *filesvr.Server
	)
	{
		userProfileServer = userprofilesvr.New(userProfileEndpoints, nil)
		emailServer = emailsvr.New(emailEndpoints, nil)
		fileServer = filesvr.New(fileEndpoints, nil)
	}

	// Initialize gRPC server with the middleware.
	srv := grpc.NewServer(
		grpcmiddleware.WithUnaryServerChain(
			grpcmdlwr.UnaryRequestID(),
			grpcmdlwr.UnaryServerLog(adapter),
		),
	)

	// Register the servers.
	user_profilepb.RegisterUserProfileServer(srv, userProfileServer)
	emailpb.RegisterEmailServer(srv, emailServer)
	filepb.RegisterFileServer(srv, fileServer)

	for svc, info := range srv.GetServiceInfo() {
		for _, m := range info.Methods {
			logger.Printf("serving gRPC method %s", svc+"/"+m.Name)
		}
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start gRPC server in a separate goroutine.
		go func() {
			lis, err := net.Listen("tcp", u.Host)
			if err != nil {
				errc <- err
			}
			logger.Printf("gRPC server listening on %q", u.Host)
			errc <- srv.Serve(lis)
		}()

		<-ctx.Done()
		logger.Printf("shutting down gRPC server at %q", u.Host)
		srv.Stop()
	}()
}
