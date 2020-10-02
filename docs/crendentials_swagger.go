package docs

import "github.com/babyplug/api-challenge-gin-gorm/dto"

// swagger:route POST /api/login Authentication authenEndPoint
// Login route for authentication.
// responses:
//   200: loginResponse

// swagger:parameters authenEndPoint
type credentailsParamsWrapper struct {
	// in:body
	Body dto.Credentials
}

// Response for login.
// swagger:response loginResponse
type credentialsResponseWrapper struct {
	// in:body
	Body dto.CredentialsResponse
}
