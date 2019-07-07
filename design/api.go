package design

import . "goa.design/goa/v3/dsl"


var _ = API("user", func() {
	Title("User Service")
	Description("HTTP service for managing user service")

	Server("user", func() {
		Description("user hosts the userProfile, file, email services.")
		Services("userProfile", "email", "file", "swagger")
		Host("localhost", func() {
			Description("default host")
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})

// JWTAuth defines a security scheme that uses JWT tokens.
var JWTAuth = JWTSecurity("jwt", func() {
	Description(`Secures endpoint by requiring a valid JWT token retrieved via the signin endpoint. Supports scopes "api:read" and "api:write".`)
	Scope("api:read", "Read-only access")
	Scope("api:write", "Read and write access")
})

var UserProfile = ResultType("UserProfile", func() {
	Description("UserProfile describes a user info in userProfile services.")
	Reference(AddUser)
	TypeName("UserProfile")

	Attributes(func() {
		Attribute("id", String, "ID is the unique id of the userProfile.", func() {
			Example("123abc")
			Meta("rpc:tag", "6")
		})
		Field(2, "email", func() {
			Pattern(`\S+@\S+\.\S+`)
			Example("123@456.com")
			Meta("rpc:tag", "2")
		})
		Field(3, "password")
		Field(4, "avator")
		Field(5, "nickname")
	})

	View("default", func() {
		Attribute("id")
		Attribute("email")
		Attribute("avator")
		Attribute("nickname")
	})

	Required("id", "email", "password")
})

var AddUser = Type("AddUser", func() {
	Description("AddUser describes a user retrieved by userProfile services.")

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
	Description("UpdateUser describes update user info by userProfile services.")
	TypeName("UpdateUser")
	Attribute("avator", String, "avator of userProfile", func() {
		Example("https://www.baidu.com/img/bd_logo1.png?where=super")
		Meta("rpc:tag", "1")
	})
	Attribute("nickname", String, "nickname of userProfile", func() {
		Example("Bobby")
		Meta("rpc:tag", "2")
	})
	TokenField(3, "token", String, func() {
		Description("JWT used for authentication")
	})

	Required("avator", "nickname", "token")
})

// Creds defines the credentials to use for authenticating to service methods.
var Creds = Type("Creds", func() {
	Field(1, "jwt", String, "JWT token", func() {
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	})
	Required("jwt")
})

var NotFound = Type("NotFound", func() {
	Description("NotFound is the type returned when attempting to show or delete a user that does not exist.")
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("user 1 not found")
		Meta("rpc:tag", "1")
	})
	Field(2, "id", String, "ID of missing user")
	Required("message", "id")
})

var EmailExist = Type("EmailExist", func() {
	Description("The email is exist and can not add user")
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("email is exist")
		Meta("rpc:tag", "1")
	})
	Field(2, "id", String, "ID of missing user")
	Required("message", "id")
})

var EmailOrPasswordError = Type("EmailOrPasswordError", func() {
	Description("The email is error or password error")
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("email or password error")
		Meta("rpc:tag", "1")
	})
	Field(2, "id", String, "ID of missing user")
	Required("message", "id")
})


var TokenInvalide = Type("TokenInvalide", func() {
	Description("The Token is invalid")
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("invalid token")
		Meta("rpc:tag", "1")
	})
	Field(2, "id", String, "ID of missing user")
	Required("message", "id")
})

var EmailInvalide = Type("EmailInvalide", func() {
	Description("The Email is invalid")
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("invalid email")
		Meta("rpc:tag", "1")
	})
	Field(2, "id", String, "ID of missing user")
	Required("message", "id")
})

var EmailNotFound = Type("EmailNotFound", func() {
	Description("The Email is not found")
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("email not found")
		Meta("rpc:tag", "1")
	})
	Field(2, "id", String, "ID of missing user")
	Required("message", "id")
})

var FileUploadErr = Type("FileUploadErr", func() {
	Description("File upload error")
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("file upload error")
		Meta("rpc:tag", "1")
	})
	Field(2, "id", String, "ID of missing user")
	Required("message", "id")
})