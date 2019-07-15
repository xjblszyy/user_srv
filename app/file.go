package app

import (
	"context"
	"log"
	file "user-srv/gen/file"
)

// file service example implementation.
// The example methods log the requests and return zero values.
type filesrvc struct {
	logger *log.Logger
}

// NewFile returns the file service implementation.
func NewFile(logger *log.Logger) file.Service {
	return &filesrvc{logger}
}

// Upload static file
func (s *filesrvc) Upload(ctx context.Context, p *file.UploadPayload) (res *file.ResponseData, err error) {
	res = &file.ResponseData{}
	s.logger.Print("file.upload")
	res.Code = 200
	res.Message = "success"
	res.Data = p.File
	return res, nil
}
