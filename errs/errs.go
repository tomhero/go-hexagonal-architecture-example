package errs

type AppError struct {
	Code    int
	Message string
}

// NOTE : @Implement จาก base error interface เหมือนใน Java ด้วย Receiver function
func (e AppError) Error() string {
	return e.Message
}
