package design

import . "goa.design/goa/v3/dsl"

var _ = Service("userProfile", func() {
	Description("The userProfile service makes it possible to view, add or remove user info.")

	HTTP(func() {
		Path("/user")
	})

	Method("retrieve", func() {
		Description("Show userProfile by Token")

		Security(JWTAuth, func() { // Use JWT to auth requests to this endpoint.
			Scope("api:read") // Enforce presence of "api:read" scope in JWT claims.
		})

		Payload(func() {
			TokenField(1, "token", String, func() {
				Description("JWT used for authentication")
			})
			Required("token")
		})
		Result(UserProfile)
		Error("not_found", NotFound, "User not found")
		Error("invalide_token", TokenInvalide, "invalide token")
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
		Description("Add new user and return its ID.")
		Payload(AddUser)
		Result(String)
		Error("email_exist", EmailExist)
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

		Payload(func() {
			Description("Credentials used to authenticate to retrieve JWT token")
			Field(1, "email", String, "Email used to perform signin", func() {
				Example("1@1.com")
			})
			Field(2, "password", String, "Password used to perform signin", func() {
				Example("123456")
			})
			Required("email", "password")
		})

		Result(String)
		Error("email_or_password_error", EmailOrPasswordError)

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

		Security(JWTAuth, func() { // Use JWT to auth requests to this endpoint.
			Scope("api:read") // Enforce presence of "api:read" scope in JWT claims.
		})
		Payload(UpdateUser)
		Error("not_found", NotFound, "user not found")
		Error("invalide_token", TokenInvalide, "invalide token")
		HTTP(func(){
			PUT("/update")
			Response("invalide_token", StatusBadRequest)
			Response(StatusNoContent)
		})
		GRPC(func(){
			Response(CodeOK)
			Response("invalide_token", CodeInvalidArgument)
		})
	})

})