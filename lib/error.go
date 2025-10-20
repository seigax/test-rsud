package lib

import (
	"net/http"
)

type CustomError struct {
	Message  string
	Field    string
	Code     int
	HTTPCode int
}

func (err CustomError) Error() string {
	return err.Message
}

var (
	ErrorForbidden = CustomError{
		Message:  "Forbidden",
		Code:     1000,
		HTTPCode: http.StatusForbidden,
	}

	ErrorInvalidParameter = CustomError{
		Message:  "Invalid Parameter",
		Code:     1001,
		HTTPCode: http.StatusUnprocessableEntity,
	}

	ErrorNotFound = CustomError{
		Message:  "Not Found",
		Code:     1002,
		HTTPCode: http.StatusNotFound,
	}

	ErrorInternalServer = CustomError{
		Message:  "Internal Server Error",
		Code:     1003,
		HTTPCode: http.StatusInternalServerError,
	}

	ErrorInvalidToken = CustomError{
		Message:  "You need a valid token!",
		Code:     1004,
		HTTPCode: http.StatusUnauthorized,
	}

	ErrorEmailAlreadyRegistered = CustomError{
		Message:  "Email Already Registered",
		Code:     1005,
		HTTPCode: http.StatusBadRequest,
	}

	ErrorPhoneAlreadyRegistered = CustomError{
		Message:  "Phone Already Registered",
		Code:     1006,
		HTTPCode: http.StatusBadRequest,
	}

	ErrorWrongEmailOrPassword = CustomError{
		Message:  "Wrong Email Or Password",
		Code:     1007,
		HTTPCode: http.StatusBadRequest,
	}

	ErrorExpiredToken = CustomError{
		Message:  "Token Expired!",
		Code:     1008,
		HTTPCode: http.StatusUnauthorized,
	}

	ErrorOnlyInternal = CustomError{
		Message:  "Only Internal User Can Login!",
		Code:     1009,
		HTTPCode: http.StatusUnauthorized,
	}
	ErrorRoleNotFarmer = CustomError{
		Message:  "Role not farmer",
		Code:     1009,
		HTTPCode: http.StatusBadRequest,
	}

	NotValidOTP = CustomError{
		Message:  "Not Valid OTP",
		Code:     1010,
		HTTPCode: http.StatusBadRequest,
	}

	OtpExpired = CustomError{
		Message:  "OTP Expired",
		Code:     1011,
		HTTPCode: http.StatusBadRequest,
	}

	PasswordNotSame = CustomError{
		Message:  "Password And Confirmation Password Not Same",
		Code:     1012,
		HTTPCode: http.StatusBadRequest,
	}

	Oauth2Error = CustomError{
		Message:  "Oauth2 Error",
		Code:     1013,
		HTTPCode: http.StatusBadRequest,
	}
)
