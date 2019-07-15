package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("user-srv", func() {
	Title("UserSrv Service")
	Description("HTTP service for managing user-srv service")

	Server("user", func() {
		Description("user-srv hosts the user, file, email services.")
		Services("user", "swagger", "file")
		Host("localhost", func() {
			Description("default host")
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})

// JWTAuth defines a security scheme that uses JWT tokens.
var JWTAuth = JWTSecurity("jwt", func() {
	Description(`Secures endpoint by requiring a valid JWT token retrieved via the signin endpoint.`)
	//Scope("api:read", "Read-only access")
	//Scope("api:write", "Read and write access")
})

var ResponseData = ResultType("ResponseData", func() {
	Description("The Response")
	TypeName("responseData")
	Attribute("code", Int, "code", func() {
		Default(200)
		Example(200)
		Meta("rpc:tag", "1")
	})
	Attribute("message", String, "message", func() {
		Example("success")
		Default("success")
		Meta("rpc:tag", "2")
	})
	Field(3, "data", String)
	Required("code", "message", "data")
})

var AddUser = Type("AddUser", func() {
	Description("AddUser describes a user retrieved by userProfile services.")
	TypeName("AddUser")
	Attribute("email", String, "email of userProfile", func() {
		Pattern(`\S+@\S+\.\S+`)
		Example("123@456.com")
		Meta("rpc:tag", "1")
	})

	Attribute("password", String, "password of userProfile", func() {
		MinLength(6)
		Example("123456")
		Meta("rpc:tag", "2")
	})

	Required("email", "password")
})

var UpdateUser = Type("UpdateUser", func(){
	Description("UpdateUser describes update user info by user services.")
	TypeName("UpdateUser")
	Attribute("avatar", String, "avatar of user", func() {
		Example("https://www.baidu.com/img/bd_logo1.png?where=super")
		Meta("rpc:tag", "1")
	})
	Attribute("nickname", String, "nickname of user", func() {
		Example("Bobby")
		Meta("rpc:tag", "2")
	})
	TokenField(3, "token", String, func() {
		Description("JWT used for authentication")
	})

	Required("avatar", "nickname", "token")
})

var Signin = Type("Signin", func() {
	Description("user sign in")
	TypeName("Signin")
	Field(1, "email", String, "Email used to perform signin", func() {
		Example("1@1.com")
	})
	Field(2, "password", String, "Password used to perform signin", func() {
		Example("123456")
	})
	Required("email", "password")
})

var SendEmail = Type("SendEmail", func() {
	Description("send email")
	TypeName("SendEmail")
	Attribute("email", String, "email of userProfile", func() {
		Pattern(`\S+@\S+\.\S+`)
		Example("123@456.com")
		Meta("rpc:tag", "1")
	})
	Required("email")
})

var EmailCode = Type("EmailCode", func() {
	Description("email code")
	TypeName("email code")
	Field(1, "code", String, "The code for email to active", func() {
		Example("123456")
		Meta("rpc:tag", "1")
	})
	Required("code")
})

// Creds defines the credentials to use for authenticating to service methods.
var Creds = Type("Creds", func() {
	TypeName("Creds")
	Field(1, "jwt", String, "JWT token", func() {
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	})
	Required("jwt")
})

//var HTTP_200_OK = Type("Ok", func() {
//	Description("Ok")
//	TypeName("Ok")
//	Attribute("message", String, "Message of error", func() {
//		Meta("struct:error:name")
//		Example("success")
//		Meta("rpc:tag", "1")
//	})
//	Attribute("code", Int, "code", func() {
//		Default(200)
//		Example(200)
//		Meta("rpc:tag", "2")
//	})
//	Required("message", "code")
//})
//
//var HTTP_201_CREATED = Type("Created", func() {
//	Description("Created")
//	TypeName("Created")
//	Attribute("message", String, "Message of error", func() {
//		Meta("struct:error:name")
//		Example("created success")
//		Meta("rpc:tag", "1")
//	})
//	Attribute("code", Int, "code", func() {
//		Default(201)
//		Example(201)
//		Meta("rpc:tag", "2")
//	})
//	Required("message", "code")
//})

