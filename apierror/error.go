package apierror

const (
	// StatusUnauthorized is used when user do not have authentication credeatials and cannot perform this action.
	StatusUnauthorized = 1

	// StatusInvalidRequest is user when client provided invalid request body.
	StatusInvalidRequest = 2

	// StatusServerError is used when server performed invalid operation and processing cannot be finished.
	StatusServerError = 666
)

var (
	// APIErrorUnauthorized is used when user do not have authentication credeatials and cannot perform this action.
	APIErrorUnauthorized = APIError{
		Description: "Cannot get authorization information. It is likely that you're unauthorized.",
		ErrorCode:   StatusUnauthorized}

	// APIInvalidRequest is user when client provided invalid request body.
	APIInvalidRequest = APIError{
		Description: "Request body is invalid, please make sure that you provided correct fields.",
		ErrorCode:   StatusInvalidRequest}

	// APIServerError is used when server performed invalid operation and processing cannot be finished.
	APIServerError = APIError{
		Description: "Server error occured, cannot process request. Please, contact administration for further information.",
		ErrorCode:   StatusServerError}
)

// APIError repesents API errors of Chlorine endpoints.
type APIError struct {
	Description string `json:"description,omitempty"`
	ErrorCode   int    `json:"error_code,omitempty"`
}
