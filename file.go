package user

import (
	"context"
	"fmt"
	"log"
	file "user/gen/file"
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
func (s *filesrvc) Upload(ctx context.Context, p *file.UploadPayload) (res string, err error) {
	s.logger.Print("file.upload")
	fmt.Println(p.File)
	return
}
