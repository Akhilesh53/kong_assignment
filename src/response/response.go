package response

import (
	"encoding/json"
	apiErr "ka/pkg/errors"
	log "ka/pkg/logging"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SendResponse(ctx *gin.Context, payload interface{}, Err apiErr.Error, err error) {
	if payload == nil {
		payload = Err
		log.Error(ctx, Err.ErrorMessage, zap.Error(err))
	}
	if err == nil {
		log.Info(ctx, Err.ErrorMessage)
	}

	log.Info(ctx, "Response", zap.Any("data", payload))
	ctx.Writer.Write(InterfaceToBytes(ctx, payload))
	ctx.AbortWithStatus(Err.StatusCode)
}

func InterfaceToBytes(ctx *gin.Context, v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		log.Error(ctx, "Error converting interface to bytes", zap.Any("error", err))
	}
	return b
}
