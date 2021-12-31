package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

// NOTE : @override จาก base error interface เหมือนใน Java ด้วย Receiver function
func (e AppError) Error() string {
	return e.Message
}

// NOTE : return ออกไปเป็น AppError ที่มาจาก error interface จริงๆ ว้าวซ่า
func NewNotFoundError(message string) error {
	return AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

// NOTE : return ออกไปเป็น AppError ที่มาจาก error interface จริงๆ ว้าวซ่า
func NewUnexpectedError() error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: "unexpected error occures",
	}
}

func NewValidationError(message string) error {
	// defaultMessage := "validation error occures"
	// if message == "" {
	// 	message = defaultMessage
	// }
	return AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}
