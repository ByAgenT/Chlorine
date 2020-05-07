package ws

import "fmt"

var (
	ParseError = Response{
		Type:        TypeResponse,
		Status:      StatusError,
		Description: "Cannot parse request data",
	}
)

func generateActionNotFound(actionName string) Response {
	return Response{
		Type:        TypeResponse,
		Status:      StatusError,
		Description: fmt.Sprintf("Action '%s' not found", actionName),
		Body:        nil,
	}
}
