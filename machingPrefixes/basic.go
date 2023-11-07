package machingPrefixes

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	// ErrInvalidParams is an error returned in case invalid params are passed
	ErrInvalidParams = errors.New("ERR_INVALID_PARAMS")
)

// Data is a type which is added in response struct
type Data interface{}

type response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Data    Data   `json:"data,omitempty"`
}

// returnSuccess adds given object in the context along with the given httpCode.
func returnSuccess(context *gin.Context, obj Data, httpCode int, healthCheckLiveness string) {
	r := response{Message: "SUCCESS", Data: obj}
	if strings.EqualFold(healthCheckLiveness, "liveness") {
		r.Message = "OK"
	}
	context.JSON(httpCode, r)
}

// returnError adds given error in the context along with the given httpCode.
func returnError(context *gin.Context, err error, errMsg string, httpCode int) {
	log.WithFields(log.Fields{
		"errorMsg": errMsg,
		"error":    err,
	}).Error()
	context.Error(err)
	r := response{Error: err.Error(), Message: "ERROR"}
	context.JSON(httpCode, r)
}
