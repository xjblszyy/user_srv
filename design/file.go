package design

import . "goa.design/goa/v3/dsl"

var _ = Service("file", func() {
	Description("The file service makes it possible to upload static file.")

	HTTP(func() {
		Path("/file")
	})

	Method("upload", func() {
		Description("Upload static file")

		Payload(func() {
			Field(1, "file", String, "file")
			Required("file")
		})
		Result(String)
		Error("file_upload_err", FileUploadErr, "file upload error")
		HTTP(func() {
			POST("/upload")
			Response(StatusOK)
			MultipartRequest()
			Response("file_upload_err", StatusBadRequest)
		})
		GRPC(func() {
			Response(CodeOK)
			Response("file_upload_err", CodeNotFound)
		})
	})
})