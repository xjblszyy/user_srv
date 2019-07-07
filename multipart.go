package user

import (
	"fmt"
	"mime/multipart"
	"user/gen/file"
	"user/utils/oss"
)

// FileUploadDecoderFunc implements the multipart decoder for service "file"
// endpoint "upload". The decoder must populate the argument p after encoding.
func FileUploadDecoderFunc(mr *multipart.Reader, p **file.UploadPayload) error {
	// Add multipart request decoder logic here
	form , err := mr.ReadForm(10<<20)
	if err != nil{
		return err
	}else{
		FileHeader := form.File["file"][0]
		Image, err:= FileHeader.Open()
		if err != nil{
			return err
		}else{
			url, err := oss.UploadFile(FileHeader.Filename, Image)
			if err != nil{
				return err
			}else{
				fmt.Println(url)
				(**p).File = url
				return nil
			}
		}
	}
}

// FileUploadEncoderFunc implements the multipart encoder for service "file"
// endpoint "upload".
func FileUploadEncoderFunc(mw *multipart.Writer, p *file.UploadPayload) error {
	// Add multipart request encoder logic here
	return nil
}
