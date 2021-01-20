
import (
	"github.com/federicoleon/golang-examples/gin_microservice/domain/httperrors"
)
type httpError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error string `json:"error"`,
}
func NewBadRequestError(message string) *httpError {
	return &httpError{
		Message: message,
		Code: http.StatusBadRequest,
		Error: "bad_request",
	}
}


func NewNotFoundError(message string) *httpError {
	return &httpError{
		Message: message,
		Code: http.StatusNotFound,
		Error: "not_found",
	}
}