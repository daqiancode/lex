package stderrorcode

type ErrorCode int64

const (
	OK            ErrorCode = 0
	InternalError ErrorCode = 1

	FieldError ErrorCode = 10
)

const (
	UserNotExist     ErrorCode = 100
	EmailExist       ErrorCode = 101
	MobileExist      ErrorCode = 102
	PasswordError    ErrorCode = 103
	EmailNotExist    ErrorCode = 104
	MobileNotExist   ErrorCode = 105
	PasswordTooShort ErrorCode = 106
	OldPasswordError ErrorCode = 107

	AccessTokenEmpty   ErrorCode = 150
	AccessTokenInvalid ErrorCode = 151
	AccessTokenExpired ErrorCode = 152
)
