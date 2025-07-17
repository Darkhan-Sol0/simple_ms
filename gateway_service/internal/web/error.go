package web

// type Response struct {
// 	Status  int         `json:"status"`
// 	Data    interface{} `json:"data"`
// 	Err     interface{} `json:"error"`
// 	Details string      `json:"details"`
// }

const (
	errorInvalidJSON             = "invalid request body"
	errorMarshalJSON             = "invalid request json"
	errorResponceInternalService = "internal error"
	errorDecodeJSON              = "invalid decode json"
	errorUnauthorized            = "unauthorized"
	errorToken                   = "Invalid token"
	errorRoleNotFound            = "user role not found"
	errorRoleType                = "invalid user role type"
	errorAccessRole              = "access denied"
	errorUUIDContext             = "user UUID not found in context"
)

func ErrorResponce(errDetails string, status int, err error) Response {
	return Response{
		Status:  status,
		Data:    nil,
		Err:     err,
		Details: errDetails,
	}
}
