package core

import (
	"net/http"
)

type ErrorCode = int

const (
	ErrorCodeIncorrectRequestBody        ErrorCode = 10000
	ErrorCodePasswordEncryptFailed       ErrorCode = 10001
	ErrorCodeCreatingToken               ErrorCode = 10002
	ErrorCodeInvalidAccessToken          ErrorCode = 10003
	ErrorCodeExtractingAccessTokenClaims ErrorCode = 10004
	ErrorCodeErrorBadAccessToken         ErrorCode = 10005
	ErrorCodeAccessTokenNotFound         ErrorCode = 10006
	ErrorCodeAccessTokenMalformed        ErrorCode = 10007
	ErrorCodeComparingPasswords          ErrorCode = 10008
	ErrorCodeEndpointSecretKeyNotFound   ErrorCode = 10009
	ErrorCodeAppSecretKeyNotFound        ErrorCode = 10010
	ErrorCodeInvalidAppSecretKey         ErrorCode = 10011
	ErrorCodeEndpointForbidden           ErrorCode = 10012
)

func errorText(code ErrorCode) string {
	switch code {
	case ErrorCodeIncorrectRequestBody:
		return "Incorrect data type in request body. Refer to documentation"
	case ErrorCodePasswordEncryptFailed:
		return "Failed to encrypt password"
	case ErrorCodeCreatingToken:
		return "Failed to create access token"
	case ErrorCodeInvalidAccessToken:
		return "Invalid access token"
	case ErrorCodeExtractingAccessTokenClaims:
		return "Failed to extract access_token claims"
	case ErrorCodeErrorBadAccessToken:
		return "Bad access token"
	case ErrorCodeAccessTokenNotFound:
		return "Access Token not found"
	case ErrorCodeAccessTokenMalformed:
		return "Access Token malformed"
	case ErrorCodeComparingPasswords:
		return "Failed to compare hashed and plain text passwords"
	case ErrorCodeEndpointSecretKeyNotFound:
		return "Server secret Key not found"
	case ErrorCodeAppSecretKeyNotFound:
		return "App secret key not found"
	case ErrorCodeInvalidAppSecretKey:
		return "Invalid app secret key"
	case ErrorCodeEndpointForbidden:
		return "You don't have access to this endpoint"
	default:
		return ""
	}
}

func errorHttpStatusCode(code ErrorCode) int {
	switch code {
	case ErrorCodeIncorrectRequestBody:
		return http.StatusBadRequest
	case ErrorCodePasswordEncryptFailed:
		return http.StatusInternalServerError
	case ErrorCodeCreatingToken:
		return http.StatusInternalServerError
	case ErrorCodeInvalidAccessToken:
		return http.StatusUnauthorized
	case ErrorCodeExtractingAccessTokenClaims:
		return http.StatusInternalServerError
	case ErrorCodeErrorBadAccessToken:
		return http.StatusUnauthorized
	case ErrorCodeAccessTokenNotFound:
		return http.StatusUnauthorized
	case ErrorCodeAccessTokenMalformed:
		return http.StatusBadRequest
	case ErrorCodeComparingPasswords:
		return http.StatusInternalServerError
	case ErrorCodeEndpointSecretKeyNotFound:
		return http.StatusUnauthorized
	case ErrorCodeAppSecretKeyNotFound:
		return http.StatusUnauthorized
	case ErrorCodeInvalidAppSecretKey:
		return http.StatusUnauthorized
	case ErrorCodeEndpointForbidden:
		return http.StatusForbidden
	default:
		return -1
	}
}

func NewError(code ErrorCode) *Error {
	return NewErrorWithDetails(code, "")
}

func NewErrorWithDetails(code ErrorCode, details string) *Error {
	info := make(map[string]string)
	if len(details) > 0 {
		info["details"] = details
	}
	return &Error{
		Code:        code,
		StatusCode:  errorHttpStatusCode(code),
		Description: errorText(code),
		Info:        info,
	}
}
