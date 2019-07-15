package design

import . "goa.design/goa/v3/dsl"

var _ = Service("user", func() {
	Description("The user service makes it possible to view, add or update user info.")

	HTTP(func() {
		Path("/user")
	})

	Method("retrieve", func() {
		Description("Show user info by Token")

		Security(JWTAuth)

		Payload(func() {
			TokenField(1, "token", String, func() {
				Description("JWT used for authentication")
			})
			Required("token")
		})
		Result(ResponseData)
		Error("not_found", ErrorResult, "User not found")
		Error("invalide_token", ErrorResult, "invalide token")
		HTTP(func() {
			GET("/current_user")
			Response(StatusOK)
			Response("invalide_token", StatusBadRequest)
			Response("not_found", StatusNotFound)
		})
		GRPC(func() {
			Response(CodeOK)
			Response("not_found", CodeNotFound)
			Response("invalide_token", CodeInvalidArgument)
		})
	})

	Method("create", func() {
		Description("Add new user")
		Payload(AddUser)
		Result(ResponseData)
		Error("email_exist", ErrorResult, "email exist")
		HTTP(func() {
			POST("/register")
			Response(StatusCreated)
			Response("email_exist", StatusBadRequest)
		})
		GRPC(func() {
			Response(CodeOK)
			Response("email_exist",CodeAlreadyExists)
		})
	})

	Method("signin", func() {
		Description("Creates a valid JWT")

		Payload(Signin)
		Result(ResponseData)
		Error("email_or_password_error", ErrorResult, "email or password error")

		HTTP(func() {
			POST("/signin")
			Response(StatusOK)
			Response("email_or_password_error", StatusBadRequest)
		})

		GRPC(func() {
			Response(CodeOK)
			Response("email_or_password_error", CodeInvalidArgument)
		})
	})

	Method("update", func(){
		Description("Update avatar and nickname about user")

		Security(JWTAuth)
		Payload(UpdateUser)
		Error("not_found", ErrorResult, "user not found")
		Error("invalide_token", ErrorResult, "invalide token")
		HTTP(func(){
			PUT("/update")
			Response("invalide_token", StatusBadRequest)
			Response("not_found", StatusNotFound)
			Response(StatusNoContent)
		})
		GRPC(func(){
			Response(CodeOK)
			Response("invalide_token", CodeInvalidArgument)
			Response("not_found", CodeNotFound)
		})
	})

	Method("sendEmail", func() {
		Description("Send email to active user")
		Payload(SendEmail)
		Result(ResponseData)
		Error("email_not_found", ErrorResult, "email not found")
		HTTP(func(){
			POST("/email")
			Response("email_not_found", StatusNotFound)
			Response(StatusOK)
		})
		GRPC(func(){
			Response(CodeOK)
			Response("email_not_found", CodeNotFound)
		})
	})

	Method("activeEmail", func() {
		Description("Active email to user")
		Payload(EmailCode)
		Result(ResponseData)
		Error("code_invalide", ErrorResult, "code_invalide")
		HTTP(func(){
			GET("/email/{code}")
			Response("code_invalide", StatusBadRequest)
			Response(StatusOK)
		})
		GRPC(func(){
			Response(CodeOK)
			Response("code_invalide", CodeInvalidArgument)
		})
	})

})