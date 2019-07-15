package app

import (
	"log"
	"mime/multipart"
	file "user-srv/gen/file"
	"user-srv/utils/oss"
)

// FileUploadDecoderFunc implements the multipart decoder for service "file"
// endpoint "upload". The decoder must populate the argument p after encoding.
func FileUploadDecoderFunc(mr *multipart.Reader, p **file.UploadPayload) error {
	// Add multipart request decoder logic here
	form , err := mr.ReadForm(10<<20)
	if err != nil{
		return err
	}

	FileHeader := form.File["file"][0]
	Image, err:= FileHeader.Open()
	defer Image.Close()
	if err != nil{
		return err
	}

	url, err := oss.UploadFile(FileHeader.Filename, Image)
	if err != nil{
		return err
	}

	log.Println(url)
	pl := &file.UploadPayload{}
	pl.File = url
	*p = pl
	return nil
}

// FileUploadEncoderFunc implements the multipart encoder for service "file"
// endpoint "upload".
func FileUploadEncoderFunc(mw *multipart.Writer, p *file.UploadPayload) error {
	// Add multipart request encoder logic here
	return nil
}
