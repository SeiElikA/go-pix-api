package models

import "fmt"

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}
