package response

import "gitlab.devgroup.tech/shkolkovo/romanych/log"

type ErrorResponse struct {
	Message string
	Code    int
}

func NewErrorResponse(Message string, Code int) *ErrorResponse{
	log.GetLogger().Warningf("Error[%d]: %s", Code, Message)
	return &ErrorResponse{
		Message: Message,
		Code:    Code,
	}
}