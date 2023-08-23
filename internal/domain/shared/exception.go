package shared

import "fmt"

type HttpException struct {
	Code    int
	Message any
}

func (e *HttpException) Error() string {
	return fmt.Sprintf("Error Code: %d, Message: %s", e.Code, e.Message)
}
