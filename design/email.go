package design

import . "goa.design/goa/v3/dsl"

var _ = Service("email", func() {
	Description("The email service makes it possible to active user and send active email.")

	HTTP(func() {
		Path("/email")
	})

	Method("active", func() {
		Description("Active user by email code")

		Payload(func() {
			Field(1, "code", String, "operand")
			Required("code")
		})
		Result(String)
		Error("invalide_code", EmailInvalide, "invalide code")
		HTTP(func() {
			GET("/active/{code}")
			Response(StatusOK)
			Response("invalide_code", StatusBadRequest)
		})
		GRPC(func() {
			Response(CodeOK)
			Response("invalide_code", CodeNotFound)
		})
	})

	Method("send_email", func() {
		Description("Send email to active user")

		Payload(func() {
			Attribute("email", String, "email of userProfile", func() {
				Pattern(`\S+@\S+\.\S+`)
				Example("123@456.com")
				Meta("rpc:tag", "1")
			})
			Required("email")
		})
		Result(String)
		Error("invalide_email", EmailInvalide, "invalide email")
		Error("email_not_found", EmailNotFound, "email not found")
		HTTP(func() {
			POST("/active")
			Response(StatusOK)
			Response("invalide_email", StatusBadRequest)
			Response("email_not_found", StatusBadRequest)
		})
		GRPC(func() {
			Response(CodeOK)
			Response("invalide_email", CodeNotFound)
			Response("email_not_found", CodeNotFound)
		})
	})

})