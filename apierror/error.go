package apierror

const (
	// StatusUnauthorized is used when user do not have authentication credeatials and cannot perform this action.
	StatusUnauthorized = 1
)

var (
	// APIErrorUnauthorized is used when user do not have authentication credeatials and cannot perform this action.
	APIErrorUnauthorized = APIError{
		Description: "Cannot get authorization information. It is likely that you're unauthorized.",
		ErrorCode:   StatusUnauthorized}
)

// APIError repesents API errors of Chlorine endpoints.
type APIError struct {
	Description string `json:"description,omitempty"`
	ErrorCode   int    `json:"error_code,omitempty"`
}
