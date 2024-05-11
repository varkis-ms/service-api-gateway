package model

var (
	ErrNoUser     = New(nil, "the specified user does not exist")
	ErrBadRequest = New(nil, "the request to the server contains a syntax error")
)

type AppError struct {
	Err     error  `json:"-"`
	Message string `json:"message,omitempty"`
}

func (e *AppError) Error() string { return e.Message }

func New(err error, message string) *AppError {
	return &AppError{
		Err:     err,
		Message: message,
	}
}
