package model

// AppError ...
type AppError interface {
	error
	Code() int32
}

// ErrorCode ...
type modelError int32

const (
	// UserNotFound indicates user not found in database
	UserNotFound modelError = iota + 1

	// InvalidUserOrPassword ...
	InvalidUserOrPassword

	// UserNotCreated ...
	UserNotCreated

	// SignJWTFailed ...
	SignJWTFailed
)

// Code ...
func (c modelError) Code() int32 {
	return int32(c)
}

func (c modelError) Error() string {
	switch int(c) {
	case 1:
		return "user not found"
	case 2:
		return "user or password is invalid"
	case 3:
		return "error occured when create user"
	case 4:
		return "sign JWT failed"
	default:
		return ""
	}
}
